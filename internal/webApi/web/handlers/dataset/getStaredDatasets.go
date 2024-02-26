package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
)

func HandleGetStaredDatasets(c *gin.Context) {
	userID := c.MustGet("userID").(string)

	datasets, err := database.DatasetRepo.GetStaredDatasets(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error finding stared datasets"})
		return
	}

	c.JSON(200, gin.H{"datasets": datasets})
	return
}
