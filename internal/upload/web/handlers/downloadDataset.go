package handlers

import (
	"archive/zip"
	"bytes"
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/utils"
	"io"
	"net/http"
)

func HandleDownloadDataset(c *gin.Context) {
	datasetID := c.Param("datasetId")

	dataset, err := database.DatasetRepo.GetDatasetByID(datasetID)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Dataset not found"})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Dataset not found"})
		return
	}

	files, err := database.FileRepo.GetFilesForDataset(datasetID)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Dataset not found"})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Dataset not found"})
		return
	}

	zipBuffer := new(bytes.Buffer)
	zipWriter := zip.NewWriter(zipBuffer)

	for _, file := range files {
		credential, err := azidentity.NewDefaultAzureCredential(nil)
		utils.HandleHandlerError(c, err)

		url := "https://synopticprojectstorage.blob.core.windows.net/"
		client, err := azblob.NewClient(url, credential, nil)
		utils.HandleHandlerError(c, err)

		containerName := "fyp-uploads"

		downloadStream, err := client.DownloadStream(context.TODO(), containerName, file.ID, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error downloading file from storage: " + err.Error()})
			return
		}

		defer downloadStream.Body.Close()

		zipFile, err := zipWriter.Create(file.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error building downloadable"})
			return
		}

		_, err = io.Copy(zipFile, downloadStream.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error building downloadable"})
			return
		}
	}

	err = zipWriter.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error building downloadable"})
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+dataset.DatasetName)
	c.Header("Content-Type", "application/octet-stream")

	// Copy the file content to the response writer
	_, err = io.Copy(c.Writer, zipBuffer)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to copy file content to response")
		return
	}

}
