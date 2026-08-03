package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, h.users.GetAll())
}
