package organizationMembers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetOrganizationMembers(c *gin.Context) {
	organizationID := c.GetInt("org_id")
	members, err := h.organizationMembers.GetByOrganizationID(organizationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, members)
}