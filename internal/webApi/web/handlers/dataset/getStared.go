package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
)

func HandleGetStared(c *gin.Context) {
	userID := c.GetString("userID")
	datasetID := c.Param("datasetId")

	stared, err := database.StaredDatasetRepo.GetStaredDataset(userID, datasetID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error finding stared dataset"})
		return
	}

	c.JSON(200, gin.H{"stared": stared})
	return
}
