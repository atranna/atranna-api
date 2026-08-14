package middlewares

import (
	"net/http"
	"strconv"

	"github.com/atranna/atranna-api/src/internal/repository"
	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware(orgMembersRepo repository.OrganizationMemberRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt("user_id")
		usersMemberships := orgMembersRepo.GetByUserID(userID)

		if orgHeader := c.GetHeader("X-Org-ID"); orgHeader != "" {
			orgID, err := strconv.Atoi(orgHeader)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "X-Org-ID must be a valid integer"})
				return
			}

			role, isMember := orgMembersRepo.GetRole(orgID, userID)
			if !isMember {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You are not a member of this organization"})
				return
			}

			c.Set("org_id", orgID)
			c.Set("role", role)
			c.Next()
			return
		}

		if len(usersMemberships) == 1 {
			c.Set("org_id", usersMemberships[0].OrganizationID)
			c.Set("role", usersMemberships[0].Role)
			c.Next()
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "X-Org-ID header is required"})
	}
}