package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/database/repositories"
	"go.uber.org/zap"
	"net/http"
)

func HandleSearch(c *gin.Context) {
	userID := c.GetString("userID")
	var r []repositories.SearchQuery
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	zap.S().Debug(r)

	datasets, files, err := database.SearchRepo.Search(r, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"datasets": datasets, "files": files})
	return
}
