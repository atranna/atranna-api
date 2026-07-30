package v1

import (
	"atranna-api/src/internal/helpers"
	"atranna-api/src/internal/models"
	"atranna-api/src/internal/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddInterface(c *gin.Context) {
	if !helpers.CheckAuthorization(c.GetHeader("Authorization")) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var newInterface models.Interface
	if err := c.ShouldBindJSON(&newInterface); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newInterface.Device_ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "device_id is required"})
		return
	}

	addedInterface := store.AddInterface(newInterface, newInterface.Device_ID)
	c.JSON(http.StatusCreated, addedInterface)
}
