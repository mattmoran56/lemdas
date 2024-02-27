package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
)

func HandleSearchForUser(c *gin.Context) {
	query, found := c.GetQuery("query")
	if !found {
		c.JSON(400, gin.H{"error": "No query provided"})
		return
	}

	users, err := database.UserRepo.SearchForUser(query)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error finding user"})
		return
	}

	c.JSON(200, gin.H{"users": users})
	return
}
