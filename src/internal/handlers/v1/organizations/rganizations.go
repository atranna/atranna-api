package organizations

import (
	"github.com/atranna/atranna-api/src/internal/repository"
)

type Handler struct {
	organizations    repository.OrganizationRepository
}

func NewHandler(organizations repository.OrganizationRepository) *Handler {
	return &Handler{organizations: organizations}
}
