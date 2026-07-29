package v1

import (
	"atranna-api/src/internal/store"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteDevice(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	for i, device := range store.Devices {
		if device.ID == id {
			store.Devices = append(store.Devices[:i], store.Devices[i+1:]...)
			c.JSON(http.StatusOK, device)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "device not found"})
}	