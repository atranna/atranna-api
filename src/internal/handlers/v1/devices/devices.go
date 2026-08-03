package devices

import (
	"github.com/atranna/atranna-api/src/internal/repository"
)

type Handler struct {
	devices    repository.DeviceRepository
	interfaces repository.InterfaceRepository
}

func NewHandler(devices repository.DeviceRepository, interfaces repository.InterfaceRepository) *Handler {
	return &Handler{devices: devices, interfaces: interfaces}
}
