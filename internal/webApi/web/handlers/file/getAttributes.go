package file

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
)

func HandleGetAttributes(c *gin.Context) {
	fileId := c.Param("fileId")
	orderBy, found := c.GetQuery("orderBy")

	if found {
		attributes, err := database.FileAttributeGroupRepo.GetFileAttributeGroupByFileIDOrderBy(fileId, orderBy)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error fetching attributes"})
			return
		}
		c.JSON(200, gin.H{"attribute_groups": attributes})
		return
	} else {
		attributes, err := database.FileAttributeGroupRepo.GetFileAttributeGroupByFileID(fileId)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error fetching attributes"})
			return
		}
		c.JSON(200, gin.H{"attribute_groups": attributes})
		return
	}

}
