package main

import (
	"net/http"
	"time"

	"github.com/breaking-fullstack/forever-server/service"
)

//Server defines a wrapper around http.Server.
//It holds run dependencies useful in handlers.
type Server struct {
	musicService service.MusicService
	*http.Server
}

//NewServer returns a new Server struct holding dependencies
func NewServer(addr string, ms service.MusicService) *Server {
	srv := &Server{
		musicService: ms,
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
