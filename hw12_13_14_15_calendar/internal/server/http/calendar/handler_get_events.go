package calendar

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/filter/event"
)

// handlerGetEvents is the handler to get the event list.
func (s *Server) handlerGetEvents() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		events, err := s.app.DI().EventService().GetEvents(ctx, event.Filter{})
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
