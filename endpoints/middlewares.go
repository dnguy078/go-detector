package endpoints

import (
	"log"
	"net/http"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// WithLogging is a middleware that wraps a http handler and logs the statuscodes.
func WithLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lrw := newLoggingResponseWriter(w)
		next.ServeHTTP(lrw, r)
		statusCode := lrw.statusCode
		log.Printf("Received Request %s %s - %d %s", r.Method, r.URL.Path, statusCode, http.StatusText(statusCode))
	}
}
