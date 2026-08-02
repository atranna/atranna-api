package devices

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"atranna-api/src/internal/models"
)

func (h *Handler) Add(c *gin.Context) {
	var newDevice models.Device
	if err := c.ShouldBindJSON(&newDevice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	addedDevice, err := h.devices.Add(newDevice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, addedDevice)
}
