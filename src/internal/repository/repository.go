package repository

import (
	"atranna-api/src/internal/models"
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
