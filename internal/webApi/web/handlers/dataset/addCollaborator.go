package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
)

func HandleAddCollaborator(c *gin.Context) {
	datasetID := c.Param("datasetId")

	type AddCollaboratorRequest struct {
		UserID string `json:"user_id" binding:"required"`
	}

	var req AddCollaboratorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	dataset, err := database.DatasetRepo.GetDatasetByID(datasetID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error finding dataset"})
		return
	}

	_, err = database.UserRepo.GetUserByID(req.UserID)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		c.JSON(500, gin.H{"error": "Error finding user"})
		return
	}

	if dataset.OwnerID == req.UserID {
		c.JSON(400, gin.H{"error": "User is owner of dataset"})
		return
	}

	_, err = database.DatasetCollaboratorRepo.GetCollaborator(datasetID, req.UserID)
	if err == nil {
		c.JSON(400, gin.H{"error": "User is already a collaborator"})
		return
	}

	collaborator, err := database.DatasetCollaboratorRepo.AddCollaborator(datasetID, req.UserID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error adding collaborator"})
		return
	}

	c.JSON(200, collaborator)
	return
}
