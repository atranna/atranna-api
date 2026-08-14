package router

import (
	"github.com/atranna/atranna-api/src/internal/auth"
	"github.com/atranna/atranna-api/src/internal/config"
	"github.com/atranna/atranna-api/src/internal/database"
	"github.com/atranna/atranna-api/src/internal/handlers"
	v1_devices "github.com/atranna/atranna-api/src/internal/handlers/v1/devices"
	v1_interfaces "github.com/atranna/atranna-api/src/internal/handlers/v1/interfaces"
	v1_networks "github.com/atranna/atranna-api/src/internal/handlers/v1/networks"
	v1_organization_members "github.com/atranna/atranna-api/src/internal/handlers/v1/organization_members"
	v1_organizations "github.com/atranna/atranna-api/src/internal/handlers/v1/organizations"
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

	organizationHandler := v1_organizations.NewHandler(organizationRepo, organizationMembersRepo)
	usersHandler := v1_users.NewHandler(userRepo, organizationMembersRepo)
	organizationMembersHandler := v1_organization_members.NewHandler(organizationMembersRepo)

	devicesHandler := v1_devices.NewHandler(deviceRepo, interfaceRepo)
	interfacesHandler := v1_interfaces.NewHandler(interfaceRepo)
	networksHandler := v1_networks.NewHandler(networkRepo)

	apiV1 := r.Group("/api/v1")
	apiV1Protected := apiV1.Group("", middlewares.AuthenticationMiddleware())

	// Organizations
	organizationsGroup := apiV1Protected.Group("/organizations", middlewares.BlockMasterTokenMiddleware())
	organizationsGroup.GET("", organizationHandler.GetOrganizations)
	organizationsGroup.GET("/:id", organizationHandler.GetOrganization)
	organizationsGroup.POST("", organizationHandler.Add)
	organizationsGroup.DELETE("/:id", middlewares.RequireOrganizationPermissionMiddleware(organizationMembersRepo, auth.OrganizationWrite), organizationHandler.Delete)

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
	organizationMembersGroup := apiV1Protected.Group("/organization-members", middlewares.BlockMasterTokenMiddleware(), middlewares.AuthorizationMiddleware(organizationMembersRepo))
	organizationMembersGroup.GET("", middlewares.RequirePermissionMiddleware(auth.OrganizationMembersRead), organizationMembersHandler.GetOrganizationMembers)
	organizationMembersGroup.GET("/:user_id", middlewares.RequirePermissionMiddleware(auth.OrganizationMembersRead), organizationMembersHandler.GetOrganizationMember)
	organizationMembersGroup.POST("", middlewares.RequirePermissionMiddleware(auth.OrganizationMembersWrite), organizationMembersHandler.AddOrganizationMember)
	organizationMembersGroup.DELETE("/:user_id", middlewares.RequirePermissionMiddleware(auth.OrganizationMembersWrite), organizationMembersHandler.DeleteOrganizationMember)

	// Devices
	devicesGroup := apiV1Protected.Group("/devices", middlewares.BlockMasterTokenMiddleware(), middlewares.AuthorizationMiddleware(organizationMembersRepo))
	devicesGroup.GET("", middlewares.RequirePermissionMiddleware(auth.DevicesRead), devicesHandler.GetDevices)
	devicesGroup.GET("/:id", middlewares.RequirePermissionMiddleware(auth.DevicesRead), devicesHandler.GetDevice)
	devicesGroup.POST("", middlewares.RequirePermissionMiddleware(auth.DevicesWrite), devicesHandler.Add)
	devicesGroup.DELETE("/:id", middlewares.RequirePermissionMiddleware(auth.DevicesWrite), devicesHandler.Delete)

	// Interfaces
	interfacesGroup := apiV1Protected.Group("/interfaces", middlewares.BlockMasterTokenMiddleware(), middlewares.AuthorizationMiddleware(organizationMembersRepo))
	interfacesGroup.GET("", middlewares.RequirePermissionMiddleware(auth.InterfacesRead), interfacesHandler.GetInterfaces)
	interfacesGroup.GET("/:id", middlewares.RequirePermissionMiddleware(auth.InterfacesRead), interfacesHandler.GetInterface)
	interfacesGroup.POST("", middlewares.RequirePermissionMiddleware(auth.InterfacesWrite), interfacesHandler.Add)
	interfacesGroup.DELETE("/:id", middlewares.RequirePermissionMiddleware(auth.InterfacesWrite), interfacesHandler.Delete)

	// Networks
	networksGroup := apiV1Protected.Group("/networks", middlewares.BlockMasterTokenMiddleware(), middlewares.AuthorizationMiddleware(organizationMembersRepo))
	networksGroup.GET("", middlewares.RequirePermissionMiddleware(auth.NetworksRead), networksHandler.GetNetworks)
	networksGroup.GET("/:id", middlewares.RequirePermissionMiddleware(auth.NetworksRead), networksHandler.GetNetwork)
	networksGroup.POST("", middlewares.RequirePermissionMiddleware(auth.NetworksWrite), networksHandler.Add)
	networksGroup.DELETE("/:id", middlewares.RequirePermissionMiddleware(auth.NetworksWrite), networksHandler.Delete)

	return r
}
