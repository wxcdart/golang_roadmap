package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/http"
	"net/http/httptest"

	"go.uber.org/zap"
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
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	// attach trace field
	logger.Info("handling request", zap.String("trace", tid), zap.String("handler", "zap-trace"))
}

func main() {
	h := withTrace(http.HandlerFunc(handler))
	req := httptest.NewRequest(http.MethodGet, "http://example.local/", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
}
