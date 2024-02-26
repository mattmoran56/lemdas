package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
)

func HandleGetFiles(c *gin.Context) {
	datasetID := c.Param("datasetId")
	//userID := c.MustGet("userID").(string)

	// check dataset exists
	_, err := database.DatasetRepo.GetDatasetByID(datasetID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error finding dataset"})
		return
	}

	files, err := database.FileRepo.GetFilesForDataset(datasetID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error finding files"})
		return
	}

	c.JSON(200, gin.H{"files": files})
	return
}
