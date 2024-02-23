package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/database/models"
)

func HandleGetDataset(c *gin.Context) {
	type DatasetResponse struct {
		models.Dataset
		OwnerName string `json:"owner_name"`
	}

	datasetID := c.Param("datasetId")

	dataset, err := database.DatasetRepo.GetDatasetByID(datasetID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error finding dataset"})
		return
	}

	user, err := database.UserRepo.GetUserByID(dataset.OwnerID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error finding dataset owner"})
		return
	}

	datasetResponse := DatasetResponse{
		Dataset:   dataset,
		OwnerName: user.FirstName + " " + user.LastName,
	}

	c.JSON(200, datasetResponse)
	return
}
