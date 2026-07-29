package v1

import (
	"atranna-api/src/internal/helpers"
	"atranna-api/src/internal/store"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetInterfacesByDeviceID(c *gin.Context) {
	if !helpers.CheckAuthorization(c.GetHeader("Authorization")) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
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

	c.JSON(http.StatusOK, store.GetInterfacesByDeviceID(device_id))
}
