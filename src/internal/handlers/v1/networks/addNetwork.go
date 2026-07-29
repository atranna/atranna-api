package networks

import (
	"atranna-api/src/internal/helpers"
	"atranna-api/src/internal/models"
	"atranna-api/src/internal/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddNetwork(c *gin.Context) {
	if !helpers.CheckAuthorization(c.GetHeader("Authorization")) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var newNetwork models.Network
	if err := c.ShouldBindJSON(&newNetwork); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	AddedNetwork := store.AddNetwork(newNetwork)
	c.JSON(http.StatusOK, AddedNetwork)
}
