package memory

import (
	"fmt"

	"github.com/atranna/atranna-api/src/internal/models"
	"github.com/atranna/atranna-api/src/internal/repository"
)

type InterfaceRepository struct {
	interfaces []models.Interface
	nextID     int
	devices    repository.DeviceRepository
}

func NewInterfaceRepository(devices repository.DeviceRepository) *InterfaceRepository {
	return &InterfaceRepository{devices: devices}
}

func (r *InterfaceRepository) Add(interf models.Interface) (models.Interface, error) {
	if _, exists := r.devices.GetByID(interf.DeviceID); !exists {
		return models.Interface{}, fmt.Errorf("device with ID %d does not exist", interf.DeviceID)
	}
	r.nextID++
	interf.ID = r.nextID
	r.interfaces = append(r.interfaces, interf)
	return interf, nil
}

func (r *InterfaceRepository) GetByID(id int) (models.Interface, bool) {
	for _, interf := range r.interfaces {
		if interf.ID == id {
			return interf, true
		}
	}
	return models.Interface{}, false
}

func (r *InterfaceRepository) GetAll() []models.Interface {
	return r.interfaces
}

func (r *InterfaceRepository) GetByDeviceID(deviceID int) []models.Interface {
	result := []models.Interface{}
	for _, interf := range r.interfaces {
		if interf.DeviceID == deviceID {
			result = append(result, interf)
		}
	}
	return result
}

func (r *InterfaceRepository) Delete(id int) (models.Interface, bool) {
	for i, interf := range r.interfaces {
		if interf.ID == id {
			r.interfaces = append(r.interfaces[:i], r.interfaces[i+1:]...)
			return interf, true
		}
	}
	return models.Interface{}, false
}

func (r *InterfaceRepository) DeleteByDeviceID(deviceID int) error {
	filtered := make([]models.Interface, 0, len(r.interfaces))
	for _, interf := range r.interfaces {
		if interf.DeviceID != deviceID {
			filtered = append(filtered, interf)
		}
	}
	r.interfaces = filtered
	return nil
}
