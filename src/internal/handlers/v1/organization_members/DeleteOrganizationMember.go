package organizationMembers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteOrganizationMember(c *gin.Context) {
	organizationID := c.GetInt("org_id")

	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	_, ok := h.organizationMembers.Delete(organizationID, userID)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "member not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User successfully removed"})
}