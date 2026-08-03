package organizations

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetOrganizations(c *gin.Context) {
	c.JSON(http.StatusOK, h.organizations.GetAll())
}
