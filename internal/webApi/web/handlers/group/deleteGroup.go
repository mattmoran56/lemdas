package group

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"net/http"
)

func HandleDeleteGroup(c *gin.Context) {
	groupId := c.Param("groupId")
	err := database.GroupRepo.DeleteGroup(groupId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting group"})
		return
	}

	err = database.GroupMemberRepo.DeleteGroupMembers(groupId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting group members"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
