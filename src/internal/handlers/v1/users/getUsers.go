package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) GetUsers(c *gin.Context) {
	requestingUserID := c.GetInt("user_id")

	requesterMemberships := h.organizationMember.GetByUserID(requestingUserID)
	requesterOrgs := make(map[int]struct{}, len(requesterMemberships))
	for _, membership := range requesterMemberships {
		requesterOrgs[membership.OrganizationID] = struct{}{}
	}

	users := []user{}
	for _, u := range h.users.GetAll() {
		if u.ID == requestingUserID {
			users = append(users, user{ID: u.ID, Username: u.Username})
			continue
		}

		for _, membership := range h.organizationMember.GetByUserID(u.ID) {
			if _, ok := requesterOrgs[membership.OrganizationID]; ok {
				users = append(users, user{ID: u.ID, Username: u.Username})
				break
			}
		}
	}

	c.JSON(http.StatusOK, users)
}
