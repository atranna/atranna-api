package organizations

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetOrganizations(c *gin.Context) {
	userID := c.GetInt("user_id")
	membershipRecords := h.organizationMembers.GetByUserID(userID)
	
	organizations := make([]map[string]interface{}, 0, len(membershipRecords))
	for _, record := range membershipRecords {
		org, ok := h.organizations.GetByID(record.OrganizationID)
		if ok != true {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve organization details"})
			return
		}
		organizations = append(organizations, map[string]interface{}{
			"id":   org.ID,
			"name": org.Name,
			"slug": org.Slug,
		})
	}

	c.JSON(http.StatusOK, organizations)
}
