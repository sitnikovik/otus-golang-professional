package http

import (
	"net/http"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
)

// loggingMiddleware is the middleware to log the request
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Debugf("Request: %s %s", r.Method, r.URL.Path)
		next(w, r)
	})
}
