package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"net/http"
)

func HandleSimpleSearch(c *gin.Context) {
	type SearchRequest struct {
		Query string `json:"query"`
	}
	var r SearchRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Check user can see files and datasets
	// TODO: Hide support files

	files, err := database.FileRepo.SearchByName(r.Query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get files"})
		return
	}

	datasets, err := database.DatasetRepo.SearchByName(r.Query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get datasets"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"files": files, "datasets": datasets})
	return
}
