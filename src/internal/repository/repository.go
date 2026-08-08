package repository

import (
	"github.com/atranna/atranna-api/src/internal/models"
)

type DeviceRepository interface {
	Add(device models.Device) (models.Device, error)
	GetByID(id int) (models.Device, bool)
	GetAll() []models.Device
	Delete(id int) (models.Device, bool)
}

type InterfaceRepository interface {
	Add(interf models.Interface) (models.Interface, error)
	GetByID(id int) (models.Interface, bool)
	GetAll() []models.Interface
	GetByDeviceID(deviceID int) []models.Interface
	Delete(id int) (models.Interface, bool)
	DeleteByDeviceID(deviceID int) error
}

type NetworkRepository interface {
	Add(network models.Network) (models.Network, error)
	GetByID(id int) (models.Network, bool)
	GetAll() []models.Network
	Delete(id int) (models.Network, bool)
}

type OrganizationRepository interface {
	Add(organization models.Organization) (models.Organization, error)
	GetByID(id int) (models.Organization, bool)
	GetAll() []models.Organization
	Delete(id int) (models.Organization, bool)
}

type UsersRepository interface {
	Add(user models.User) (models.User, error)
	GetByID(id int) (models.User, bool)
	GetByUsername(username string) (models.User, bool)
	GetAll() []models.User
	Delete(id int) (models.User, bool)
}

type OrganizationMemberRepository interface {
	Add(member models.OrganizationMember) (models.OrganizationMember, error)
	GetByOrganizationID(organizationID int) ([]models.OrganizationMember, error)
	GetByUserID(userID int) []models.OrganizationMember
	Delete(organizationID int, userID int) (models.OrganizationMember, bool)
	GetRole(organizationID int, userID int) (string, bool)
}