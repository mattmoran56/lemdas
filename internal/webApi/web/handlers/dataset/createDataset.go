package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/database/models"
)

func HandleCreateDataset(c *gin.Context) {
	type DatasetRequest struct {
		DatasetName string `json:"dataset_name"`
	}
	var dataset DatasetRequest
	err := c.BindJSON(&dataset)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	userId := c.MustGet("userID").(string)

	newDataset := models.Dataset{
		DatasetName: dataset.DatasetName,
		OwnerID:     userId,
	}
	createdDataset, err := database.DatasetRepo.CreateDataset(newDataset)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error adding new dataset to database"})
		return
	}

	c.JSON(200, createdDataset)
	return
}
