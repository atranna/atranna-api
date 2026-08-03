package middlewares

import (
	"atranna-api/src/internal/config"
	"atranna-api/src/internal/helpers"
	"strings"

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
            c.Next()
            return
        }

        token := strings.TrimSpace(authorization)
        if strings.HasPrefix(strings.ToLower(token), "bearer ") {
            token = strings.TrimSpace(token[len("bearer "):])
        }

        c.Next()
    }
}