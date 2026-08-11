package middlewares

import (
	"net/http"

	"github.com/atranna/atranna-api/src/internal/auth"
	"github.com/gin-gonic/gin"
)

func RequirePermissionMiddleware(key auth.PermissionKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")
		permissions, ok := auth.SystemRoles[role]
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "invalid role"})
			return
		}

		for _, permission := range permissions {
			if permission == key {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "your role does not have the required permission to perform this action: " + string(key)})
	}
}