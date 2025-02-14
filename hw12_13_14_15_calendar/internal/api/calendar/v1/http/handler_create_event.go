package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
)

// handlerGetEvent is the handler to create the event.
func (s *Server) handlerCreateEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var err error

		// Read the request body
		defer r.Body.Close()
		bb, err := io.ReadAll(r.Body)
		if err != nil {
			errorHandler(w,
				fmt.Errorf("failed to read body: %w", err),
				http.StatusInternalServerError,
			)
			return
		}

		// Unmarshal the request body to the event model
		var v eventModel.Event
		err = json.Unmarshal(bb, &v)
		if err != nil {
			errorHandler(w,
				fmt.Errorf("failed to unmarshal body: %w", err),
				http.StatusInternalServerError,
			)
			return
		}

		// Create the event
		id, err := s.eventService.CreateEvent(ctx, &v)
		if err != nil {
			errorHandler(w,
				fmt.Errorf("failed to  create event: %w", err),
				http.StatusInternalServerError,
			)
			return
		}

		// Write response
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"id": "` + fmt.Sprint(id) + `"}`))
	}
}
