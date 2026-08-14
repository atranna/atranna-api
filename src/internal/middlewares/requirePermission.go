package middlewares

import (
	"net/http"
	"strconv"

	"github.com/atranna/atranna-api/src/internal/auth"
	"github.com/atranna/atranna-api/src/internal/repository"
	"github.com/gin-gonic/gin"
)

func RequirePermissionMiddleware(key auth.PermissionKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")
		if _, ok := auth.SystemRoles[role]; !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "invalid role"})
			return
		}

		if auth.RoleHasPermission(role, key) {
			c.Next()
			return
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "your role does not have the required permission to perform this action: " + string(key)})
	}
}

func RequireOrganizationPermissionMiddleware(orgMembersRepo repository.OrganizationMemberRepository, key auth.PermissionKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetInt("user_id") == -1 {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "master token cannot be used for this action"})
			return
		}

		orgID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid organization id"})
			return
		}

		role, isMember := orgMembersRepo.GetRole(orgID, c.GetInt("user_id"))
		if !isMember {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "you are not a member of this organization"})
			return
		}

		if auth.RoleHasPermission(role, key) {
			c.Next()
			return
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "your role does not have the required permission to perform this action: " + string(key)})
	}
}
