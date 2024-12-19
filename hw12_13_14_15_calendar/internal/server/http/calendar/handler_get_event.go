package calendar

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// handlerGetEvent is the handler to get the event.
func (s *Server) handlerGetEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Get the event ID from URI path param
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		if id <= 0 {
			errorHandler(
				w,
				fmt.Errorf("event ID is empty or invalid"),
				http.StatusBadRequest,
			)
			return
		}

		// Get the event
		event, err := s.app.DI().EventService().GetEvent(ctx, uint64(id))
		if err != nil {
			errorHandler(
				w,
				err,
				http.StatusInternalServerError,
			)
			return
		}

		// Write response
		bb, err := json.Marshal(event)
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
