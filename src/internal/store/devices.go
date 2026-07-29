package store

import "atranna-api/src/internal/models"

var nextID int = 1

var Devices = []models.Device{}

func AddDevice(device models.Device) models.Device {
	device.ID = nextID
	nextID++
	Devices = append(Devices, device)
	return device
}