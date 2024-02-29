package share

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"net/http"
)

func HandleDeleteUserShare(c *gin.Context) {
	datasetID := c.Param("datasetId")
	shareeID := c.Param("userId")

	err := database.UserShareDatasetRepo.DeleteUserShareDataset(datasetID, shareeID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete share"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
	return
}
