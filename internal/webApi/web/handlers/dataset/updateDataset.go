package dataset

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/database/models"
)

func HandleUpdateDataset(c *gin.Context) {
	type DatasetUpdate struct {
		DatasetName string `json:"dataset_name" binding:"required"`
		IsPublic    bool   `json:"is_public"`
		OwnerID     string `json:"owner_id" binding:"required"`
	}

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

	if datasetUpdate.IsPublic {
		details := map[string]interface{}{
			"dataset_name": dataset.DatasetName,
		}
		detailsJSON, err := json.Marshal(details)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error marshaling details to JSON"})
			return
		}

		activity := models.Activity{
			UserID:  datasetUpdate.OwnerID,
			Type:    "make_public",
			Details: string(detailsJSON),
		}
		_, err = database.ActivityRepo.CreateActivity(activity)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error creating activity"})
			return
		}
	}

	c.JSON(201, dataset)

}
