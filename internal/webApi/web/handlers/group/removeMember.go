package group

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"net/http"
)

func HandleDeleteMember(c *gin.Context) {
	groupID := c.Param("groupId")
	userID := c.Param("userId")

	isMember, err := database.GroupMemberRepo.IsUserInGroup(userID, groupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking membership"})
		return
	}
	if !isMember {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	err = database.GroupMemberRepo.RemoveUserFromGroup(userID, groupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error removing member"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
