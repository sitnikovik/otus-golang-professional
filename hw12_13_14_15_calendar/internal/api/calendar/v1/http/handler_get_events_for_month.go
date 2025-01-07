package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetEventsForMonth is the handler to get the events for current month.
func (s *Server) GetEventsForMonth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Get the events
		events, err := s.eventService.GetEventsForMonth(ctx)
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
