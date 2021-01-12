package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getRoutes() http.Handler {
	router := gin.Default()
	return router
}
