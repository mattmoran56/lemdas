package web

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/apitesting"
	"testing"
)

func TestInitiateServer(t *testing.T) {
	tests := []apitesting.Test{
		{
			Name:         "No code provided",
			Method:       "POST",
			Path:         "/token",
			Auth:         "",
			Body:         nil,
			ResponseCode: 400,
			ResponseBody: map[string]interface{}{"error": "Invalid request body"},
		},
		{
			Name:         "Invalid code provided",
			Method:       "POST",
			Path:         "/token",
			Auth:         "",
			Body:         map[string]interface{}{"code": "invalid"},
			ResponseCode: 500,
			ResponseBody: nil,
		},

		{
			Name:         "No Token provided",
			Method:       "POST",
			Path:         "/verify",
			Auth:         "",
			Body:         nil,
			ResponseCode: 401,
			ResponseBody: map[string]interface{}{"error": "Invalid or expired token"},
		},
		{
			Name:         "Invalid Token provided",
			Method:       "POST",
			Path:         "/verify",
			Auth:         "",
			Body:         map[string]interface{}{"token": "invalid"},
			ResponseCode: 401,
			ResponseBody: map[string]interface{}{"error": "Invalid or expired token"},
		},
		{
			Name:         "JWT signed with different key provided",
			Method:       "POST",
			Path:         "/verify",
			Auth:         "",
			Body:         map[string]interface{}{"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNWRiZDFhMTUtNmU2Ni00YmQzLTgwNWItYTAzMWUzY2MzMDA2IiwiZW1haWwiOiJzYzIwbW1AbGVlZHMuYWMudWsiLCJmaXJzdF9uYW1lIjoiTWF0dGhldyIsImxhc3RfbmFtZSI6Ik1vcmFuIiwiZXhwIjoxNzA4NTk3ODkzLCJpYXQiOjE3MDg1MTE0OTN9.yBsbUHsH-mo8Hj_9zynYwgeGPJ3Q70N8w0w-_tdF290"},
			ResponseCode: 401,
			ResponseBody: map[string]interface{}{"error": "Invalid or expired token"},
		},
	}

	gin.SetMode(gin.TestMode)
	router := InitiateServer()

	apitesting.TestServer(t, tests, router)
}
