package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
)

func HandleStarDataset(c *gin.Context) {
	userID := c.GetString("userID")
	datasetID := c.Param("datasetId")

	stared, err := database.StaredDatasetRepo.GetStaredDataset(userID, datasetID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error finding stared dataset"})
		return
	}

	if stared {
		err = database.StaredDatasetRepo.UnstarDataset(userID, datasetID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error unstaring dataset"})
			return
		}
		c.JSON(200, gin.H{"stared": false})
		return
	} else {
		_, err = database.StaredDatasetRepo.StarDataset(userID, datasetID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error staring dataset"})
			return
		}
		c.JSON(200, gin.H{"stared": true})
		return
	}
}
