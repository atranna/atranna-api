package networks

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetNetwork(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	network, found := h.networks.GetByID(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "network not found"})
		return
	}

	c.JSON(http.StatusOK, network)
}
