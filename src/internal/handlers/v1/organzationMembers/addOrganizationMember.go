package organizationMembers

import (
	"atranna-api/src/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddOrganizationMember(c *gin.Context) {
	var newMember models.OrganizationMember
	if err := c.ShouldBindJSON(&newMember); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	addedMember, err := h.organizationMembers.Add(newMember)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, addedMember)
}