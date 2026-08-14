package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) GetUsers(c *gin.Context) {
	allUsers := h.users.GetAll()

	users := make([]user, 0, len(allUsers))
	for _, u := range allUsers {
		users = append(users, user{ID: u.ID, Username: u.Username})
	}

	c.JSON(http.StatusOK, users)
}
