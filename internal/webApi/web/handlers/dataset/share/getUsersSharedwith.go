package share

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"go.uber.org/zap"
	"net/http"
)

func HandleGetUsersSharedWith(c *gin.Context) {
	datasetID := c.Param("datasetId")
	zap.S().Debug(datasetID)

	users, err := database.UserShareDatasetRepo.GetUserShareDatasetsForDatasetId(datasetID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
	return
}
