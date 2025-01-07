package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Server) handlerDeleteEvent() http.HandlerFunc {
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

		// Delete the event
		err := s.eventService.DeleteEvent(ctx, uint64(id))
		if err != nil {
			errorHandler(
				w,
				fmt.Errorf("failed to delete event: %w", err),
				http.StatusInternalServerError,
			)
			return
		}

		// Write response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	}
}
