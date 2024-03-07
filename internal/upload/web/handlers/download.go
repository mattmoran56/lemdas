package handlers

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/utils"
	"io"
	"net/http"
)

func HandleDownload(c *gin.Context) {
	fileId := c.Param("fileId")

	file, err := database.FileRepo.GetFileByID(fileId)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// TODO: Check file permissions

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	utils.HandleHandlerError(c, err)

	url := "https://synopticprojectstorage.blob.core.windows.net/"
	client, err := azblob.NewClient(url, credential, nil)
	utils.HandleHandlerError(c, err)

	containerName := "fyp-uploads"

	downloadStream, err := client.DownloadStream(context.TODO(), containerName, fileId, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error downloading file from storage: " + err.Error()})
		return
	}

	defer downloadStream.Body.Close()

	//_, err = io.ReadAll(downloadStream.Body)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading file from storage: " + err.Error()})
	//	return
	//}

	c.Header("Content-Disposition", "attachment; filename="+file.Name)
	c.Header("Content-Type", "application/octet-stream")

	// Copy the file content to the response writer
	_, err = io.Copy(c.Writer, downloadStream.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to copy file content to response")
		return
	}

}
