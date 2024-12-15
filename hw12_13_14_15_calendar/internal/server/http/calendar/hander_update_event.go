package calendar

import (
	"encoding/json"
	"fmt"
	"net/http"

	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
)

// handlerUpdateEvent is the handler to update an event
func (s *Server) handlerUpdateEvent() http.HandlerFunc {
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
		event := &eventModel.Event{}
		err = json.Unmarshal(bb, event)
		if err != nil {
			errorHandler(w,
				fmt.Errorf("failed to unmarshal body: %v", err),
				http.StatusInternalServerError,
			)
			return
		}

		// Update the event
		err = s.app.DI().EventService().UpdateEvent(ctx, event)
		if err != nil {
			errorHandler(w,
				fmt.Errorf("failed to update event: %v", err),
				http.StatusInternalServerError,
			)
			return
		}

		// Write response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	}
}
