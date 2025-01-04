package http

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/config"
	eventFilter "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/filter/event"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
)

// eventService describes the event service.
type eventService interface {
	// CreateEvent creates a new event.
	CreateEvent(ctx context.Context, event *eventModel.Event) (uint64, error)
	// UpdateEvent updates the event.
	UpdateEvent(ctx context.Context, event *eventModel.Event) error
	// DeleteEvent deletes the event.
	DeleteEvent(ctx context.Context, eventID uint64) error
	// GetEvent returns the event by ID.
	GetEvent(ctx context.Context, eventID uint64) (*eventModel.Event, error)
	// GetEvents returns the events by filter.
	GetEvents(ctx context.Context, filter eventFilter.Filter) ([]*eventModel.Event, error)
	// GetEventsBeforeDays returns the events before days.
	GetEventsBeforeDays(ctx context.Context, days uint32) ([]*eventModel.Event, error)
	// GetEventsForMonth returns a list of events that will occur in the current month.
	GetEventsForMonth(ctx context.Context) ([]*eventModel.Event, error)
	// GetEventsForWeek returns a list of events that will occur in the current week.
	GetEventsForWeek(ctx context.Context) ([]*eventModel.Event, error)
	// GetEventsForToday returns a list of events that will occur today.
	GetEventsForToday(ctx context.Context) ([]*eventModel.Event, error)
}

// Server is the HTTP server.
type Server struct {
	httpconf config.HTTPConf
	server   *http.Server

	// eventService describes the event service to work with events.
	eventService eventService
}

// NewServer creates a new HTTP server.
func NewServer(conf config.HTTPConf, eventsService eventService) *Server {
	s := &Server{
		httpconf:     conf,
		eventService: eventsService,
	}
	// Create a new HTTP server
	s.server = &http.Server{
		Addr:              conf.Host + ":" + conf.Port,
		Handler:           s.routes(),
		ReadTimeout:       time.Duration(conf.Timeout) * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	return s
}

// Serve starts the HTTP server.
func (s *Server) Serve(ctx context.Context) error {
	// Run the server in a separate goroutine
	go func() {
		logger.Infof("Starting HTTP server on %s", s.server.Addr)
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Panicf("listen err: %s", err)
		}
	}()

	// Wait for the context to be done
	<-ctx.Done()

	// Shutdown the server with a timeout
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.GracefulShutdown(ctxShutDown)

	return nil
}

// GracefulShutdown gracefully stops the HTTP server.
func (s *Server) GracefulShutdown(ctx context.Context) {
	if err := s.server.Shutdown(ctx); err != nil {
		logger.Emergencyf("failed to stop http server: %v", err)
	}
}

// Stop stops the HTTP server.
func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

// errorHandler is the handler to handle errors.
func errorHandler(w http.ResponseWriter, err error, httpCode int) {
	logger.Error(err.Error())
	http.Error(w, err.Error(), httpCode)
}
