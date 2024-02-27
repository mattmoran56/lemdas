package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"github.com/mattmoran/fyp/api/pkg/utils"
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
		DatasetID string `form:"dataset_id" binding:"required"`
		IsPublic  bool   `form:"is_public"`
	}

	var r UploadFileRequest
	if err := c.ShouldBind(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dataset or is_public is missing"})
		return
	}

	form, _ := c.MultipartForm()
	files := form.File["file"]

	for _, file := range files {
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

		// delete the file after upload
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
			OwnerID:   userId,
			DatasetID: r.DatasetID,
			Status:    "uploaded",
			IsPublic:  r.IsPublic,
		}

		uploadedId, err := database.FileRepo.CreateFile(fileObject)
		utils.HandleHandlerError(c, err)

		// TODO: Trigger upload processing
		type ProcessorRequest struct {
			FileId string `json:"file_id"`
		}
		request := ProcessorRequest{
			FileId: uploadedId,
		}
		// marshall data to json (like json_encode)
		marshalled, err := json.Marshal(request)

		processorUrl := os.Getenv("PROCESSOR_URL")

		req, err := http.NewRequest("POST", processorUrl+"/process", bytes.NewReader(marshalled))
		utils.HandleHandlerError(c, err)

		httpClient := &http.Client{}
		_, err = httpClient.Do(req)
		utils.HandleHandlerError(c, err)

	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
	return
}
