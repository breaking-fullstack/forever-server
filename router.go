package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) getRoutes() http.Handler {
	router := gin.Default()

	router.GET("/", s.handleRoot)

	return router
}
