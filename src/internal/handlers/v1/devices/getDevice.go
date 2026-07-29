package devices

import (
	"atranna-api/src/internal/helpers"
	"atranna-api/src/internal/store"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDevice(c *gin.Context) {
	if !helpers.CheckAuthorization(c.GetHeader("Authorization")) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	device, found := store.GetDeviceByID(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "device not found"})
		return
	}

	c.JSON(http.StatusOK, device)
}
