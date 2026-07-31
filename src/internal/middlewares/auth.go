package middlewares

import (
	"atranna-api/src/internal/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        if !helpers.CheckAuthorization(c.GetHeader("Authorization")) {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }
        c.Next()
    }
}