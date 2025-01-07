package middleware

import (
	"net/http"
	"time"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
)

// LoggingMiddleware is the middleware to log the request.
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// Create a response writer wrapper to capture the status code
		wwrap := &responseWriter{w, http.StatusOK}

		// Call the next handler
		next(wwrap, r)

		// Log the request details
		clientIP := r.RemoteAddr
		if ip := r.Header.Get("X-Real-IP"); ip != "" {
			clientIP = ip
		} else if ip = r.Header.Get("X-Forwarded-For"); ip != "" {
			clientIP = ip
		}

		logger.Debugf(
			"%s [%s] %s %s %s %d %s \"%s\"",
			clientIP,
			start.Format("02/Jan/2006:15:04:05 -0700"),
			r.Method,
			r.URL.Path,
			r.Proto,
			wwrap.statusCode,
			time.Since(start),
			r.UserAgent(),
		)
	})
}

// responseWriter is a wrapper around http.ResponseWriter that captures the status code.
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code.
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
