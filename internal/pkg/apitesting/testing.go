package apitesting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type Test struct {
	Name   string
	Method string
	Path   string
	Auth   string
	Body   map[string]interface{}

	ResponseCode      int
	ResponseBody      map[string]interface{}
	ManualCompareBody bool `default:"false"`
}

func mapsEqual(map1, map2 map[string]interface{}) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, val1 := range map1 {
		val2, ok := map2[key]
		if !ok || !reflect.DeepEqual(val1, val2) {
			return false
		}
	}

	return true
}

func TestServer(t *testing.T, tests []Test, router *gin.Engine) {
	// TODO: Set token to a valid JWT token

	for _, tt := range tests {
		w := httptest.NewRecorder()

		if tt.ManualCompareBody {
			tt.Name = fmt.Sprintf("MANUAL CHECK: %s", tt.Name)
		}
		t.Run(tt.Name, func(t *testing.T) {
			jsonBody, _ := json.Marshal(tt.Body)
			req, _ := http.NewRequest(tt.Method, tt.Path, bytes.NewBuffer(jsonBody))
			if tt.Auth != "" {
				req.Header.Set("Authorization", "Bearer "+tt.Auth)
			}
			router.ServeHTTP(w, req)

			if tt.ResponseCode != w.Code {
				t.Errorf("Test: %s - Expected response code %d, got %d", tt.Name, tt.ResponseCode, w.Code)
			}

			// Only check body if we don't have internal server error
			if w.Code != 500 {
				response := make(map[string]interface{})
				_ = json.Unmarshal(w.Body.Bytes(), &response)
				//response = make(map[string]interface{})
				if tt.ManualCompareBody {
					fmt.Printf("Diff: %v\n", cmp.Diff(tt.ResponseBody, response))
					t.Error("Manual compare body set to true, no comparison done.")
				} else if !mapsEqual(tt.ResponseBody, response) {
					fmt.Printf("Diff: %v\n", cmp.Diff(tt.ResponseBody, response))
					t.Errorf("Test: %s - Expected response body %s, got %s", tt.Name, tt.ResponseBody, response)
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
