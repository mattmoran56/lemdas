package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"go.uber.org/zap"
)

func HandleGetDatasets(c *gin.Context) {
	orderBy, found := c.GetQuery("orderBy")
	zap.S().Debug(orderBy)
	if !found {
		datasets, err := database.DatasetRepo.GetDatasets()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, datasets)
	} else {
		datasets, err := database.DatasetRepo.GetDatasetsOrderBy(orderBy)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, datasets)
	}

	return
}
