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

	user, found := h.users.GetByID(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
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
