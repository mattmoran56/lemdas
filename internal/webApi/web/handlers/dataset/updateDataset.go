package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"go.uber.org/zap"
)

func HandleUpdateDataset(c *gin.Context) {
	type DatasetUpdate struct {
		DatasetName string `json:"dataset_name" binding:"required"`
		IsPublic    bool   `json:"is_public"`
		OwnerID     string `json:"owner_id" binding:"required"`
	}

	zap.S().Info(c.Request.Body)

	var datasetUpdate DatasetUpdate
	if err := c.BindJSON(&datasetUpdate); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	datasetID := c.Param("datasetId")

	updatedDataset := models.Dataset{
		Base:        models.Base{ID: datasetID},
		DatasetName: datasetUpdate.DatasetName,
		IsPublic:    datasetUpdate.IsPublic,
		OwnerID:     datasetUpdate.OwnerID,
	}

	dataset, err := database.DatasetRepo.UpdateDataset(updatedDataset)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error updating dataset"})
		return
	}

	c.JSON(201, dataset)

}
