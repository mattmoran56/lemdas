package file

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
)

func HandleGetAttributes(c *gin.Context) {
	fileId := c.Param("fileId")
	orderBy, found := c.GetQuery("orderBy")

	if found {
		attributes, err := database.FileAttributeRepo.GetFileAttributeByFileIDOrderBy(fileId, orderBy)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error fetching attributes"})
			return
		}
		c.JSON(200, attributes)
		return
	} else {
		attributes, err := database.FileAttributeRepo.GetFileAttributeByFileID(fileId)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error fetching attributes"})
			return
		}
		c.JSON(200, attributes)
		return
	}

}
