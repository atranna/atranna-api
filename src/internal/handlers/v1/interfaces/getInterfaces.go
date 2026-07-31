package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetInterfaces(c *gin.Context) {
	c.JSON(http.StatusOK, h.interfaces.GetAll())
}
