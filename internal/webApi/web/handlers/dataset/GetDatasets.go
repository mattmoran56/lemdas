package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
)

func HandleGetDatasets(c *gin.Context) {
	datasets, err := database.DatasetRepo.GetDatasets()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, datasets)
	return
}
