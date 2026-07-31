package devices

import (
	"atranna-api/src/internal/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDevices(c *gin.Context) {
	c.JSON(http.StatusOK, store.Devices)
}
