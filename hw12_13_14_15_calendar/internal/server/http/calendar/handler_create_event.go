package calendar

import (
	"encoding/json"
	"fmt"
	"net/http"

	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
)

// handlerGetEvent is the handler to create the event
func (s *Server) handlerCreateEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var err error

		// Read the request body
		defer r.Body.Close()
		var v eventModel.Event
		if err = json.NewDecoder(r.Body).Decode(&v); err != nil {
			errorHandler(w,
				fmt.Errorf("failed to decode body: %v", err),
				http.StatusBadRequest,
			)
			return
		}

		// Create the event
		id, err := s.app.DI().EventService().CreateEvent(ctx, &v)
		if err != nil {
			errorHandler(w,
				fmt.Errorf("failed to  create event: %v", err),
				http.StatusInternalServerError,
			)
			return
		}

		// Write response
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"id": "` + fmt.Sprint(id) + `"}`))
	}
}
