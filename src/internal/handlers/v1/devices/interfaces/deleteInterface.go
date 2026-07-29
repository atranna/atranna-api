package v1

import (
	"atranna-api/src/internal/store"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteInterface(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("interface_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	interf, deleted := store.DeleteInterface(id)
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": "interface not found"})
		return
	}

	c.JSON(http.StatusOK, interf)
}
