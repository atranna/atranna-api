package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userWithOrgs struct {
	ID 		int      `json:"id"`
	Email 	string   `json:"email"`
	Username string   `json:"username"`
	DisplayName string `json:"display_name"`
	Orgs     []int `json:"orgs"`
}

func (h *Handler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	requestingUserID := c.GetInt("user_id")
	if requestingUserID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, found := h.users.GetByID(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	requestingUserMemberships := h.organizationMember.GetByUserID(requestingUserID)
	if len(requestingUserMemberships) == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are not a member of any organizations and therefore cannot view other users"})
		return
	}

	requestedUserMemberships := h.organizationMember.GetByUserID(id)
	if len(requestedUserMemberships) == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "requested user is not a member of any organization that you are a member of"})
		return
	}

	commonOrgID := -1
	for _, requesterMembership := range requestingUserMemberships {
		for _, requestedMembership := range requestedUserMemberships {
			if requesterMembership.OrganizationID == requestedMembership.OrganizationID {
				commonOrgID = requesterMembership.OrganizationID
				break
			}
		}
		if commonOrgID != -1 {
			break
		}
	}

	if commonOrgID == -1 {
		c.JSON(http.StatusForbidden, gin.H{"error": "user is not in the same organization"})
		return
	}

	orgMembers := h.organizationMember.GetByUserID(id)
	orgs := make([]int, len(orgMembers))
	for i, om := range orgMembers {
		orgs[i] = om.OrganizationID
	}

	userWithOrgs := userWithOrgs{
		ID:          user.ID,
		Email:       user.Email,
		Username:    user.Username,
		DisplayName: user.DisplayName,
		Orgs:        orgs,
	}

	c.JSON(http.StatusOK, userWithOrgs)
}
