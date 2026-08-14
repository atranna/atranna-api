package devices

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/atranna/atranna-api/src/internal/models"
)

func (h *Handler) Add(c *gin.Context) {
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

	var req struct {
		Hostname string `json:"hostname" binding:"required"`
		IP       string `json:"ip" binding:"required"`
		Vendor   string `json:"vendor" binding:"required"`
		Model    string `json:"model" binding:"required"`
		Type     string `json:"type" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newDevice := models.Device{
		Hostname: req.Hostname,
		IP:       req.IP,
		Vendor:   req.Vendor,
		Model:    req.Model,
		Type:     req.Type,
		OrgID:    orgID,
	}

	addedDevice, err := h.devices.Add(newDevice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, addedDevice)
}
