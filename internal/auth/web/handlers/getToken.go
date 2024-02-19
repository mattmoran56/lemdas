package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"github.com/mattmoran/fyp/api/pkg/utils"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
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

	decodedResp := TokenResponse{}
	err = json.NewDecoder(resp.Body).Decode(&decodedResp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	}

	user, err := database.UserRepo.GetUserByEmail(claims.Email)
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
