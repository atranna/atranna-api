package v1

import (
	"atranna-api/src/internal/store"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDevice(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	for _, device := range store.Devices {
		if device.ID == id {
			c.JSON(http.StatusOK, device)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "device not found"})
}