package middlewares

import (
	"net/http"
	"strings"

	"github.com/atranna/atranna-api/src/internal/config"
	"github.com/atranna/atranna-api/src/internal/helpers"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        if !config.Current.Auth.Enable {
            c.Next()
            return
        }

        authorization := c.GetHeader("Authorization")

        if helpers.CheckAuthorization(authorization, config.Current) {
            c.Set("user_id", -1)
            c.Next()
            return
        }

        token := strings.TrimSpace(authorization)
        if strings.HasPrefix(strings.ToLower(token), "bearer ") {
            token = strings.TrimSpace(token[len("bearer "):])
        }
        if token == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }

        userID, err := helpers.ValidateJWT(token)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }

        c.Set("user_id", userID)
        c.Next()
    }
}