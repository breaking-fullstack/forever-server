package main

import (
	"net/http"

	"github.com/breaking-fullstack/forever-server/middleware"
	"github.com/gin-gonic/gin"
)

func (s *Server) getRoutes() http.Handler {
	router := gin.Default()

	router.GET("/", s.handleRoot)

	//musicService
	msRouter := router.Group("/music", middleware.Auth(s.tokenVerifier))
	{
		msRouter.GET("", s.handleGetMusic)
	}

	return router
}
