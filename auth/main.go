// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package auth

import (
	"context"
	"encoding/hex"
	"net"
	"net/http"
	"net/url"
	"sort"
	"time"

	"github.com/spacemonkeygo/monkit/v3"
	"github.com/zeebo/errs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"

	"storj.io/common/sync2"
	"storj.io/gateway-mt/auth/authdb"
	"storj.io/gateway-mt/auth/drpcauth"
	"storj.io/gateway-mt/auth/httpauth"
	"storj.io/gateway-mt/auth/satellitelist"
	"storj.io/gateway-mt/pkg/server"
)

const serverShutdownTimeout = 10 * time.Second

// Config holds authservice's configuration.
type Config struct {
	Endpoint          string        `help:"Gateway endpoint URL to return to clients" default:""`
	AuthToken         string        `help:"auth security token to validate requests" releaseDefault:"" devDefault:""`
	AllowedSatellites []string      `help:"list of satellite NodeURLs allowed for incoming access grants" default:"https://www.storj.io/dcs-satellites"`
	CacheExpiration   time.Duration `help:"length of time satellite addresses are cached for" default:"10m"`

	KVBackend string `help:"key/value store backend url" default:""`
	Migration bool   `help:"create or update the database schema, and then continue service startup" default:"false"`

	ListenAddr    string `user:"true" help:"public HTTP address to listen on" default:":8000"`
	ListenAddrTLS string `user:"true" help:"public HTTPS and DRPC+TLS address to listen on" default:":8443"`

	LetsEncrypt bool   `user:"true" help:"use lets-encrypt to handle TLS certificates" default:"false"`
	CertFile    string `user:"true" help:"server certificate file" default:""`
	KeyFile     string `user:"true" help:"server key file" default:""`
	PublicURL   string `user:"true" help:"public url for the server, for the TLS certificate" devDefault:"http://localhost:8080" releaseDefault:""`

	DeleteUnused DeleteUnusedConfig
}

// DeleteUnusedConfig is a config struct for configuring unused records deletion
// chores.
type DeleteUnusedConfig struct {
	Run                bool          `help:"whether to run unused records deletion chore" default:"false"`
	Interval           time.Duration `help:"interval unused records deletion chore waits to start next iteration" default:"24h"`
	AsOfSystemInterval time.Duration `help:"the interval specified in AS OF SYSTEM in unused records deletion chore query as negative interval" default:"5s"`
	SelectSize         int           `help:"batch size of records selected for deletion at a time" default:"10000"`
	DeleteSize         int           `help:"batch size of records to delete from selected records at a time" default:"1000"`
}

