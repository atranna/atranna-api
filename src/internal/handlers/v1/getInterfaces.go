package v1

import (
	"atranna-api/src/internal/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetInterfaces(c *gin.Context) {

	c.JSON(http.StatusOK, store.Interfaces)
}
