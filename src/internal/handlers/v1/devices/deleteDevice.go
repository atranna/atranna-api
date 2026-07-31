package devices

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	_, deleted := h.devices.Delete(id)
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": "device not found"})
		return
	}

	h.interfaces.DeleteByDeviceID(id)

	c.Status(http.StatusNoContent)
}
