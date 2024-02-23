package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"go.uber.org/zap"
	"net/http"
)

func HandleSimpleSearch(c *gin.Context) {
	type SearchRequest struct {
		Query string `json:"query" binding:"required"`
	}
	var r SearchRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No search query found"})
		return
	}

	userID := c.GetString("userID")
	zap.S().Info("userID: ", userID)
	// TODO: Hide support files

	files, err := database.FileRepo.SearchByName(r.Query, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get files"})
		return
	}

	datasets, err := database.DatasetRepo.SearchByName(r.Query, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get datasets"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"files": files, "datasets": datasets})
	return
}
