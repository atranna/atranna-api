package users

import (
	"atranna-api/src/internal/repository"
)

type Handler struct {
	users    repository.UsersRepository
}

func NewHandler(users repository.UsersRepository) *Handler {
	return &Handler{users: users}
}
