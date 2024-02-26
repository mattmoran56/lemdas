package file

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/database/models"
)

func HandleUpdateFile(c *gin.Context) {
	fileID := c.Param("fileId")

	type FileUpdate struct {
		Name      string `json:"name" binding:"required"`
		IsPublic  bool   `json:"is_public"`
		OwnerID   string `json:"owner_id" binding:"required"`
		Status    string `json:"status" binding:"required"`
		DatasetID string `json:"dataset_id" binding:"required"`
	}

	var fileUpdate FileUpdate
	if err := c.BindJSON(&fileUpdate); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	updatedFile := models.File{
		Base:      models.Base{ID: fileID},
		IsPublic:  fileUpdate.IsPublic,
		OwnerID:   fileUpdate.OwnerID,
		Name:      fileUpdate.Name,
		Status:    fileUpdate.Status,
		DatasetID: fileUpdate.DatasetID,
	}
	file, err := database.FileRepo.UpdateFile(updatedFile)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error updating file"})
		return
	}

	c.JSON(201, file)
	return
}
