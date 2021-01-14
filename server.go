package main

import (
	"context"
	"net/http"
	"time"

	"github.com/breaking-fullstack/forever-server/service"
	"github.com/breaking-fullstack/forever-server/verifier"
)

//Server defines a wrapper around http.Server.
//It holds run dependencies useful in handlers.
type Server struct {
	musicService  service.Music
	tokenVerifier verifier.Verifier
	*http.Server
}

//NewServer returns a new Server struct holding dependencies
func NewServer(addr string, ms service.Music, tv verifier.Verifier) *Server {
	srv := &Server{
		musicService:  ms,
		tokenVerifier: tv,
		Server: &http.Server{
			WriteTimeout: 5 * time.Second,
			ReadTimeout:  5 * time.Second,
			IdleTimeout:  5 * time.Second,
			Addr:         addr,
		},
	}
	srv.Server.Handler = srv.getRoutes()
	return srv
}

//Start runs the server
func (s *Server) Start() error {
	return s.ListenAndServe()
}

//Stop shutsdown the server
func (s *Server) Stop() error {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.Shutdown(shutdownCtx)
}
