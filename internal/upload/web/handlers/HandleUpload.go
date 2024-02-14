package handlers

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"github.com/mattmoran/fyp/api/pkg/utils"
	"go.uber.org/zap"
	"net/http"
	"os"
	"path/filepath"
)

func getFileExtension(filename string) string {
	ext := filepath.Ext(filename)
	if ext != "" {
		// Remove the dot (.) from the extension
		return ext[1:]
	}
	return ""
}

func HandleUpload(c *gin.Context) {
	type UploadFileRequest struct {
		DatasetID string `form:"dataset_id"`
		IsPublic  bool   `form:"is_public"`
	}

	var r UploadFileRequest
	if err := c.ShouldBind(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error binding to request: " + err.Error()})
		return
	}

	zap.S().Debug(r)

	form, _ := c.MultipartForm()
	files := form.File["file"]
	zap.S().Debug(form)
	zap.S().Debug(files)

	for _, file := range files {
		zap.S().Debug("Uploading file: ", file.Filename)
		randomFileId := uuid.New().String() + "." + getFileExtension(file.Filename)

		if err := c.SaveUploadedFile(file, ".temp/"+randomFileId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error saving file: " + err.Error()})
			return
		}

		credential, err := azidentity.NewDefaultAzureCredential(nil)
		utils.HandleHandlerError(c, err)

		url := "https://synopticprojectstorage.blob.core.windows.net/"
		client, err := azblob.NewClient(url, credential, nil)
		utils.HandleHandlerError(c, err)

		containerName := "fyp-uploads"

		// Open the file to upload
		fileHandler, err := os.Open(".temp/" + randomFileId)
		utils.HandleHandlerError(c, err)

		// close the file after it is no longer required.
		defer func(file *os.File) {
			err = file.Close()
			utils.HandleHandlerError(c, err)
		}(fileHandler)

		// delete the local file if required.
		defer func(name string) {
			err = os.Remove(name)
			utils.HandleHandlerError(c, err)
		}(".temp/" + randomFileId)

		_, err = client.UploadFile(context.TODO(), containerName, randomFileId, fileHandler, &azblob.UploadBufferOptions{})
		utils.HandleHandlerError(c, err)

		// Add the file to the database
		userId := c.MustGet("userID").(string)
		fileObject := models.File{
			Base:      models.Base{ID: randomFileId},
			Name:      file.Filename,
			OwnerId:   userId,
			DatasetID: r.DatasetID,
			Status:    "uploaded",
			IsPublic:  r.IsPublic,
		}

		err = database.FileRepo.CreateFile(fileObject)
		utils.HandleHandlerError(c, err)

		// TODO: Trigger upload processing
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
	return
}
