package apitesting

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Test struct {
	Name   string
	Method string
	Path   string
	Auth   bool
	Body   any

	ResponseCode int
	ResponseBody any
}

func TestServer(t *testing.T, tests []Test, router *gin.Engine) {
	// TODO: Set token to a valid JWT token
	token := "eyJhb"

	for _, tt := range tests {
		w := httptest.NewRecorder()
		t.Run(tt.Name, func(t *testing.T) {
			jsonBody, _ := json.Marshal(tt.Body)
			req, _ := http.NewRequest(tt.Method, tt.Path, bytes.NewBuffer(jsonBody))
			if tt.Auth {
				req.Header.Set("Authorization", "Bearer "+token)
			}
			router.ServeHTTP(w, req)

			if tt.ResponseCode != w.Code {
				t.Errorf("Test: %s - Expected response code %d, got %d", tt.Name, tt.ResponseCode, w.Code)
			}
			// Only check body if we have an OK response
			if w.Code == 200 {
				response := tt.ResponseBody
				_ = json.Unmarshal(w.Body.Bytes(), &response)
				if tt.ResponseBody != w.Body {
					t.Errorf("Test: %s - Expected response body %s, got %s", tt.Name, tt.ResponseBody, w.Body)
				}
			} else {
				var response struct {
					Error string `json:"error"`
				}
				_ = json.Unmarshal(w.Body.Bytes(), &response)
				if response.Error == "" {
					t.Errorf("Test: %s - Expected response body to contain error", tt.Name)
				}
			}
		})
	}
}
