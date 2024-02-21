package web

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

type test struct {
	name   string
	method string
	path   string
	auth   bool
	body   any

	responseCode int
	responseBody any
}

func TestInitiateServer(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// TODO: Set token to a valid JWT token
	token := "eyJhb"

	tests := []test{
		{
			name:         "No code provided",
			method:       "POST",
			path:         "/token",
			auth:         false,
			body:         nil,
			responseCode: 400,
			responseBody: nil,
		},
		{
			name:   "Invalid code provided",
			method: "POST",
			path:   "/token",
			auth:   false,
			body: struct {
				Code string `json:"code"`
			}{Code: "invalid"},
			responseCode: 500,
			responseBody: nil,
		},

		{
			name:         "No Token provided",
			method:       "POST",
			path:         "/verify",
			auth:         false,
			body:         nil,
			responseCode: 401,
			responseBody: nil,
		},
		{
			name:   "Invalid Token provided",
			method: "POST",
			path:   "/verify",
			auth:   false,
			body: struct {
				Token string `json:"token"`
			}{Token: "invalid"},
			responseCode: 401,
			responseBody: nil,
		},
		{
			name:   "JWT signed with different key provided",
			method: "POST",
			path:   "/verify",
			auth:   false,
			body: struct {
				Token string `json:"token"`
			}{Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNWRiZDFhMTUtNmU2Ni00YmQzLTgwNWItYTAzMWUzY2MzMDA2IiwiZW1haWwiOiJzYzIwbW1AbGVlZHMuYWMudWsiLCJmaXJzdF9uYW1lIjoiTWF0dGhldyIsImxhc3RfbmFtZSI6Ik1vcmFuIiwiZXhwIjoxNzA4NTk3ODkzLCJpYXQiOjE3MDg1MTE0OTN9.yBsbUHsH-mo8Hj_9zynYwgeGPJ3Q70N8w0w-_tdF290"},
			responseCode: 401,
			responseBody: nil,
		},
	}

	router := InitiateServer()

	for _, tt := range tests {
		w := httptest.NewRecorder()
		t.Run(tt.name, func(t *testing.T) {
			jsonBody, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest(tt.method, tt.path, bytes.NewBuffer(jsonBody))
			if tt.auth {
				req.Header.Set("Authorization", "Bearer "+token)
			}
			router.ServeHTTP(w, req)

			if tt.responseCode != w.Code {
				t.Errorf("Test: %s - Expected response code %d, got %d", tt.name, tt.responseCode, w.Code)
			}
			// Only check body if we have an OK response
			if w.Code == 200 {
				response := tt.responseBody
				_ = json.Unmarshal(w.Body.Bytes(), &response)
				if tt.responseBody != w.Body {
					t.Errorf("Test: %s - Expected response body %s, got %s", tt.name, tt.responseBody, w.Body)
				}
			} else {
				var response struct {
					Error string `json:"error"`
				}
				_ = json.Unmarshal(w.Body.Bytes(), &response)
				if response.Error == "" {
					t.Errorf("Test: %s - Expected response body to contain error", tt.name)
				}
			}
		})
	}
}
