package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"plan-O/modules/crm"
	"plan-O/repository"
	"time"

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
	srvHTTP := &http.Server{
		ReadTimeout: 10 * time.Second,
		Handler:     s.r,
		Addr:        addr,
	}
	listenerHTTP, err := net.Listen("tcp4", addr)
	if err != nil {
		return fmt.Errorf("ошибка сервера HTTP: %w", err)
	}
	go func() {
		err = srvHTTP.Serve(listenerHTTP)
		if err != nil {
			log.Println(err, "ошибка сервера HTTP")
		}
	}()
	return nil
}
