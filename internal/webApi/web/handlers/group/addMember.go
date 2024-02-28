package group

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"net/http"
)

func HandleAddMember(c *gin.Context) {
	groupID := c.Param("groupId")

	type AddMemberRequest struct {
		UserID string `json:"user_id" binding:"required"`
	}

	var r AddMemberRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := database.GroupMemberRepo.AddUserToGroup(r.UserID, groupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding member to group"})
		return
	}

	c.JSON(http.StatusCreated, nil)
}
