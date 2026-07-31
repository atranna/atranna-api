package interfaces

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetInterface(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	interf, found := h.interfaces.GetByID(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "interface not found"})
		return
	}

	c.JSON(http.StatusOK, interf)
}
