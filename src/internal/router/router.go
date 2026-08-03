package router

import (
	"atranna-api/src/internal/config"
	"atranna-api/src/internal/database"
	"atranna-api/src/internal/handlers"
	v1_devices "atranna-api/src/internal/handlers/v1/devices"
	v1_interfaces "atranna-api/src/internal/handlers/v1/interfaces"
	v1_networks "atranna-api/src/internal/handlers/v1/networks"
	v1_organizations "atranna-api/src/internal/handlers/v1/organizations"
	"atranna-api/src/internal/middlewares"
	"atranna-api/src/internal/repository"
	"atranna-api/src/internal/repository/memory"
	"atranna-api/src/internal/repository/postgres"
	"atranna-api/src/internal/repository/sqlite"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	if !config.Current.Debug {
		gin.SetMode(gin.ReleaseMode)
	} 
	
	r := gin.Default()

	r.GET("/ping", handlers.Ping)

	var organizationRepo repository.OrganizationRepository

	var deviceRepo repository.DeviceRepository
	var interfaceRepo repository.InterfaceRepository
	var networkRepo repository.NetworkRepository
	
	switch config.Current.Storage.Backend {
	case "memory":
		organizationRepo = memory.NewOrganizationRepository()

		deviceRepo = memory.NewDeviceRepository()
		interfaceRepo = memory.NewInterfaceRepository(deviceRepo)
		networkRepo = memory.NewNetworkRepository()
	case "postgres":
		database.Init()
		database.ApplyPostgresMigrations()

		organizationRepo = postgres.NewOrganizationRepository(database.DB)

		deviceRepo = postgres.NewDeviceRepository(database.DB)
		interfaceRepo = postgres.NewInterfaceRepository(database.DB, deviceRepo)
		networkRepo = postgres.NewNetworkRepository(database.DB)
	case "sqlite":
		database.Init()
		database.ApplySQLiteMigrations()

		organizationRepo = sqlite.NewOrganizationRepository(database.DB)

		deviceRepo = sqlite.NewDeviceRepository(database.DB)
		interfaceRepo = sqlite.NewInterfaceRepository(database.DB, deviceRepo)
		networkRepo = sqlite.NewNetworkRepository(database.DB)
	}

	organizationHandler := v1_organizations.NewHandler(organizationRepo)

	devicesHandler := v1_devices.NewHandler(deviceRepo, interfaceRepo)
	interfacesHandler := v1_interfaces.NewHandler(interfaceRepo)
	networksHandler := v1_networks.NewHandler(networkRepo)

	apiV1 := r.Group("/api/v1", middlewares.AuthMiddleware())

	// Organizations
	apiV1.GET("/organizations", organizationHandler.GetOrganizations)
	apiV1.GET("/organizations/:id", organizationHandler.GetOrganization)
	apiV1.POST("/organizations", organizationHandler.Add)
	apiV1.DELETE("/organizations/:id", organizationHandler.Delete)

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
