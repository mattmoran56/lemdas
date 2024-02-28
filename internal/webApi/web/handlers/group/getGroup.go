package group

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"net/http"
)

func HandleGetGroup(c *gin.Context) {
	groupId := c.Param("groupId")

	group, err := database.GroupRepo.GetGroupById(groupId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching group"})
		return
	}

	c.JSON(http.StatusOK, group)
}
