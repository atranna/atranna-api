package networks

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	_, deleted := h.networks.Delete(id)
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": "network not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
