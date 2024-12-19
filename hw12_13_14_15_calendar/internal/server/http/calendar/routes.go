package calendar

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/server/http/middleware"
)

// routes defines the routes of the HTTP server
func (s *Server) routes() http.Handler {
	mux := mux.NewRouter()

	// Index
	mux.HandleFunc(
		"/",
		middleware.LoggingMiddleware(
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Hello, calendar!"))
			},
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

	return mux
}
