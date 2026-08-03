package users

import (
	"atranna-api/src/internal/helpers"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	var request struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	users := h.users.GetAll()

	for _, user := range users {
		if user.Username == request.Username {
			if !helpers.CheckPasswordHash(request.Password, user.PasswordHash) {
				c.JSON(401, gin.H{"error": "Invalid username or password"})
				return
			}

			token, err := helpers.GenerateJWT(user.ID)
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to generate token"})
				return
			}

			c.JSON(200, gin.H{"token": token})
			return
		}
	}
	c.JSON(401, gin.H{"error": "Invalid username or password"})
}