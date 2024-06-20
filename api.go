package goapi

import (
	"log"
	"net/http"
)

type Server struct {
	Addres string
	Router *Router
	Logger *logger
}

func NewServer(addres string) *Server {
	return &Server{Addres: addres, Router: &Router{}, Logger: &logger{}}
}

func (s *Server) Serve() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Info(r.URL.Path)

		rc, err := s.Router.FindByPath(r.URL.Path, r.Method)

		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte(err.Error()))
			return
		}

		if rc == nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("{\"error\": \"Not found route config\"}"))
		} else {
			rc.HandlerFunc(w, r)
		}
	})

	s.Logger.Info("Listening at: http://" + s.Addres)

	if err := http.ListenAndServe(s.Addres, nil); err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
