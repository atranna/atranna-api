package networks

import (
	"atranna-api/src/internal/helpers"
	"atranna-api/src/internal/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNetworks(c *gin.Context) {
	if !helpers.CheckAuthorization(c.GetHeader("Authorization")) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	c.JSON(http.StatusOK, store.Networks)
}
