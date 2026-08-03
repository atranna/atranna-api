package organizationMembers

import (
	"atranna-api/src/internal/repository"
)

type Handler struct {
	organizationMembers repository.OrganizationMemberRepository
}

func NewHandler(organizationMembers repository.OrganizationMemberRepository) *Handler {
	return &Handler{organizationMembers: organizationMembers}
}
