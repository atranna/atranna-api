package organizationMembers

import (
	"net/http"
	"strconv"

	"github.com/atranna/atranna-api/src/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddOrganizationMember(c *gin.Context) {
	organizationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid organization ID"})
		return
	}

	var request struct {
		UserID int `json:"user_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newMember := models.OrganizationMember{OrganizationID: organizationID, UserID: request.UserID}
	addedMember, err := h.organizationMembers.Add(newMember)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, addedMember)
}