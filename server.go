package main

import (
	"log"
	"net/http"

	"github.com/mannanmcc/helloworld/config"
)

type Server struct {
	config *config.Config
}

func newServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (s *Server) Run(handler http.Handler) {
	log.Println("app launching on port:::::" + s.config.Port)
	http.ListenAndServe(":"+s.config.Port, handler)
}
