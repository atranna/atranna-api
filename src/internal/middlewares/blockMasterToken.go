package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BlockMasterTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetInt("user_id") == -1 {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "master token cannot be used for this action"})
			return
		}
		c.Next()
	}
}