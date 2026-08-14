package memory

import (
	"github.com/atranna/atranna-api/src/internal/models"
)

type DeviceRepository struct {
	devices []models.Device
	nextID  int
}

func NewDeviceRepository() *DeviceRepository {
	return &DeviceRepository{}
}

func (r *DeviceRepository) Add(device models.Device) (models.Device, error) {
	r.nextID++
	device.ID = r.nextID
	r.devices = append(r.devices, device)
	return device, nil
}

func (r *DeviceRepository) GetByID(id int) (models.Device, bool) {
	for _, device := range r.devices {
		if device.ID == id {
			return device, true
		}
	}
	return models.Device{}, false
}

func (r *DeviceRepository) GetAll(orgID int) []models.Device {
	var devices []models.Device
	for _, device := range r.devices {
		if device.OrgID == orgID {
			devices = append(devices, device)
		}
	}
	return devices
}

func (r *DeviceRepository) Delete(id int) (models.Device, bool) {
	for i, device := range r.devices {
		if device.ID == id {
			r.devices = append(r.devices[:i], r.devices[i+1:]...)
			return device, true
		}
	}
	return models.Device{}, false
}
