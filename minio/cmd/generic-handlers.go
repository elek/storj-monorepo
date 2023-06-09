/*
 * MinIO Cloud Storage, (C) 2015-2020 MinIO, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"context"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	humanize "github.com/dustin/go-humanize"

	"storj.io/minio/cmd/crypto"
	xhttp "storj.io/minio/cmd/http"
	"storj.io/minio/cmd/http/stats"
	"storj.io/minio/cmd/logger"
)

// Adds limiting body size middleware

// Maximum allowed form data field values. 64MiB is a guessed practical value
// which is more than enough to accommodate any form data fields and headers.
const requestFormDataSize = 64 * humanize.MiByte

// For any HTTP request, request body should be not more than 16GiB + requestFormDataSize
// where, 16GiB is the maximum allowed object size for object upload.
const requestMaxBodySize = globalMaxObjectSize + requestFormDataSize

func setRequestSizeLimitHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Restricting read data to a given maximum length
		r.Body = http.MaxBytesReader(w, r.Body, requestMaxBodySize)
		h.ServeHTTP(w, r)
	})
}

const (
	// Maximum size for http headers - See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html
	maxHeaderSize = 8 * 1024
	// Maximum size for user-defined metadata - See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html
	maxUserDataSize = 2 * 1024
)

// ServeHTTP restricts the size of the http header to 8 KB and the size
// of the user-defined metadata to 2 KB.
func setRequestHeaderSizeLimitHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isHTTPHeaderSizeTooLarge(r.Header) {
			WriteErrorResponse(r.Context(), w, errorCodes.ToAPIErr(ErrMetadataTooLarge), r.URL, guessIsBrowserReq(r))
			atomic.AddUint64(&globalHTTPStats.rejectedRequestsHeader, 1)
			return
		}
		h.ServeHTTP(w, r)
	})
}

// isHTTPHeaderSizeTooLarge returns true if the provided
// header is larger than 8 KB or the user-defined metadata
// is larger than 2 KB.
func isHTTPHeaderSizeTooLarge(header http.Header) bool {
	var size, usersize int
	for key := range header {
		length := len(key) + len(header.Get(key))
		size += length
		for _, prefix := range userMetadataKeyPrefixes {
			if strings.HasPrefix(strings.ToLower(key), prefix) {
				usersize += length
				break
			}
		}
		if usersize > maxUserDataSize || size > maxHeaderSize {
			return true
		}
	}
	return false
}

// ReservedMetadataPrefix is the prefix of a metadata key which
// is reserved and for internal use only.
const (
	ReservedMetadataPrefix      = "X-Minio-Internal-"
	ReservedMetadataPrefixLower = "x-minio-internal-"
)

// ServeHTTP fails if the request contains at least one reserved header which
// would be treated as metadata.
func filterReservedMetadata(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if containsReservedMetadata(r.Header) {
			WriteErrorResponse(r.Context(), w, errorCodes.ToAPIErr(ErrUnsupportedMetadata), r.URL, guessIsBrowserReq(r))
			return
		}
		h.ServeHTTP(w, r)
	})
}

// containsReservedMetadata returns true if the http.Header contains
// keys which are treated as metadata but are reserved for internal use
// and must not set by clients
func containsReservedMetadata(header http.Header) bool {
	for key := range header {
		if strings.HasPrefix(strings.ToLower(key), ReservedMetadataPrefixLower) {
			return true
		}
	}
	return false
}

// Reserved bucket.
const (
	minioReservedBucket     = "minio"
	minioReservedBucketPath = SlashSeparator + minioReservedBucket
	loginPathPrefix         = SlashSeparator + "login"
)

func setRedirectHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !shouldProxy() || guessIsRPCReq(r) || guessIsBrowserReq(r) ||
			guessIsHealthCheckReq(r) || guessIsMetricsReq(r) || isAdminReq(r) {
			h.ServeHTTP(w, r)
			return
		}
		// if this server is still initializing, proxy the request
		// to any other online servers to avoid 503 for any incoming
		// API calls.
		if idx := getOnlineProxyEndpointIdx(); idx >= 0 {
			proxyRequest(context.TODO(), w, r, globalProxyEndpoints[idx])
			return
		}
		h.ServeHTTP(w, r)
	})
}

func setBrowserRedirectHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Re-direction is handled specifically for browser requests.
		if globalBrowserEnabled && guessIsBrowserReq(r) {
			// Fetch the redirect location if any.
			redirectLocation := getRedirectLocation(r.URL.Path)
			if redirectLocation != "" {
				// Employ a temporary re-direct.
				http.Redirect(w, r, redirectLocation, http.StatusTemporaryRedirect)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func shouldProxy() bool {
	if newObjectLayerFn() == nil {
		return true
	}
	return !GlobalIAMSys.Initialized()
}

// Fetch redirect location if urlPath satisfies certain
// criteria. Some special names are considered to be
// redirectable, this is purely internal function and
// serves only limited purpose on redirect-handler for
// browser requests.
func getRedirectLocation(urlPath string) (rLocation string) {
	if urlPath == minioReservedBucketPath {
		rLocation = minioReservedBucketPath + SlashSeparator
	}
	if contains([]string{
		SlashSeparator,
		"/webrpc",
		"/login",
		"/favicon-16x16.png",
		"/favicon-32x32.png",
		"/favicon-96x96.png",
	}, urlPath) {
		rLocation = minioReservedBucketPath + urlPath
	}
	return rLocation
}

// guessIsBrowserReq - returns true if the request is browser.
// This implementation just validates user-agent and
// looks for "Mozilla" string. This is no way certifiable
// way to know if the request really came from a browser
// since User-Agent's can be arbitrary. But this is just
// a best effort function.
func guessIsBrowserReq(req *http.Request) bool {
	if req == nil {
		return false
	}
	aType := getRequestAuthType(req)
	return strings.Contains(req.Header.Get("User-Agent"), "Mozilla") && globalBrowserEnabled &&
		(aType == authTypeJWT || aType == authTypeAnonymous)
}

// guessIsHealthCheckReq - returns true if incoming request looks
// like healthcheck request
func guessIsHealthCheckReq(req *http.Request) bool {
	if req == nil {
		return false
	}
	aType := getRequestAuthType(req)
	return aType == authTypeAnonymous && (req.Method == http.MethodGet || req.Method == http.MethodHead) &&
		(req.URL.Path == healthCheckPathPrefix+healthCheckLivenessPath ||
			req.URL.Path == healthCheckPathPrefix+healthCheckReadinessPath ||
			req.URL.Path == healthCheckPathPrefix+healthCheckClusterPath ||
			req.URL.Path == healthCheckPathPrefix+healthCheckClusterReadPath)
}

// guessIsMetricsReq - returns true if incoming request looks
// like metrics request
func guessIsMetricsReq(req *http.Request) bool {
	if req == nil {
		return false
	}
	aType := getRequestAuthType(req)
	return (aType == authTypeAnonymous || aType == authTypeJWT) &&
		req.URL.Path == minioReservedBucketPath+prometheusMetricsPathLegacy ||
		req.URL.Path == minioReservedBucketPath+prometheusMetricsV2ClusterPath ||
		req.URL.Path == minioReservedBucketPath+prometheusMetricsV2NodePath
}

// guessIsRPCReq - returns true if the request is for an RPC endpoint.
func guessIsRPCReq(req *http.Request) bool {
	if req == nil {
		return false
	}
	return req.Method == http.MethodPost &&
		strings.HasPrefix(req.URL.Path, minioReservedBucketPath+SlashSeparator)
}

// Adds Cache-Control header
func setBrowserCacheControlHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if globalBrowserEnabled && r.Method == http.MethodGet && guessIsBrowserReq(r) {
			// For all browser requests set appropriate Cache-Control policies
			if HasPrefix(r.URL.Path, minioReservedBucketPath+SlashSeparator) {
				if HasSuffix(r.URL.Path, ".js") || r.URL.Path == minioReservedBucketPath+"/favicon.ico" {
					// For assets set cache expiry of one year. For each release, the name
					// of the asset name will change and hence it can not be served from cache.
					w.Header().Set(xhttp.CacheControl, "max-age=31536000")
				} else {
					// For non asset requests we serve index.html which will never be cached.
					w.Header().Set(xhttp.CacheControl, "no-store")
				}
			}
		}

		h.ServeHTTP(w, r)
	})
}

// Check to allow access to the reserved "bucket" `/minio` for Admin
// API requests.
func isAdminReq(r *http.Request) bool {
	return strings.HasPrefix(r.URL.Path, adminPathPrefix)
}

// guessIsLoginSTSReq - returns true if incoming request is Login STS user
func guessIsLoginSTSReq(req *http.Request) bool {
	if req == nil {
		return false
	}
	return strings.HasPrefix(req.URL.Path, loginPathPrefix) ||
		(req.Method == http.MethodPost && req.URL.Path == SlashSeparator &&
			getRequestAuthType(req) == authTypeSTS)
}

// Adds verification for incoming paths.
func setReservedBucketHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// For all other requests reject access to reserved buckets
		bucketName, _ := request2BucketObjectName(r)
		if isMinioReservedBucket(bucketName) || isMinioMetaBucket(bucketName) {
			if !guessIsRPCReq(r) && !guessIsBrowserReq(r) && !guessIsHealthCheckReq(r) && !guessIsMetricsReq(r) && !isAdminReq(r) {
				WriteErrorResponse(r.Context(), w, errorCodes.ToAPIErr(ErrAllAccessDisabled), r.URL, guessIsBrowserReq(r))
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

// Supported Amz date formats.
var amzDateFormats = []string{
	time.RFC1123,
	time.RFC1123Z,
	iso8601Format,
	// Add new AMZ date formats here.
}

// Supported Amz date headers.
var amzDateHeaders = []string{
	"x-amz-date",
	"date",
}

// parseAmzDate - parses date string into supported amz date formats.
func parseAmzDate(amzDateStr string) (amzDate time.Time, apiErr APIErrorCode) {
	for _, dateFormat := range amzDateFormats {
		amzDate, err := time.Parse(dateFormat, amzDateStr)
		if err == nil {
			return amzDate, ErrNone
		}
	}
	return time.Time{}, ErrMissingDateHeader
}

// parseAmzDateHeader - parses supported amz date headers, in
// supported amz date formats.
func parseAmzDateHeader(req *http.Request) (time.Time, APIErrorCode) {
	for _, amzDateHeader := range amzDateHeaders {
		amzDateStr := req.Header.Get(amzDateHeader)
		if amzDateStr != "" {
			return parseAmzDate(amzDateStr)
		}
	}
	// Date header missing.
	return time.Time{}, ErrMissingDateHeader
}

// setTimeValidityHandler to validate parsable time over http header
func setTimeValidityHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		aType := getRequestAuthType(r)
		if aType == authTypeSigned || aType == authTypeSignedV2 || aType == authTypeStreamingSigned {
			// Verify if date headers are set, if not reject the request
			amzDate, errCode := parseAmzDateHeader(r)
			if errCode != ErrNone {
				// All our internal APIs are sensitive towards Date
				// header, for all requests where Date header is not
				// present we will reject such clients.
				WriteErrorResponse(r.Context(), w, errorCodes.ToAPIErr(errCode), r.URL, guessIsBrowserReq(r))
				atomic.AddUint64(&globalHTTPStats.rejectedRequestsTime, 1)
				return
			}
			// Verify if the request date header is shifted by less than globalMaxSkewTime parameter in the past
			// or in the future, reject request otherwise.
			curTime := UTCNow()
			if curTime.Sub(amzDate) > globalMaxSkewTime || amzDate.Sub(curTime) > globalMaxSkewTime {
				WriteErrorResponse(r.Context(), w, errorCodes.ToAPIErr(ErrRequestTimeTooSkewed), r.URL, guessIsBrowserReq(r))
				atomic.AddUint64(&globalHTTPStats.rejectedRequestsTime, 1)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

// setHttpStatsHandler sets a http Stats handler to gather HTTP statistics
func setHTTPStatsHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Meters s3 connection stats.
		meteredRequest := &stats.IncomingTrafficMeter{ReadCloser: r.Body}
		meteredResponse := &stats.OutgoingTrafficMeter{ResponseWriter: w}

		// Execute the request
		r.Body = meteredRequest
		h.ServeHTTP(meteredResponse, r)

		if strings.HasPrefix(r.URL.Path, minioReservedBucketPath) {
			globalConnStats.incInputBytes(meteredRequest.BytesCount())
			globalConnStats.incOutputBytes(meteredResponse.BytesCount())
		} else {
			globalConnStats.incS3InputBytes(meteredRequest.BytesCount())
			globalConnStats.incS3OutputBytes(meteredResponse.BytesCount())
		}
	})
}

// Bad path components to be rejected by the path validity handler.
const (
	dotdotComponent = ".."
	dotComponent    = "."
)

// Check if the incoming path has bad path components,
// such as ".." and "."
func hasBadPathComponent(path string) bool {
	path = strings.TrimSpace(path)
	for _, p := range strings.Split(path, SlashSeparator) {
		switch strings.TrimSpace(p) {
		case dotdotComponent:
			return true
		case dotComponent:
			return true
		}
	}
	return false
}

// Check if client is sending a malicious request.
func hasMultipleAuth(r *http.Request) bool {
	authTypeCount := 0
	for _, hasValidAuth := range []func(*http.Request) bool{isRequestSignatureV2, isRequestPresignedSignatureV2, isRequestSignatureV4, isRequestPresignedSignatureV4, isRequestJWT, isRequestPostPolicySignatureV4} {
		if hasValidAuth(r) {
			authTypeCount++
		}
	}
	return authTypeCount > 1
}

// requestValidityHandler validates all the incoming paths for
// any malicious requests.
func setRequestValidityHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for bad components in URL path.
		if hasBadPathComponent(r.URL.Path) {
			WriteErrorResponse(r.Context(), w, errorCodes.ToAPIErr(ErrInvalidResourceName), r.URL, guessIsBrowserReq(r))
			atomic.AddUint64(&globalHTTPStats.rejectedRequestsInvalid, 1)
			return
		}
		// Check for bad components in URL query values.
		for q, vv := range r.URL.Query() {
			// TODO: ideally we'd have a precise list of fields to check here
			// however, we know that "delimiter" must be skipped for the Ceph / Splunk tests
			if q == "delimiter" {
				continue
			}
			for _, v := range vv {
				if hasBadPathComponent(v) {
					WriteErrorResponse(r.Context(), w, errorCodes.ToAPIErr(ErrInvalidResourceName), r.URL, guessIsBrowserReq(r))
					atomic.AddUint64(&globalHTTPStats.rejectedRequestsInvalid, 1)
					return
				}
			}
		}
		if hasMultipleAuth(r) {
			WriteErrorResponse(r.Context(), w, errorCodes.ToAPIErr(ErrInvalidRequest), r.URL, guessIsBrowserReq(r))
			atomic.AddUint64(&globalHTTPStats.rejectedRequestsInvalid, 1)
			return
		}
		h.ServeHTTP(w, r)
	})
}

// customHeaderHandler sets x-amz-request-id header.
// Previously, this value was set right before a response was sent to
// the client. So, logger and Error response XML were not using this
// value. This is set here so that this header can be logged as
// part of the log entry, Error response XML and auditing.
func addCustomHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set custom headers such as x-amz-request-id for each request.
		w.Header().Set(xhttp.AmzRequestID, mustGetRequestID(UTCNow()))
		h.ServeHTTP(logger.NewResponseWriter(w), r)
	})
}

func addSecurityHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Set("X-XSS-Protection", "1; mode=block")                  // Prevents against XSS attacks
		header.Set("Content-Security-Policy", "block-all-mixed-content") // prevent mixed (HTTP / HTTPS content)
		h.ServeHTTP(w, r)
	})
}

// criticalErrorHandler handles critical server failures caused by
// `panic(logger.ErrCritical)` as done by `logger.CriticalIf`.
//
// It should be always the first / highest HTTP handler.
type criticalErrorHandler struct{ handler http.Handler }

func (h criticalErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err == logger.ErrCritical { // handle
			WriteErrorResponse(r.Context(), w, errorCodes.ToAPIErr(ErrInternalError), r.URL, guessIsBrowserReq(r))
			return
		} else if err != nil {
			panic(err) // forward other panic calls
		}
	}()
	h.handler.ServeHTTP(w, r)
}

// sseTLSHandler enforces certain rules for SSE requests which are made / must be made over TLS.
func setSSETLSHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Deny SSE-C requests if not made over TLS
		if !GlobalIsTLS && (crypto.SSEC.IsRequested(r.Header) || crypto.SSECopy.IsRequested(r.Header)) {
			if r.Method == http.MethodHead {
				writeErrorResponseHeadersOnly(w, errorCodes.ToAPIErr(ErrInsecureSSECustomerRequest))
			} else {
				WriteErrorResponse(r.Context(), w, errorCodes.ToAPIErr(ErrInsecureSSECustomerRequest), r.URL, guessIsBrowserReq(r))
			}
			return
		}
		h.ServeHTTP(w, r)
	})
}
