package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
)

func HandleGetAttributes(c *gin.Context) {
	datasetId := c.Param("datasetId")
	userID := c.MustGet("userID").(string)
	orderBy, found := c.GetQuery("orderBy")

	if found {
		attributes, err := database.DatasetAttributeRepo.GetDatasetAttributeByDatasetIDOrderBy(datasetId, userID, orderBy)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error fetching attributes"})
			return
		}
		c.JSON(200, gin.H{"attributes": attributes})
		return
	} else {
		attributes, err := database.DatasetAttributeRepo.GetDatasetAttributeByDatasetID(datasetId, userID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error fetching attributes"})
			return
		}
		c.JSON(200, gin.H{"attributes": attributes})
		return
	}

}
