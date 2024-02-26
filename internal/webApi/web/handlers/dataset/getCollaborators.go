package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
)

func HandleGetCollaborators(c *gin.Context) {
	datasetID := c.Param("datasetId")

	collaborators, err := database.DatasetCollaboratorRepo.GetCollaborators(datasetID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error finding collaborators"})
		return
	}

	c.JSON(200, gin.H{"collaborators": collaborators})
	return
}
