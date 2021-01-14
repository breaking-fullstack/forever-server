package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) handleRoot(c *gin.Context) {
	c.String(http.StatusOK, "https://github.com/breaking-fullstack/forever")
}
