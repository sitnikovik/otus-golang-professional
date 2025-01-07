package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetEventsBeforeDays is the handler to get the events before days.
func (s *Server) GetEventsBeforeDays() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Get days from URI path param
		vars := mux.Vars(r)
		days, _ := strconv.Atoi(vars["days"])
		if days <= 0 {
			errorHandler(
				w,
				fmt.Errorf("days is empty or invalid"),
				http.StatusBadRequest,
			)
			return
		}

		// Get the events
		events, err := s.eventService.GetEventsBeforeDays(ctx, uint32(days))
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
