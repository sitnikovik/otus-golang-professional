package http

import "net/http"

func (s *Server) handlerHello() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, calendar!"))
	}
}
