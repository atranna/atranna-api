package devices

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetDevices(c *gin.Context) {
	c.JSON(http.StatusOK, h.devices.GetAll())
}
