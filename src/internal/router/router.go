package router

import (
	"atranna-api/src/internal/handlers"
	v1_devices "atranna-api/src/internal/handlers/v1/devices"
	v1_interfaces "atranna-api/src/internal/handlers/v1/devices/interfaces"
	v1_networks "atranna-api/src/internal/handlers/v1/networks"
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

	apiV1 := r.Group("/api/v1")

	// Devices
	apiV1.GET("/devices", v1_devices.GetDevices)
	apiV1.GET("/devices/:id", v1_devices.GetDevice)
	apiV1.POST("/devices", v1_devices.AddDevice)
	apiV1.DELETE("/devices/:id", v1_devices.DeleteDevice)

	// Interfaces
	apiV1.GET("/interfaces", v1_interfaces.GetInterfaces)
	apiV1.GET("/interfaces/:id", v1_interfaces.GetInterface)
	apiV1.POST("/interfaces", v1_interfaces.AddInterface)
	apiV1.DELETE("/interfaces/:id", v1_interfaces.DeleteInterface)

	// Networks
	apiV1.GET("/networks", v1_networks.GetNetworks)
	apiV1.GET("/networks/:id", v1_networks.GetNetwork)
	apiV1.POST("/networks", v1_networks.AddNetwork)
	apiV1.DELETE("/networks/:id", v1_networks.DeleteNetwork)

	return r
}
