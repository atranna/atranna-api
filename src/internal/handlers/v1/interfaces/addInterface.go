package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/atranna/atranna-api/src/internal/models"
)

func (h *Handler) Add(c *gin.Context) {
	var newInterface models.Interface
	if err := c.ShouldBindJSON(&newInterface); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newInterface.DeviceID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "device_id is required"})
		return
	}

	addedInterface, err := h.interfaces.Add(newInterface)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, addedInterface)
}
