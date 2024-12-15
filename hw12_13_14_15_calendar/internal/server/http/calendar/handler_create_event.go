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

		// Read the request body
		reader, err := r.GetBody()
		if err != nil {
			errorHandler(w,
				fmt.Errorf("failed to get body: %v", err),
				http.StatusInternalServerError,
			)
			return
		}
		defer reader.Close()
		bb := make([]byte, 0, 1024)
		_, err = reader.Read(bb)
		if err != nil {
			errorHandler(w,
				fmt.Errorf("failed to read body: %v", err),
				http.StatusInternalServerError,
			)
			return
		}

		// Unmarshal the request body to the event model
		v := &eventModel.Event{}
		err = json.Unmarshal(bb, v)
		if err != nil {
			errorHandler(w,
				fmt.Errorf("failed to unmarshal body: %v", err),
				http.StatusInternalServerError,
			)
			return
		}

		// Create the event
		id, err := s.app.DI().EventService().CreateEvent(ctx, v)
		if err != nil {
			errorHandler(w,
				fmt.Errorf("failed to  create event: %v", err),
				http.StatusInternalServerError,
			)
			return
		}

		// Write response
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"id": "` + id + `"}`))
	}
}
