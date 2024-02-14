package file

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"net/http"
)

func HandleUpdateAttribute(c *gin.Context) {
	fileId := c.Param("fileId")

	type attributeUpdate struct {
		AttributeID    string `json:"attribute_id"`
		AttributeName  string `json:"attribute_name"`
		AttributeValue string `json:"attribute_value"`
	}

	var r attributeUpdate
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attribute := models.FileAttribute{
		Base: models.Base{
			ID: r.AttributeID,
		},
		FileID:         fileId,
		AttributeName:  r.AttributeName,
		AttributeValue: r.AttributeValue,
	}

	err := database.FileAttributeRepo.UpdateFileAttribute(attribute)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error updating attribute"})
		return
	}

	c.JSON(200, gin.H{})
	return
}
