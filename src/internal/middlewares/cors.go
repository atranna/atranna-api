package middlewares

import (
	"net/http"
	"strings"

	"github.com/atranna/atranna-api/src/internal/config"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if origin != "" && isAllowedOrigin(origin) {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Vary", "Origin")
		}

		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Max-Age", "86400")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func isAllowedOrigin(origin string) bool {
	allowedOrigins := strings.Split(config.Current.CORS.AllowedOrigins, ",")
	for _, allowedOrigin := range allowedOrigins {
		trimmed := strings.TrimSpace(allowedOrigin)
		if trimmed == "*" || trimmed == origin {
			return true
		}
	}

	return false
}
