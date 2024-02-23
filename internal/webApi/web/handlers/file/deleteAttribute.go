package file

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
)

func HandleDeleteAttribute(c *gin.Context) {
	attributeId := c.Param("fileAttributeId")
	fileId := c.Param("fileId")

	// Check if attribute exists
	attribute, err := database.FileAttributeRepo.GetFileAttributeByID(attributeId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error finding attribute"})
		return
	}

	if attribute.FileID != fileId {
		c.JSON(404, gin.H{"error": "Attribute not found on file"})
		return
	}

	err = database.FileAttributeRepo.DeleteFileAttribute(attributeId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error deleting attribute"})
		return
	}

	c.JSON(204, nil)
	return
}
