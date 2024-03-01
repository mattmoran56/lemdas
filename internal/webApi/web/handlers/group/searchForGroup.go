package group

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"go.uber.org/zap"
)

func HandleSearchForGroup(c *gin.Context) {
	userID := c.GetString("userID")
	zap.S().Debug(userID)

	query, found := c.GetQuery("query")
	if !found {
		c.JSON(400, gin.H{"error": "No query provided"})
		return
	}

	groups, err := database.GroupRepo.SearchForGroupUserIsIn(query, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error finding group"})
		return
	}

	c.JSON(200, gin.H{"groups": groups})
	return
}
