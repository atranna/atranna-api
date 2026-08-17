package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func (h *Handler) ValidateToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "valid"})
}
