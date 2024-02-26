package file

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
)

func HandleDeleteFile(c *gin.Context) {
	fileID := c.Param("fileId")

	// check file exists
	_, err := database.FileRepo.GetFileByID(fileID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error finding file"})
		return
	}

	err = database.FileRepo.DeleteFile(fileID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error deleting file"})
		return
	}

	c.JSON(204, nil)
	return
}
