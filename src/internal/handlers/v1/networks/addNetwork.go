package networks

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"atranna-api/src/internal/models"
)

func (h *Handler) Add(c *gin.Context) {
	var newNetwork models.Network
	if err := c.ShouldBindJSON(&newNetwork); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	addedNetwork, err := h.networks.Add(newNetwork)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, addedNetwork)
}
