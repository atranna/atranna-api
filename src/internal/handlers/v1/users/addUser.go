package users

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/atranna/atranna-api/src/internal/helpers"
	"github.com/atranna/atranna-api/src/internal/models"
)

func (h *Handler) Add(c *gin.Context) {
	var request struct {
		Email       string `json:"email"`
		Username    string `json:"username" binding:"required"`
		Password    string `json:"password" binding:"required"`
		DisplayName string `json:"display_name"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passwordHash, err := helpers.HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newUser := models.User{
		Email:        request.Email,
		Username:     request.Username,
		PasswordHash: passwordHash,
		DisplayName:  request.DisplayName,
	}

	addedUser, err := h.users.Add(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":    addedUser.ID,
		"email": addedUser.Email,
		"username": addedUser.Username,
		"display_name": addedUser.DisplayName,
	})
}
