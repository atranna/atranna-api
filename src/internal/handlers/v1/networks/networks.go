package networks

import (
	"github.com/atranna/atranna-api/src/internal/repository"
)

type Handler struct {
	networks repository.NetworkRepository
}

func NewHandler(networks repository.NetworkRepository) *Handler {
	return &Handler{networks: networks}
}
