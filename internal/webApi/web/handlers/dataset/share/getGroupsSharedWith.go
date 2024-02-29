package share

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"net/http"
)

func HandleGetGroupsSharedWith(c *gin.Context) {
	datasetID := c.Param("datasetId")

	groups, err := database.GroupShareDatasetRepo.GetGroupShareDatasetsForDatasetId(datasetID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get groups shared with"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"groups": groups})
	return
}
