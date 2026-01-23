package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/rs/zerolog"
)

type traceKey struct{}

func newTraceID() string {
	b := make([]byte, 8)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func withTrace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tid := r.Header.Get("X-Trace-ID")
		if tid == "" {
			tid = newTraceID()
		}
		ctx := context.WithValue(r.Context(), traceKey{}, tid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	tid, _ := r.Context().Value(traceKey{}).(string)
	// attach trace to log
	zerolog.TimeFieldFormat = time.RFC3339
	logger := zerolog.New(zerolog.ConsoleWriter{Out: w}).With().Timestamp().Str("trace", tid).Logger()
	logger.Info().Str("handler", "zerolog-trace").Msg("handling request")
}

func main() {
	// simulate a request through middleware
	h := withTrace(http.HandlerFunc(handler))
	req := httptest.NewRequest(http.MethodGet, "http://example.local/", nil)
	// optional: set incoming trace id header
	req.Header.Set("X-Trace-ID", "")
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	fmt.Print("--- response body above ---\n")
}