// Run contains everything needed to start authservice.
func Run(ctx context.Context, config Config, confDir string, log *zap.Logger) error {
	if len(config.AllowedSatellites) == 0 {
		return errs.New("allowed satellites parameter '--allowed-satellites' is required")
	}
	allowedSats, areSatsDynamic, err := satellitelist.LoadSatelliteIDs(ctx, config.AllowedSatellites)
	if err != nil {
		return errs.Wrap(err)
	}
	if len(allowedSats) == 0 {
		return errs.New("allowed satellites parameter '--allowed-satellites' resolved to zero satellites")
	}

	if config.Endpoint == "" {
		return errs.New("endpoint parameter '--endpoint' is required")
	}
	endpoint, err := url.Parse(config.Endpoint)
	if err != nil {
		return errs.Wrap(err)
	}
	if endpoint.Scheme != "http" && endpoint.Scheme != "https" {
		return errs.New("unexpected scheme found in endpoint parameter %s", endpoint.Scheme)
	}

	kv, err := OpenKV(ctx, log.Named("db"), config.KVBackend)
	if err != nil {
		return errs.Wrap(err)
	}
	defer func() { err = errs.Combine(err, errs.Wrap(kv.Close())) }()

	db := authdb.NewDatabase(kv, allowedSats)
	res := httpauth.New(log.Named("resources"), db, endpoint, config.AuthToken)

	tlsInfo := &TLSInfo{
		LetsEncrypt: config.LetsEncrypt,
		CertFile:    config.CertFile,
		KeyFile:     config.KeyFile,
		PublicURL:   config.PublicURL,
		ConfigDir:   confDir,
	}

	tlsConfig, handler, err := configureTLS(tlsInfo, res)
	if err != nil {
		return errs.Wrap(err)
	}

	// logging. do not log paths - paths have access keys in them.
	handler = server.LogResponses(log, server.LogRequests(log, handler, false), false)

	httpServer := &http.Server{
		Addr:    config.ListenAddr,
		Handler: handler,
	}

	group, groupCtx := errgroup.WithContext(ctx)

	if areSatsDynamic {
		sync2.NewCycle(config.CacheExpiration).Start(groupCtx, group, func(ctx context.Context) error {
			log.Debug("Reloading allowed satellite list")
			allowedSatelliteIDs, _, err := satellitelist.LoadSatelliteIDs(ctx, config.AllowedSatellites)
			if err != nil {
				log.Warn("Error reloading allowed satellite list", zap.Error(err))
			} else {
				db.SetAllowedSatellites(allowedSatelliteIDs)
			}
			return nil
		})
	}

	group.Go(func() error {
		<-groupCtx.Done()

		// Don't wait more than a couple of seconds to shut down the server.
		ctxWithTimeout, canc := context.WithTimeout(context.Background(), serverShutdownTimeout)
		defer canc()

		// We received an interrupt signal, shut down.
		if err := httpServer.Shutdown(ctxWithTimeout); err != nil {
			// Error from closing listeners, or context timeout:
			log.Error("HTTP server Shutdown", zap.Error(err))
		}

		return nil
	})

	group.Go(func() error {
		log.Info("listening for incoming HTTP connections", zap.String("addr", config.ListenAddr))

		if !errs.Is(httpServer.ListenAndServe(), http.ErrServerClosed) {
			// Error starting or closing listener:
			return err
		}

		return nil
	})

	group.Go(func() (err error) {
		listener, err := net.Listen("tcp", config.ListenAddrTLS)
		if err != nil {
			return err
		}

		defer func() { err = errs.Combine(err, listener.Close()) }()

		res.SetStartupDone()

		server := drpcauth.NewServer(groupCtx, log, db, endpoint)

		return listenAndServe(groupCtx, log, listener, tlsConfig, server, handler)
	})

	if config.DeleteUnused.Run {
		sync2.NewCycle(config.DeleteUnused.Interval).Start(groupCtx, group, func(ctx context.Context) error {
			log.Info("Beginning of next iteration of unused records deletion chore")

			count, rounds, heads, err := db.DeleteUnused(
				ctx,
				config.DeleteUnused.AsOfSystemInterval,
				config.DeleteUnused.SelectSize,
				config.DeleteUnused.DeleteSize)
			if err != nil {
				log.Warn("Error deleting unused records", zap.Error(err))
			}

			log.Info(
				"Deleted unused records",
				zap.Int64("count", count),
				zap.Int64("rounds", rounds),
				zap.Array("heads", headsMapToLoggableHeads(heads)))

			monkit.Package().IntVal("authservice_deleted_unused_records_count").Observe(count)
			monkit.Package().IntVal("authservice_deleted_unused_records_rounds").Observe(rounds)

			for h, c := range heads {
				monkit.Package().IntVal(
					"authservice_deleted_unused_records_deletes_per_head",
					monkit.NewSeriesTag("head", hex.EncodeToString([]byte(h))),
				).Observe(c)
			}

			return nil
		})
	}

	return errs.Wrap(group.Wait())
}

func headsMapToLoggableHeads(heads map[string]int64) zapcore.ArrayMarshalerFunc {
	type loggableHead struct {
		head  string
		count int64
	}

	var loggableHeads []loggableHead

	for k, v := range heads {
		loggableHeads = append(loggableHeads, loggableHead{head: k, count: v})
	}

	sort.Slice(loggableHeads, func(i, j int) bool {
		return loggableHeads[i].count > loggableHeads[j].count
	})

	return zapcore.ArrayMarshalerFunc(func(ae zapcore.ArrayEncoder) error {
		for _, h := range loggableHeads {
			if err := ae.AppendObject(zapcore.ObjectMarshalerFunc(func(oe zapcore.ObjectEncoder) error {
				oe.AddInt64(hex.EncodeToString([]byte(h.head)), h.count)
				return nil
			})); err != nil {
				return err
			}
		}
		return nil
	})
}
