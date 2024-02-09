package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/database/models"
)

func HandleCreateDataset(c *gin.Context) {
	var dataset models.Dataset
	err := c.BindJSON(&dataset)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	err = database.DatasetRepo.CreateDataset(dataset)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error adding new dataset to database"})
		return
	}

	dataset, err = database.DatasetRepo.GetDatasetByName(dataset.DatasetName)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error retrieving dataset from database"})
		return
	}

	c.JSON(200, dataset)
	return
}
