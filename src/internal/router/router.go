package router

import (
	"atranna-api/src/internal/handlers"
	v1 "atranna-api/src/internal/handlers/v1"
	"os"

	"github.com/gin-gonic/gin"
)
func New() *gin.Engine {
	if os.Getenv("DEBUG") == "true" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	r.GET("/ping", handlers.Ping)
	r.GET("/api/v1/devices", v1.GetDevices)
	r.GET("/api/v1/devices/:id", v1.GetDevice)
	r.POST("/api/v1/devices/", v1.AddDevice)
	r.DELETE("/api/v1/devices/:id", v1.DeleteDevice)
	
	return r
}