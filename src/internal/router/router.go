package router

import (
	"atranna-api/src/internal/handlers"

	"github.com/gin-gonic/gin"
)
func New() *gin.Engine {
	r := gin.Default()
	
	r.GET("/ping", handlers.Ping)
	
	return r
}