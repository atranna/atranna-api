package v1

import (
	"atranna-api/src/internal/models"
	"atranna-api/src/internal/store"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddInterface(c *gin.Context) {
	var newInterface models.Interface
	if err := c.ShouldBindJSON(&newInterface); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	device_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if _, found := store.GetDeviceByID(device_id); !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "device not found"})
		return
	}
	addedInterface := store.AddInterface(newInterface, device_id)
	c.JSON(http.StatusOK, addedInterface)
}
