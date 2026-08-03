package organizations

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/atranna/atranna-api/src/internal/models"
)

func (h *Handler) Add(c *gin.Context) {
	var newOrganization models.Organization
	if err := c.ShouldBindJSON(&newOrganization); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	addedOrganization, err := h.organizations.Add(newOrganization)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, addedOrganization)
}
