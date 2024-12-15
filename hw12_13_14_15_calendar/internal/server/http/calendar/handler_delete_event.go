package calendar

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) handlerDeleteEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Get the event ID from URI path param
		vars := mux.Vars(r)
		id := vars["id"]
		if id == "" {
			errorHandler(
				w,
				fmt.Errorf("event ID is required"),
				http.StatusBadRequest,
			)
			return
		}

		// Delete the event
		err := s.app.DI().EventService().DeleteEvent(ctx, id)
		if err != nil {
			errorHandler(
				w,
				fmt.Errorf("failed to delete event: %v", err),
				http.StatusInternalServerError,
			)
			return
		}

		// Write response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	}
}
