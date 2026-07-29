package router

import (
	"atranna-api/src/internal/handlers"
	v1_devices "atranna-api/src/internal/handlers/v1/devices"
	v1_interfaces "atranna-api/src/internal/handlers/v1/devices/interfaces"
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

	// API v1 routes

	// Devices
	r.GET("/api/v1/devices", v1_devices.GetDevices)
	r.GET("/api/v1/devices/:id", v1_devices.GetDevice)
	r.POST("/api/v1/devices/", v1_devices.AddDevice)
	r.DELETE("/api/v1/devices/:id", v1_devices.DeleteDevice)

	// Interfaces

	r.GET("/api/v1/interfaces/", v1_interfaces.GetInterfaces)
	r.GET("/api/v1/interfaces/:interface_id", v1_interfaces.GetInterface)
	r.DELETE("/api/v1/interfaces/:interface_id", v1_interfaces.DeleteInterface)

	// Interfaces with device ID
	r.GET("/api/v1/devices/:id/interfaces/", v1_interfaces.GetInterfacesByDeviceID)
	r.POST("/api/v1/devices/:id/interfaces/", v1_interfaces.AddInterface)
	
	
	return r
}