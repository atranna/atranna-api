package users

import (
	"github.com/atranna/atranna-api/src/internal/repository"
)

type Handler struct {
	users    repository.UsersRepository
	organizationMember repository.OrganizationMemberRepository
}

func NewHandler(users repository.UsersRepository, organizationMember repository.OrganizationMemberRepository) *Handler {
	return &Handler{users: users, organizationMember: organizationMember}
}
