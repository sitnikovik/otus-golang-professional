package http

import (
	"net/http"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
)

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc { //nolint:unused
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Debugf("Request: %s %s", r.Method, r.URL.Path)
		next(w, r)
	})
}
