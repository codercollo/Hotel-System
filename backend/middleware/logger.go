package middleware

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/codercollo/hotel-system/backend/pkg/logger"
)

// Logger logs every HTTP request with method, path, status, latency, and request ID.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rw, r)
		l := logger.WithRequestID(GetRequestID(r.Context()))
		l.Info().
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Int("status", rw.status).
			Dur("latency", time.Since(start)).
			Str("ip", r.RemoteAddr).
			Msg("request")
	})
}

// statusRecorder captures the HTTP status code written by the handler.
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (sr *statusRecorder) WriteHeader(code int) {
	sr.status = code
	sr.ResponseWriter.WriteHeader(code)
}

// Hijack implements http.Hijacker so that WebSocket upgrades work when the
// logger middleware is in the chain. Delegates to the underlying ResponseWriter.
// Without this, golang.org/x/net/websocket panics with:
//
//	"interface conversion: *middleware.statusRecorder is not http.Hijacker"
func (sr *statusRecorder) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := sr.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, fmt.Errorf("middleware: underlying ResponseWriter %T does not implement http.Hijacker", sr.ResponseWriter)
	}
	return h.Hijack()
}
