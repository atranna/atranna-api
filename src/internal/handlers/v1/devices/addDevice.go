package v1

import (
	"atranna-api/src/internal/models"
	"atranna-api/src/internal/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddDevice(c *gin.Context) {
	var newDevice models.Device
	if err := c.ShouldBindJSON(&newDevice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	AddedDevice := store.AddDevice(newDevice)
	c.JSON(http.StatusOK, AddedDevice)
}