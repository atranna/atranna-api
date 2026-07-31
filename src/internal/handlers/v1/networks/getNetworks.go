package networks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetNetworks(c *gin.Context) {
	c.JSON(http.StatusOK, h.networks.GetAll())
}
