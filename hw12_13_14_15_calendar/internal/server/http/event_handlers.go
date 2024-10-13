package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/storage"
)

func (s *Server) handlerIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, calendar!"))
	}
}

// handlerGetEvent is the handler to get the event
func (s *Server) handlerGetEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		errorHandler(w, fmt.Errorf("not implemented"), http.StatusNotImplemented)
	}
}

// handlerGetEvent is the handler to create the event
func (s *Server) handlerCreateEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		reader, err := r.GetBody()
		if err != nil {
			errorHandler(w,
				fmt.Errorf("failed to get body: %v", err),
				http.StatusInternalServerError)
			return
		}
		defer reader.Close()

		bb := make([]byte, 0, 1024)
		_, err = reader.Read(bb)
		if err != nil {
			errorHandler(w,
				fmt.Errorf("failed to read body: %v", err),
				http.StatusInternalServerError)
			return
		}

		v := &storage.Event{}
		err = json.Unmarshal(bb, v)
		if err != nil {
			errorHandler(w,
				fmt.Errorf("failed to unmarshal body: %v", err),
				http.StatusInternalServerError)
			return
		}

		id, err := s.app.DIContainer().Postgres().CreateEvent(ctx, v)
		if err != nil {
			errorHandler(w,
				fmt.Errorf("failed to  create event: %v", err),
				http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"id": "` + id + `"}`))
	}
}

func (s *Server) handlerUpdateEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		errorHandler(w, fmt.Errorf("not implemented"), http.StatusNotImplemented)
	}
}

func (s *Server) handlerDeleteEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		errorHandler(w, fmt.Errorf("not implemented"), http.StatusNotImplemented)
	}
}

// errorHandler is the handler to handle errors
func errorHandler(w http.ResponseWriter, err error, httpCode int) {
	logger.Error(err.Error())
	http.Error(w, err.Error(), httpCode)
}
