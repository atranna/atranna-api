package organizations

import (
	"github.com/atranna/atranna-api/src/internal/repository"
)

type Handler struct {
	organizations    repository.OrganizationRepository
	organizationMembers repository.OrganizationMemberRepository
}

func NewHandler(organizations repository.OrganizationRepository, organizationMembers repository.OrganizationMemberRepository) *Handler {
	return &Handler{organizations: organizations, organizationMembers: organizationMembers}
}
