package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
)

func HandleDeleteDataset(c *gin.Context) {
	datasetID := c.Param("datasetId")

	files, err := database.FileRepo.GetFilesForDataset(datasetID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error finding files"})
		return
	}
	if len(files) > 0 {
		c.JSON(400, gin.H{"error": "Dataset has files"})
		return
	}

	err = database.DatasetRepo.DeleteDatasetByID(datasetID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error deleting dataset"})
		return
	}

	c.JSON(204, nil)
}
