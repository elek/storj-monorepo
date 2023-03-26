// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package rpc

import (
	"context"
	"math/rand"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"storj.io/common/storj"
)

// quicRolloutPercent is the node rollout between 0 and 100.
// -1 means no satellites either.
const quicRolloutPercentDefault = -1

var quicRolloutPercent int

func init() {
	if percent, err := strconv.Atoi(os.Getenv("STORJ_QUIC_ROLLOUT_PERCENT")); err == nil {
		quicRolloutPercent = percent
	} else {
		quicRolloutPercent = quicRolloutPercentDefault
	}
}

var (
	rngState uint64
	rngInc   = uint64(time.Now().UnixNano())
)

func intn(n int) int {
	return rand.New(rand.NewSource(int64(atomic.AddUint64(&rngState, rngInc)))).Intn(n)
}

type quicRolloutPercentKey struct{}

// WithQUICRolloutPercent sets the QUIC rollout percent for any dials that
// use the returned context.
func WithQUICRolloutPercent(ctx context.Context, percent int) context.Context {
	return context.WithValue(ctx, quicRolloutPercentKey{}, percent)
}

// QUICRolloutPercent returns the currently configured QUIC rollout percent
// for the given context.
func QUICRolloutPercent(ctx context.Context) int {
	if ctxpercent, ok := ctx.Value(quicRolloutPercentKey{}).(int); ok {
		return ctxpercent
	}
	return quicRolloutPercent
}

func checkQUICRolloutState(ctx context.Context, id storj.NodeID) bool {
	percent := QUICRolloutPercent(ctx)

	if percent >= 100 {
		return true
	}

	if percent < 0 {
		return false
	}

	if quicRolloutSatelliteIDs[id] {
		return true
	}

	if intn(100) < percent {
		return true
	}

	return false
}

// setQUICRollout set up forced TCP if not yet set, based on rollout rules.
func setQUICRollout(ctx context.Context, nodeURL storj.NodeURL) context.Context {
	forced := ctx.Value(hybridConnectorForcedKind{})
	if forced != nil && forced != "" {
		return ctx
	}
	if !checkQUICRolloutState(ctx, nodeURL.ID) {
		ctx = WithForcedKind(ctx, "tcp")
	}
	return ctx
}

var quicRolloutSatelliteIDs = map[storj.NodeID]bool{
	{0xa2, 0x8b, 0x4f, 0x4, 0xe1, 0xb, 0xae, 0x85, 0xd6, 0x7f, 0x4c, 0x6c, 0xb8, 0x2b, 0xf8, 0xd4, 0xc0, 0xf0, 0xf4, 0x7a, 0x8e, 0xa7, 0x26, 0x27, 0x52, 0x4d, 0xeb, 0x6e, 0xc0, 0x0, 0x0, 0x0}:   true,
	{0x84, 0xa7, 0x4c, 0x2c, 0xd4, 0x3c, 0x5b, 0xa7, 0x65, 0x35, 0xe1, 0xf4, 0x2f, 0x5d, 0xf7, 0xc2, 0x87, 0xed, 0x68, 0xd3, 0x35, 0x22, 0x78, 0x2f, 0x4a, 0xfa, 0xbf, 0xdb, 0x40, 0x0, 0x0, 0x0}: true,
	{0xaf, 0x2c, 0x42, 0x0, 0x3e, 0xfc, 0x82, 0x6a, 0xb4, 0x36, 0x1f, 0x73, 0xf9, 0xd8, 0x90, 0x94, 0x21, 0x46, 0xfe, 0xe, 0xbe, 0x80, 0x67, 0x86, 0xf8, 0xe7, 0x19, 0x8, 0x0, 0x0, 0x0, 0x0}:     true,
	{0xf4, 0x74, 0x53, 0x5a, 0x19, 0xdb, 0x0, 0xdb, 0x4f, 0x80, 0x71, 0xa1, 0xbe, 0x6c, 0x25, 0x51, 0xf4, 0xde, 0xd6, 0xa6, 0xe3, 0x8f, 0x8, 0x18, 0xc6, 0x8c, 0x68, 0xd0, 0x0, 0x0, 0x0, 0x0}:    true,
	{0x7b, 0x2d, 0xe9, 0xd7, 0x2c, 0x2e, 0x93, 0x5f, 0x19, 0x18, 0xc0, 0x58, 0xca, 0xaf, 0x8e, 0xd0, 0xf, 0x5, 0x81, 0x63, 0x90, 0x8, 0x70, 0x73, 0x17, 0xff, 0x1b, 0xd0, 0x0, 0x0, 0x0, 0x0}:     true,
	{0x4, 0x48, 0x9f, 0x52, 0x45, 0xde, 0xd4, 0x8d, 0x2a, 0x8a, 0xc8, 0xfb, 0x5f, 0x5c, 0xd1, 0xc6, 0xa6, 0x38, 0xf7, 0xc6, 0xe7, 0x5e, 0xfd, 0x80, 0xe, 0xf2, 0xd7, 0x20, 0x0, 0x0, 0x0, 0x0}:    true,
}
