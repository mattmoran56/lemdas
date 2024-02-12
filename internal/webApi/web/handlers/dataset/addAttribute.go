package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"net/http"
)

func HandleCreateAttribute(c *gin.Context) {
	datasetId := c.Param("datasetId")
	type NewAttribute struct {
		AttributeName  string `json:"attribute_name"`
		AttributeValue string `json:"attribute_value"`
	}

	var r NewAttribute
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	datasetAttribute := models.DatasetAttribute{
		DatasetID:      datasetId,
		AttributeName:  r.AttributeName,
		AttributeValue: r.AttributeValue,
	}

	err := database.DatasetAttributeRepo.CreateDatasetAttribute(datasetAttribute)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error creating attribute"})
		return
	}

	c.JSON(200, gin.H{"success": "Attribute created"})
	return
}
