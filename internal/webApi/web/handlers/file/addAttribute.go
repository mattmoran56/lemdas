package file

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"net/http"
)

func HandleCreateAttribute(c *gin.Context) {
	fileId := c.Param("fileId")
	type NewAttribute struct {
		AttributeName  string `json:"attribute_name"`
		AttributeValue string `json:"attribute_value"`
	}

	var r NewAttribute
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileAttribute := models.FileAttribute{
		FileID:         fileId,
		AttributeName:  r.AttributeName,
		AttributeValue: r.AttributeValue,
	}

	id, err := database.FileAttributeRepo.CreateFileAttribute(fileAttribute)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error creating attribute"})
		return
	}

	c.JSON(200, gin.H{"attribute_id": id})
	return
}
