package router

import (
	"atranna-api/src/internal/config"
	"atranna-api/src/internal/database"
	"atranna-api/src/internal/handlers"
	v1_devices "atranna-api/src/internal/handlers/v1/devices"
	v1_interfaces "atranna-api/src/internal/handlers/v1/interfaces"
	v1_networks "atranna-api/src/internal/handlers/v1/networks"
	"atranna-api/src/internal/middlewares"
	"atranna-api/src/internal/repository"
	"atranna-api/src/internal/repository/memory"
	"atranna-api/src/internal/repository/sqlite"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	if !config.Current.Debug {
		gin.SetMode(gin.ReleaseMode)
	} 
	
	r := gin.Default()

	r.GET("/ping", handlers.Ping)

	var deviceRepo repository.DeviceRepository
	var interfaceRepo repository.InterfaceRepository
	var networkRepo repository.NetworkRepository
	
	switch config.Current.Storage.Backend {
	case "memory":
		deviceRepo = memory.NewDeviceRepository()
		interfaceRepo = memory.NewInterfaceRepository(deviceRepo)
		networkRepo = memory.NewNetworkRepository()
	case "postgres":
		database.Init()
		database.ApplyPostgresMigrations()

		// deviceRepo = postgres.NewDeviceRepository()
		// interfaceRepo = postgres.NewInterfaceRepository(deviceRepo)
		// networkRepo = postgres.NewNetworkRepository()
	case "sqlite":
		database.Init()
		database.ApplySQLiteMigrations()

		deviceRepo = sqlite.NewDeviceRepository(database.DB)
		interfaceRepo = sqlite.NewInterfaceRepository(database.DB, deviceRepo)
		networkRepo = sqlite.NewNetworkRepository(database.DB)
	}

	devicesHandler := v1_devices.NewHandler(deviceRepo, interfaceRepo)
	interfacesHandler := v1_interfaces.NewHandler(interfaceRepo)
	networksHandler := v1_networks.NewHandler(networkRepo)

	apiV1 := r.Group("/api/v1", middlewares.AuthMiddleware())

	// Devices
	apiV1.GET("/devices", devicesHandler.GetDevices)
	apiV1.GET("/devices/:id", devicesHandler.GetDevice)
	apiV1.POST("/devices", devicesHandler.Add)
	apiV1.DELETE("/devices/:id", devicesHandler.Delete)

	// Interfaces
	apiV1.GET("/interfaces", interfacesHandler.GetInterfaces)
	apiV1.GET("/interfaces/:id", interfacesHandler.GetInterface)
	apiV1.POST("/interfaces", interfacesHandler.Add)
	apiV1.DELETE("/interfaces/:id", interfacesHandler.Delete)

	// Networks
	apiV1.GET("/networks", networksHandler.GetNetworks)
	apiV1.GET("/networks/:id", networksHandler.GetNetwork)
	apiV1.POST("/networks", networksHandler.Add)
	apiV1.DELETE("/networks/:id", networksHandler.Delete)

	return r
}
