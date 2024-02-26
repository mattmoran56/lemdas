package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
)

func HandleDeleteCollaborator(c *gin.Context) {
	datasetID := c.Param("datasetId")
	collaboratorID := c.Param("collaboratorId")

	_, err := database.DatasetCollaboratorRepo.GetCollaborator(datasetID, collaboratorID)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(404, gin.H{"error": "Collaborator not found"})
			return
		}
		c.JSON(500, gin.H{"error": "Error finding collaborator"})
		return
	}

	err = database.DatasetCollaboratorRepo.RemoveCollaborator(datasetID, collaboratorID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error deleting collaborator"})
		return
	}

	c.JSON(204, nil)
	return
}
