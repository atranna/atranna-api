package organizationMembers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetOrganizationMember(c *gin.Context) {
	organizationID := c.GetInt("org_id")
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	members, err := h.organizationMembers.GetByOrganizationID(organizationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for _, member := range members {
		if member.UserID == userID {
			c.JSON(http.StatusOK, member)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "member not found"})
}
