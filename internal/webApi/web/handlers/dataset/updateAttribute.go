package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"net/http"
)

func HandleUpdateAttribute(c *gin.Context) {
	datasetID := c.Param("datasetId")
	attributeID := c.Param("attributeId")

	type attributeUpdate struct {
		AttributeName  string `json:"attribute_name" binding:"required"`
		AttributeValue string `json:"attribute_value" binding:"required"`
	}

	var r attributeUpdate
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	attribute := models.DatasetAttribute{
		Base: models.Base{
			ID: attributeID,
		},
		DatasetID:      datasetID,
		AttributeName:  r.AttributeName,
		AttributeValue: r.AttributeValue,
	}

	attribute, err := database.DatasetAttributeRepo.UpdateDatasetAttribute(attribute)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error updating attribute"})
		return
	}

	c.JSON(201, attribute)
	return
}
