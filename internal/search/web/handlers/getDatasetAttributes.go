package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
)

func HandleGetDatasetAttributes(c *gin.Context) {
	attributes, err := database.DatasetAttributeRepo.GetAllAttributeNames()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error getting dataset attributes"})
		return
	}

	c.JSON(200, gin.H{"attributes": attributes})
	return
}
