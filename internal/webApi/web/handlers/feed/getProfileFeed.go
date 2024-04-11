package feed

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
)

func HandleGetProfileFeed(c *gin.Context) {
	userId := c.Param("userId")

	activities, err := database.ActivityRepo.GetActivitiesByUserIDAndType(userId, "make_public")
	if err != nil {
		c.JSON(500, gin.H{"error": "Error fetching activities"})
		return
	}

	c.JSON(200, gin.H{"activities": activities})
}
