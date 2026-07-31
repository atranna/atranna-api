package store

import (
	"atranna-api/src/internal/models"
)

var nextDeviceID int = 1

var Devices = []models.Device{}

func AddDevice(device models.Device) models.Device {
	device.ID = nextDeviceID
	nextDeviceID++
	Devices = append(Devices, device)
	return device
}

func GetDeviceByID(id int) (models.Device, bool) {
	for _, device := range Devices {
		if device.ID == id {
			return device, true
		}
	}
	return models.Device{}, false
}

func DeleteDevice(id int) (models.Device, bool) {
	for i, device := range Devices {
		if device.ID == id {
			Devices = append(Devices[:i], Devices[i+1:]...)
			DeleteInterfacesByDeviceID(id)
			return device, true
		}
	}

	return models.Device{}, false
}

func DeviceExists(id int) bool {
	for _, device := range Devices {
		if device.ID == id {
			return true
		}
	}
	return false
}