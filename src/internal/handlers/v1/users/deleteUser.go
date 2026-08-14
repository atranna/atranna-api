package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	requestingUserID := c.GetInt("user_id")
	if requestingUserID != id {
		c.JSON(http.StatusForbidden, gin.H{"error": "you can only delete your own account"})
		return
	}

	_, deleted := h.users.Delete(id)
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
