package users

import (
	"github.com/atranna/atranna-api/src/internal/auth"
	"github.com/gin-gonic/gin"
)

type selfOrganization struct {
	OrganizationID int                  `json:"organization_id"`
	Role           string               `json:"role"`
	Permissions    []auth.PermissionKey `json:"permissions"`
}

func (h *Handler) GetSelf(c *gin.Context) {
	user, found := h.users.GetByID(c.GetInt("user_id"))
	if !found {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	memberships := h.organizationMember.GetByUserID(user.ID)

	organizations := make([]selfOrganization, 0, len(memberships))
	for _, membership := range memberships {
		organizations = append(organizations, selfOrganization{
			OrganizationID: membership.OrganizationID,
			Role:           membership.Role,
			Permissions:    auth.RolePermissions(membership.Role),
		})
	}

	response := gin.H{
		"user":          user,
		"organizations": organizations,
	}

	c.JSON(200, response)
}
