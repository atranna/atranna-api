package v1

import (
	"atranna-api/src/internal/models"
	"atranna-api/src/internal/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddInterface(c *gin.Context) {
	var newInterface models.Interface
	if err := c.ShouldBindJSON(&newInterface); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newInterface.Device_ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "device_id is required"})
		return
	}

	addedInterface, error := store.AddInterface(newInterface, newInterface.Device_ID)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}
	c.JSON(http.StatusCreated, addedInterface)
}
