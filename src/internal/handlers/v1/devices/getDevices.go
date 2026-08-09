package devices

import (
	"net/http"
	"strconv"

	"github.com/atranna/atranna-api/src/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetDevices(c *gin.Context) {
	orgID := c.GetInt("org_id")
	if orgID == 0 {
		orgIDHeader := c.GetHeader("X-Org-ID")
		if orgIDHeader == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "X-Org-ID header is required"})
			return
		}
		parsedOrgID, err := strconv.Atoi(orgIDHeader)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "X-Org-ID must be a valid integer"})
			return
		}
		orgID = parsedOrgID
	}

	devices := h.devices.GetAll(orgID)
	if devices == nil {
		devices = []models.Device{}
	}
	c.JSON(http.StatusOK, devices)
}
