package file

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"github.com/mattmoran/fyp/api/pkg/utils"
)

func HandleGetFile(c *gin.Context) {
	type FileResponse struct {
		models.File
		DatasetName string `json:"dataset_name"`
		OwnerName   string `json:"owner_name"`
	}

	fileId := c.Param("fileId")

	file, err := database.FileRepo.GetFileByID(fileId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error finding file"})
		return
	}

	if utils.CheckUserHasPermission(fileId, c.MustGet("userId").(string)) == false {
		c.JSON(401, gin.H{"error": "User does not have permission to view this file"})
		return
	}

	dataset, err := database.DatasetRepo.GetDatasetByID(file.DatasetID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error finding dataset"})
		return
	}

	user, err := database.UserRepo.GetUserByID(file.OwnerId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error finding file owner"})
		return
	}

	var fileResponse = FileResponse{
		File:        file,
		DatasetName: dataset.DatasetName,
		OwnerName:   user.FirstName + " " + user.LastName,
	}

	c.JSON(200, fileResponse)
	return
}
