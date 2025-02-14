package http

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/api/calendar/v1/http/middleware"
)

// routes defines the routes of the HTTP server.
func (s *Server) routes() http.Handler {
	mux := mux.NewRouter()

	// Hello page
	mux.HandleFunc(
		"/hello",
		middleware.LoggingMiddleware(
			s.handlerHello(),
		),
	).Methods(http.MethodGet)

	// Create a new event
	mux.HandleFunc(
		"/event",
		middleware.LoggingMiddleware(
			s.handlerCreateEvent(),
		),
	).Methods(http.MethodPut)

	// Get event list
	mux.HandleFunc(
		"/event",
		middleware.LoggingMiddleware(
			s.handlerGetEvents(),
		),
	).Methods(http.MethodGet)

	// Get event by ID
	mux.HandleFunc(
		"/event/{id}",
		middleware.LoggingMiddleware(
			s.handlerGetEvent(),
		),
	).Methods(http.MethodGet)

	// Update an event
	mux.HandleFunc(
		"/event/{id}",
		middleware.LoggingMiddleware(
			s.handlerUpdateEvent(),
		),
	).Methods(http.MethodPost)

	// Delete an event
	mux.HandleFunc(
		"/event/{id}",
		middleware.LoggingMiddleware(
			s.handlerDeleteEvent(),
		),
	).Methods(http.MethodDelete)

	// Get events before days
	mux.HandleFunc(
		"/event/before/{days}",
		middleware.LoggingMiddleware(
			s.GetEventsBeforeDays(),
		),
	)

	// Get events for today
	mux.HandleFunc(
		"/event/for/today",
		middleware.LoggingMiddleware(
			s.GetEventsForToday(),
		),
	)

	// Get events for week
	mux.HandleFunc(
		"/event/for/week",
		middleware.LoggingMiddleware(
			s.GetEventsForWeek(),
		),
	)

	// Get events for month
	mux.HandleFunc(
		"/event/for/month",
		middleware.LoggingMiddleware(
			s.GetEventsForMonth(),
		),
	)

	return mux
}
