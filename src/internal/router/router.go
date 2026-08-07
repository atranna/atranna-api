package router

import (
	"github.com/atranna/atranna-api/src/internal/config"
	"github.com/atranna/atranna-api/src/internal/database"
	"github.com/atranna/atranna-api/src/internal/handlers"
	v1_devices "github.com/atranna/atranna-api/src/internal/handlers/v1/devices"
	v1_interfaces "github.com/atranna/atranna-api/src/internal/handlers/v1/interfaces"
	v1_networks "github.com/atranna/atranna-api/src/internal/handlers/v1/networks"
	v1_organizations "github.com/atranna/atranna-api/src/internal/handlers/v1/organizations"
	organizationMembers "github.com/atranna/atranna-api/src/internal/handlers/v1/organzationMembers"
	v1_users "github.com/atranna/atranna-api/src/internal/handlers/v1/users"
	"github.com/atranna/atranna-api/src/internal/middlewares"
	"github.com/atranna/atranna-api/src/internal/repository"
	"github.com/atranna/atranna-api/src/internal/repository/memory"
	"github.com/atranna/atranna-api/src/internal/repository/postgres"
	"github.com/atranna/atranna-api/src/internal/repository/sqlite"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	if !config.Current.Debug {
		gin.SetMode(gin.ReleaseMode)
	} 
	
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	r.GET("/ping", handlers.Ping)

	var organizationRepo repository.OrganizationRepository
	var userRepo repository.UsersRepository
	var organizationMembersRepo repository.OrganizationMemberRepository

	var deviceRepo repository.DeviceRepository
	var interfaceRepo repository.InterfaceRepository
	var networkRepo repository.NetworkRepository
	
	switch config.Current.Storage.Backend {
	case "memory":
		organizationRepo = memory.NewOrganizationRepository()
		userRepo = memory.NewUserRepository()
		organizationMembersRepo = memory.NewOrganizationMemberRepository()

		deviceRepo = memory.NewDeviceRepository()
		interfaceRepo = memory.NewInterfaceRepository(deviceRepo)
		networkRepo = memory.NewNetworkRepository()
	case "postgres":
		database.Init()
		database.ApplyPostgresMigrations()

		organizationRepo = postgres.NewOrganizationRepository(database.DB)
		userRepo = postgres.NewUserRepository(database.DB)
		organizationMembersRepo = postgres.NewOrganizationMemberRepository(database.DB)

		deviceRepo = postgres.NewDeviceRepository(database.DB)
		interfaceRepo = postgres.NewInterfaceRepository(database.DB, deviceRepo)
		networkRepo = postgres.NewNetworkRepository(database.DB)
	case "sqlite":
		database.Init()
		database.ApplySQLiteMigrations()

		organizationRepo = sqlite.NewOrganizationRepository(database.DB)
		userRepo = sqlite.NewUserRepository(database.DB)
		organizationMembersRepo = sqlite.NewOrganizationMemberRepository(database.DB)

		deviceRepo = sqlite.NewDeviceRepository(database.DB)
		interfaceRepo = sqlite.NewInterfaceRepository(database.DB, deviceRepo)
		networkRepo = sqlite.NewNetworkRepository(database.DB)
	}

	organizationHandler := v1_organizations.NewHandler(organizationRepo)
	usersHandler := v1_users.NewHandler(userRepo, organizationMembersRepo)
	organizationMembersHandler := organizationMembers.NewHandler(organizationMembersRepo)

	devicesHandler := v1_devices.NewHandler(deviceRepo, interfaceRepo)
	interfacesHandler := v1_interfaces.NewHandler(interfaceRepo)
	networksHandler := v1_networks.NewHandler(networkRepo)

	apiV1 := r.Group("/api/v1")
	apiV1Protected := apiV1.Group("", middlewares.AuthMiddleware())

	// Organizations
	organizationsGroup := apiV1Protected.Group("/organizations", middlewares.BlockMasterTokenMiddleware())
	organizationsGroup.GET("", organizationHandler.GetOrganizations)
	organizationsGroup.GET("/:id", organizationHandler.GetOrganization)
	organizationsGroup.POST("", organizationHandler.Add)
	organizationsGroup.DELETE("/:id", organizationHandler.Delete)

	// Users
	usersGroup := apiV1Protected.Group("/users")
	usersGroup.GET("", usersHandler.GetUsers)
	usersGroup.GET("/:id", usersHandler.GetUser)
	usersGroup.POST("", usersHandler.Add)
	usersGroup.DELETE("/:id", usersHandler.Delete)
	usersGroup.GET("/me", middlewares.BlockMasterTokenMiddleware(), usersHandler.GetSelf)

	// Auth
	authGroup := apiV1.Group("/auth")
	authGroup.POST("/login", usersHandler.Login)

	// Organization Members
	organizationsGroup.GET("/:id/members", organizationMembersHandler.GetOrganizationMembers)
	organizationsGroup.POST("/:id/members", organizationMembersHandler.AddOrganizationMember)
	organizationsGroup.DELETE("/:id/members/:user_id", organizationMembersHandler.DeleteOrganizationMember)

	// Devices
	devicesGroup := apiV1Protected.Group("/devices", middlewares.BlockMasterTokenMiddleware())
	devicesGroup.GET("", devicesHandler.GetDevices)
	devicesGroup.GET("/:id", devicesHandler.GetDevice)
	devicesGroup.POST("", devicesHandler.Add)
	devicesGroup.DELETE("/:id", devicesHandler.Delete)

	// Interfaces
	interfacesGroup := apiV1Protected.Group("/interfaces", middlewares.BlockMasterTokenMiddleware())
	interfacesGroup.GET("", interfacesHandler.GetInterfaces)
	interfacesGroup.GET("/:id", interfacesHandler.GetInterface)
	interfacesGroup.POST("", interfacesHandler.Add)
	interfacesGroup.DELETE("/:id", interfacesHandler.Delete)

	// Networks
	networksGroup := apiV1Protected.Group("/networks", middlewares.BlockMasterTokenMiddleware())
	networksGroup.GET("", networksHandler.GetNetworks)
	networksGroup.GET("/:id", networksHandler.GetNetwork)
	networksGroup.POST("", networksHandler.Add)
	networksGroup.DELETE("/:id", networksHandler.Delete)

	return r
}
