package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
)

func HandleGetAttributes(c *gin.Context) {
	datasetId := c.Param("datasetId")
	orderBy, found := c.GetQuery("orderBy")

	if found {
		attributes, err := database.DatasetAttributeRepo.GetDatasetAttributeByDatasetIDOrderBy(datasetId, orderBy)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error fetching attributes"})
			return
		}
		c.JSON(200, attributes)
		return
	} else {
		attributes, err := database.DatasetAttributeRepo.GetDatasetAttributeByDatasetID(datasetId)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error fetching attributes"})
			return
		}
		c.JSON(200, attributes)
		return
	}

}
