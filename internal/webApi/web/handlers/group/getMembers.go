package group

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"net/http"
)

func HandleGetMembers(c *gin.Context) {
	groupId := c.Param("groupId")

	members, err := database.GroupMemberRepo.GetGroupMembers(groupId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching members"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"members": members})
}
