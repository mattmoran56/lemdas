package handlers

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/service"
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"github.com/mattmoran/fyp/api/pkg/utils"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"
)

type Claims struct {
	Email     string `json:"unique_name"`
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
}

func HandleToken(c *gin.Context) {
	type TokenRequest struct {
		Code string `json:"code"`
	}

	type TokenResponse struct {
		AccessToken string `json:"access_token"`
	}

	var r TokenRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if r.Code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	defer writer.Close()

	// Write form fields to get token
	writer.WriteField("code", r.Code)
	writer.WriteField("client_id", os.Getenv("CLIENT_ID"))
	writer.WriteField("client_secret", os.Getenv("CLIENT_SECRET"))
	writer.WriteField("grant_type", "authorization_code")

	writer.Close()
	// Set the Content-Type header
	req, err := http.NewRequest("POST", "https://login.microsoftonline.com/bdeaeda8-c81d-45ce-863e-5232a535b7cb/oauth2/v2.0/token", body)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the Content-Type header with the boundary from the multipart writer
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Make the request to get token
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if resp.StatusCode != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting token"})
		return
	}

	decodedResp := TokenResponse{}
	err = json.NewDecoder(resp.Body).Decode(&decodedResp)

	if decodedResp.AccessToken == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No access token returned"})
		return
	}

	parts := strings.Split(decodedResp.AccessToken, ".")
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var claims Claims
	err = json.Unmarshal(payload, &claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	checkUserExists, err := database.UserRepo.CheckUserExistsByEmail(claims.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if checkUserExists == false {
		user := models.User{
			Email:     claims.Email,
			FirstName: claims.FirstName,
			LastName:  claims.LastName,
		}
		err = database.UserRepo.CreateUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = getAvatar(user, decodedResp.AccessToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

	}

	user, err := database.UserRepo.GetUserByEmail(claims.Email)
	if user.Avatar == "" {
		err = getAvatar(user, decodedResp.AccessToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	jwtClaims := utils.JWTPayload{
		UserId:    user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		// Add extra details such as research groups etc here
	}

	token, err := utils.CreateJWT(jwtClaims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token})

}

func getAvatar(user models.User, accessToken string) error {
	req, err := http.NewRequest("GET", "https://graph.microsoft.com/v1.0/me/photo/$value", nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		filePath := ".temp/" + user.ID + ".jpg" // Specify the file path to save the avatar image
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		// Copy the response body (avatar image data) to the file
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return err
		}

		credential, err := azidentity.NewDefaultAzureCredential(nil)
		if err != nil {
			return err
		}

		url := "https://synopticprojectstorage.blob.core.windows.net/"
		client, err := azblob.NewClient(url, credential, nil)
		if err != nil {
			return err
		}

		containerName := "fyp-profile-images"

		// Open the file to upload
		fileHandler, err := os.Open(".temp/" + user.ID + ".jpg")
		if err != nil {
			return err
		}

		// close the file after it is no longer required.
		defer func(file *os.File) {
			err = file.Close()
			if err != nil {
				return
			}
		}(fileHandler)

		// delete the file after upload
		defer func(name string) {
			err = os.Remove(name)
			if err != nil {
				return
			}
		}(".temp/" + user.ID + ".jpg")

		_, err = client.UploadFile(context.TODO(), containerName, user.ID+".jpg", fileHandler, &azblob.UploadBufferOptions{})
		if err != nil {
			return err
		}

		containerName = "fyp-profile-images"

		cred, err := azidentity.NewDefaultAzureCredential(nil)
		if err != nil {
			return err
		}

		// Get the azure blob client
		client, err = azblob.NewClient(url, cred, nil)
		if err != nil {
			return err
		}

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
		if err != nil {
			return err
		}

		sasQueryParams, err := sas.BlobSignatureValues{
			Protocol:      sas.ProtocolHTTPS,
			StartTime:     time.Now().UTC().Add(time.Second * -10),
			ExpiryTime:    time.Now().UTC().Add(1000 * time.Hour),
			Permissions:   to.Ptr(sas.ContainerPermissions{Read: true, List: true}).String(),
			ContainerName: containerName,
			BlobName:      user.ID + ".jpg",
		}.SignWithUserDelegation(udc)
		if err != nil {
			return err
		}
		sasURL := url + containerName + "/" + user.ID + ".jpg" + "?" + sasQueryParams.Encode()

		user.Avatar = sasURL

		err = database.UserRepo.UpdateUser(user)
		if err != nil {
			return err
		}

	}

	return nil
}
