package calendar

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
)

// handlerUpdateEvent is the handler to update an event.
func (s *Server) handlerUpdateEvent() http.HandlerFunc {
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
		v.ID = uint64(id)

		// Update the event
		err = s.app.DI().EventService().UpdateEvent(ctx, &v)
		if err != nil {
			errorHandler(w,
				fmt.Errorf("failed to update event: %w", err),
				http.StatusInternalServerError,
			)
			return
		}

		// Write response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	}
}
