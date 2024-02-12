package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
)

func HandleGetFiles(c *gin.Context) {
	datasetID := c.Param("datasetId")
	userId := c.MustGet("userID").(string)

	files, err := database.FileRepo.GetFilesForDataset(datasetID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error finding files"})
		return
	}

	// Check each file has permission to be accessed by user
	for i, el := range files {
		if el.OwnerId != userId {
			files = append(files[:i], files[i+1:]...)
		}
	}

	c.JSON(200, files)
	return
}
