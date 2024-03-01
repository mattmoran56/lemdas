package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
)

func HandleGetSharedDatasets(c *gin.Context) {
	userID := c.GetString("userID")

	datasets, err := database.DatasetRepo.GetUsersSharedDatasets(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "internal error"})
		return
	}

	c.JSON(200, gin.H{"datasets": datasets})
	return
}
