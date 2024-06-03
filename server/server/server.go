package server

import (
	"net/http"
	"plan-O/modules/crm"
	"plan-O/repository"

	"github.com/gorilla/mux"
)

type Server struct {
	crm *crm.Service
	r   *mux.Router
}

func New(db repository.Repo) *Server {
	s := &Server{
		crm: crm.NewService(db),
		r:   mux.NewRouter(),
	}

	s.r.PathPrefix("/api/crm").Handler(http.StripPrefix("/api/crm", s.crm))

	return s
}

func (s *Server) Run(addr string) error {
	return http.ListenAndServe(addr, s.r)
}
