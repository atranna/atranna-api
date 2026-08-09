package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        org := c.GetHeader("X-Org-ID")
        if org == "" {
            c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "X-Org-ID header is required"})
            return
        }

        orgID, err := strconv.Atoi(org)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "X-Org-ID must be a valid integer"})
            return
        }
        c.Set("org_id", orgID)

        c.Next()
    }
}