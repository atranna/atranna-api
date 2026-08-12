package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUsers(c *gin.Context) {
	allUsers := h.users.GetAll()
	type user struct {
		ID    int    `json:"id"`
		Username  string `json:"username"`
	}

	users := make([]user, len(allUsers))
	for i := range allUsers {
		users[i] = user{
			ID:    allUsers[i].ID,
			Username:  allUsers[i].Username,
		}
	}

	c.JSON(http.StatusOK, users)
}
