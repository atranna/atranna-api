package networks

import (
	"atranna-api/src/internal/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNetworks(c *gin.Context) {
	c.JSON(http.StatusOK, store.Networks)
}
