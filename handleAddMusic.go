package main

import (
	"net/http"

	"github.com/breaking-fullstack/forever-server/entity"
	"github.com/gin-gonic/gin"
)

func (s *Server) handleAddMusic(c *gin.Context) {
	userID := c.GetString("UID")

	var newMusic entity.Music
	err := c.BindJSON(&newMusic)
	if err != nil {
		c.Error(err)
		return
	}

	err = s.musicService.Save(c.Request.Context(), userID, newMusic)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, newMusic)
}
