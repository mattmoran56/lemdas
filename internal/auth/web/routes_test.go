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
			Auth:         false,
			Body:         nil,
			ResponseCode: 400,
			ResponseBody: nil,
		},
		{
			Name:   "Invalid code provided",
			Method: "POST",
			Path:   "/token",
			Auth:   false,
			Body: struct {
				Code string `json:"code"`
			}{Code: "invalid"},
			ResponseCode: 500,
			ResponseBody: nil,
		},

		{
			Name:         "No Token provided",
			Method:       "POST",
			Path:         "/verify",
			Auth:         false,
			Body:         nil,
			ResponseCode: 401,
			ResponseBody: nil,
		},
		{
			Name:   "Invalid Token provided",
			Method: "POST",
			Path:   "/verify",
			Auth:   false,
			Body: struct {
				Token string `json:"token"`
			}{Token: "invalid"},
			ResponseCode: 401,
			ResponseBody: nil,
		},
		{
			Name:   "JWT signed with different key provided",
			Method: "POST",
			Path:   "/verify",
			Auth:   false,
			Body: struct {
				Token string `json:"token"`
			}{Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNWRiZDFhMTUtNmU2Ni00YmQzLTgwNWItYTAzMWUzY2MzMDA2IiwiZW1haWwiOiJzYzIwbW1AbGVlZHMuYWMudWsiLCJmaXJzdF9uYW1lIjoiTWF0dGhldyIsImxhc3RfbmFtZSI6Ik1vcmFuIiwiZXhwIjoxNzA4NTk3ODkzLCJpYXQiOjE3MDg1MTE0OTN9.yBsbUHsH-mo8Hj_9zynYwgeGPJ3Q70N8w0w-_tdF290"},
			ResponseCode: 401,
			ResponseBody: nil,
		},
	}

	gin.SetMode(gin.TestMode)
	router := InitiateServer()

	apitesting.TestServer(t, tests, router)
}
