package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
)

func HandleDeleteAttribute(c *gin.Context) {
	attributeId := c.Param("datasetAttributeId")

	err := database.DatasetAttributeRepo.DeleteDatasetAttribute(attributeId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error deleting attribute"})
		return
	}

	c.JSON(204, nil)
	return
}
