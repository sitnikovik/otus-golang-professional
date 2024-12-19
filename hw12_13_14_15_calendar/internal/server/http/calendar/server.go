package calendar

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/app"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/config"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
)

// Server is the HTTP server
type Server struct {
	app      *app.App
	httpconf config.HTTPConf
	server   *http.Server
}

// NewServer creates a new HTTP server
func NewServer(app *app.App, httpconf config.HTTPConf) *Server {
	return &Server{
		app:      app,
		httpconf: httpconf,
	}
}

// Start starts the HTTP server
func (s *Server) Start(ctx context.Context) error {
	// Create a new HTTP server
	s.server = &http.Server{
		Addr:    ":" + s.httpconf.Port,
		Handler: s.routes(),
	}

	// Recover from panics
	defer func() {
		if r := recover(); r != nil {
			logger.Emergencyf("recovered from panic: %v", r)
		}
	}()

	// Run the server in a separate goroutine
	go func() {
		logger.Infof("Starting HTTP server on %s", s.server.Addr)
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Emergencyf("listen err: %s", err)
		}
	}()

	// Wait for the context to be done
	<-ctx.Done()

	// Shutdown the server with a timeout
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.server.Shutdown(ctxShutDown)
}

// Stop stops the HTTP server
func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

// errorHandler is the handler to handle errors
func errorHandler(w http.ResponseWriter, err error, httpCode int) {
	logger.Error(err.Error())
	http.Error(w, err.Error(), httpCode)
}
