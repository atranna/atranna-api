package users

import "github.com/gin-gonic/gin"

func (h *Handler) GetSelf(c *gin.Context) {
	user, found := h.users.GetByID(c.GetInt("user_id"))
	if !found {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}
	
	user.PasswordHash = "REDACTED"

	c.JSON(200, user)
}