package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
)

func HandleDeleteAttribute(c *gin.Context) {
	attributeId := c.Param("datasetAttributeId")
	datasetId := c.Param("datasetId")

	// Check if attribute exists
	attribute, err := database.DatasetAttributeRepo.GetDatasetAttributeByID(attributeId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error finding attribute"})
		return
	}

	if attribute.DatasetID != datasetId {
		c.JSON(404, gin.H{"error": "Attribute not found on dataset"})
		return
	}

	err = database.DatasetAttributeRepo.DeleteDatasetAttribute(attributeId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error deleting attribute"})
		return
	}

	c.JSON(200, gin.H{})
	return
}
