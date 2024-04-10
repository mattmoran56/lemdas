package file

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"net/http"
)

func HandleUpdateAttribute(c *gin.Context) {
	fileId := c.Param("fileId")
	attributeID := c.Param("fileAttributeId")

	type attributeUpdate struct {
		AttributeName    string `json:"attribute_name" binding:"required"`
		AttributeValue   string `json:"attribute_value" binding:"required"`
		AttributeGroupID string `json:"attribute_group_id" binding:"required"`
	}

	var r attributeUpdate
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	attribute := models.FileAttribute{
		Base: models.Base{
			ID: attributeID,
		},
		FileID:           fileId,
		AttributeName:    r.AttributeName,
		AttributeValue:   r.AttributeValue,
		AttributeGroupID: r.AttributeGroupID,
	}

	attribute, err := database.FileAttributeRepo.UpdateFileAttribute(attribute)

	if err != nil {
		c.JSON(500, gin.H{"error": "Error updating attribute"})
		return
	}

	c.JSON(200, attribute)
	return
}
