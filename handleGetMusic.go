package main

import "github.com/gin-gonic/gin"

func (s *Server) handleGetMusic(c *gin.Context) {
	userID := c.GetString("UID")
	musicList, err := s.musicService.GetAll(c.Request.Context(), userID)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, musicList)
}
