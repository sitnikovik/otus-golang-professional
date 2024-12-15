package http

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/app/depinjection"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/config"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
)

type Server struct {
	app      Application
	httpconf config.HTTPConf
	server   *http.Server
}

type Application interface { // TODO
	DIContainer() depinjection.DIContainer
}

// NewServer creates a new HTTP server
func NewServer(app Application, conf config.HTTPConf) *Server {
	return &Server{
		app:      app,
		httpconf: conf,
	}
}

// Start starts the HTTP server
func (s *Server) Start(ctx context.Context) error {
	// Create a new HTTP server
	s.server = &http.Server{
		Addr:    ":" + s.httpconf.Port,
		Handler: s.routes(),
	}
	// Run the server in a separate goroutine
	go func() {
		logger.Infof("Starting HTTP server on %s", s.server.Addr)
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Criticalf("listen err: %s", err)
		}
	}()

	// Wait for the context to be done
	<-ctx.Done()

	// Shutdown the server with a timeout
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.server.Shutdown(ctxShutDown)
}

func (s *Server) Stop(ctx context.Context) error {
	// TODO: Implement any additional cleanup if needed
	return s.server.Shutdown(ctx)
}

// routes defines the routes of the HTTP server
func (s *Server) routes() http.Handler {
	mux := mux.NewRouter()

	// Index
	mux.HandleFunc(
		"/",
		loggingMiddleware(
			s.handlerIndex(),
		),
	).Methods(http.MethodGet)

	// Create a new event
	mux.HandleFunc(
		"/event",
		loggingMiddleware(
			s.handlerCreateEvent(),
		),
	).Methods(http.MethodPut)

	// Get event by ID
	mux.HandleFunc(
		"/event/{id}",
		loggingMiddleware(
			s.handlerGetEvent(),
		),
	).Methods(http.MethodGet)

	// Update an event
	mux.HandleFunc(
		"/event/{id}",
		loggingMiddleware(
			s.handlerUpdateEvent(),
		),
	).Methods(http.MethodPost)

	// Delete an event
	mux.HandleFunc(
		"/event/{id}",
		loggingMiddleware(
			s.handlerDeleteEvent(),
		),
	).Methods(http.MethodDelete)

	return mux
}
