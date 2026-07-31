package networks

import (
	"atranna-api/src/internal/store"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteNetwork(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	network, deleted := store.DeleteNetwork(id)
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": "network not found"})
		return
	}

	c.JSON(http.StatusOK, network)
}
