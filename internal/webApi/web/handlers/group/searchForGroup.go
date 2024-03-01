package group

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
)

func HandleSearchForGroup(c *gin.Context) {
	query, found := c.GetQuery("query")
	if !found {
		c.JSON(400, gin.H{"error": "No query provided"})
		return
	}

	groups, err := database.GroupRepo.SearchForGroup(query)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error finding group"})
		return
	}

	c.JSON(200, gin.H{"groups": groups})
	return
}
