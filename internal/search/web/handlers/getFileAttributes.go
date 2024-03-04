package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
)

func HandleGetFileAttributes(c *gin.Context) {
	attributes, err := database.FileAttributeRepo.GetAllAttributeNames()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error getting attribute names"})
		return
	}

	c.JSON(200, gin.H{"attributes": attributes})
	return
}
