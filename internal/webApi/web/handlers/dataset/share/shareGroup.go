package share

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"net/http"
)

func HandleShareWithGroup(c *gin.Context) {
	datasetID := c.Param("datasetId")

	type ShareRequest struct {
		ShareeEID   string `json:"user_id" binding:"required"`
		WriteAccess bool   `json:"write_access" default:"false"`
	}

	var shareRequest ShareRequest
	if err := c.ShouldBindJSON(&shareRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	_, err := database.GroupRepo.GetGroupById(shareRequest.ShareeEID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Group not found"})
		return
	}

	_, err = database.GroupShareDatasetRepo.ShareDatasetWithGroup(datasetID, shareRequest.ShareeEID, shareRequest.WriteAccess)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to share dataset"})
		return
	}

	c.JSON(http.StatusCreated, nil)
	return
}
