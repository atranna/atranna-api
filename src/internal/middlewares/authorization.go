package middlewares

import (
	"net/http"
	"strconv"

	"github.com/atranna/atranna-api/src/internal/repository"
	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware(orgMembersRepo repository.OrganizationMemberRepository) gin.HandlerFunc {
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

        userID := c.GetInt("user_id")
        if _, isMember := orgMembersRepo.GetRole(orgID, userID); !isMember {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You are not a member of this organization"})
            return
        }

        c.Set("org_id", orgID)

        c.Next()
    }
}