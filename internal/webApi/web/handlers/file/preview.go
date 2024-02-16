package file

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/service"
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/utils"
	"strings"
	"time"
)

func HandlePreview(c *gin.Context) {
	fileId := c.Param("fileId")

	// Check user has permission to view file
	userId := c.MustGet("userID").(string)
	if utils.CheckUserHasPermission(fileId, userId) == false {
		c.JSON(401, gin.H{"error": "User does not have permission to view this file"})
		return
	}

	file := strings.Split(fileId, ".")[0] + ".png"

	containerName := "fyp-previews"
	storageUrl := "https://synopticprojectstorage.blob.core.windows.net/"

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	utils.HandleHandlerError(c, err)

	// Get the azure blob client
	client, err := azblob.NewClient(storageUrl, cred, nil)
	utils.HandleHandlerError(c, err)

	serviceClient := client.ServiceClient()

	// Prepare SAS Token Input field
	now := time.Now().UTC().Add(-10 * time.Second)
	expiry := now.Add(48 * time.Hour)
	info := service.KeyInfo{
		Start:  to.Ptr(now.UTC().Format(sas.TimeFormat)),
		Expiry: to.Ptr(expiry.UTC().Format(sas.TimeFormat)),
	}

	// Get Delegation Key
	udc, err := serviceClient.GetUserDelegationCredential(context.TODO(), info, nil)
	utils.HandleHandlerError(c, err)

	sasQueryParams, err := sas.BlobSignatureValues{
		Protocol:      sas.ProtocolHTTPS,
		StartTime:     time.Now().UTC().Add(time.Second * -10),
		ExpiryTime:    time.Now().UTC().Add(15 * time.Minute),
		Permissions:   to.Ptr(sas.ContainerPermissions{Read: true, List: true}).String(),
		ContainerName: containerName,
		BlobName:      file,
	}.SignWithUserDelegation(udc)
	utils.HandleHandlerError(c, err)

	sasURL := storageUrl + containerName + "/" + file + "?" + sasQueryParams.Encode()

	//// This URL can be used to authenticate requests now
	//azClient, err := azblob.NewClientWithNoCredential(sasURL, nil)
	//utils.HandleHandlerError(c, err)

	c.JSON(200, gin.H{"url": sasURL})
}
