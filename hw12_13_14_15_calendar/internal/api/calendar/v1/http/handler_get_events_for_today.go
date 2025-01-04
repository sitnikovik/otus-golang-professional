package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetEventsForToday is the handler to get the events for today.
func (s *Server) GetEventsForToday() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Get the events
		events, err := s.eventService.GetEventsForToday(ctx)
		if err != nil {
			errorHandler(
				w,
				err,
				http.StatusInternalServerError,
			)
			return
		}

		// Write response
		bb, err := json.Marshal(events)
		if err != nil {
			errorHandler(
				w,
				fmt.Errorf("failed to marshal events: %w", err),
				http.StatusInternalServerError,
			)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(bb)
	}
}
