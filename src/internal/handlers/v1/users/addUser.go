package users

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"atranna-api/src/internal/models"
)

func (h *Handler) Add(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	addedUser, err := h.users.Add(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, addedUser)
}
