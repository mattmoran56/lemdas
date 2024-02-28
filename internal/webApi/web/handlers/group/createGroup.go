package group

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"net/http"
)

func HandleCreateGroup(c *gin.Context) {
	userID := c.GetString("userID")

	type GroupRequest struct {
		Name string `json:"group_name" binding:"required"`
	}

	var groupRequest GroupRequest
	if err := c.ShouldBindJSON(&groupRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	newGroup := models.UserGroup{
		GroupName: groupRequest.Name,
		OwnerID:   userID,
	}
	group, err := database.GroupRepo.Create(newGroup)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error creating group"})
		return
	}

	c.JSON(http.StatusCreated, group)
	return
}
