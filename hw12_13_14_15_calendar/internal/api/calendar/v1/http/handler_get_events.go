package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/filter/event"
)

// handlerGetEvents is the handler to get the event list.
func (s *Server) handlerGetEvents() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Get the events
		events, err := s.eventService.GetEvents(ctx, parseFilter(r))
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
				fmt.Errorf("failed to marshal event: %w", err),
				http.StatusInternalServerError,
			)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(bb)
	}
}

// parseFilter parses the filter from the request query params.
func parseFilter(r *http.Request) event.Filter {
	var filter event.Filter

	ids := r.URL.Query()["ids"]
	if len(ids) > 0 {
		for _, id := range ids {
			idUint, _ := strconv.ParseUint(id, 10, 64)
			if idUint > 0 {
				filter.IDs = append(filter.IDs, idUint)
			}
		}
	}

	limit := r.URL.Query().Get("limit")
	if limit != "" {
		limitUint, _ := strconv.ParseUint(limit, 10, 64)
		if limitUint > 0 {
			filter.Limit = limitUint
		}
	}

	return filter
}
