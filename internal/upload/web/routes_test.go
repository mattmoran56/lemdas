package web

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/apitesting"
	"github.com/mattmoran/fyp/api/pkg/utils"
	"testing"
)

func TestInitiateServer(t *testing.T) {
	token, _ := utils.CreateJWT(utils.JWTPayload{
		UserId:         "test",
		Email:          "",
		FirstName:      "",
		LastName:       "",
		StandardClaims: jwt.StandardClaims{},
	})

	tests := []apitesting.Test{
		{
			Name:         "No token provided and no body",
			Method:       "POST",
			Path:         "/upload",
			Auth:         "",
			Body:         nil,
			ResponseCode: 401,
			ResponseBody: map[string]interface{}{"error": "JWT token is missing or in an invalid format"},
		},
		{
			Name:         "No body provided",
			Method:       "POST",
			Path:         "/upload",
			Auth:         token,
			Body:         nil,
			ResponseCode: 400,
			ResponseBody: map[string]interface{}{"error": "Dataset or is_public is missing"},
		},
	}

	gin.SetMode(gin.TestMode)
	router := InitiateServer()

	apitesting.TestServer(t, tests, router)
}
