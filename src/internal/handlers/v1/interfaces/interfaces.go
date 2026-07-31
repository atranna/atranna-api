package interfaces

import (
	"atranna-api/src/internal/repository"
)

type Handler struct {
	interfaces repository.InterfaceRepository
}

func NewHandler(interfaces repository.InterfaceRepository) *Handler {
	return &Handler{interfaces: interfaces}
}
