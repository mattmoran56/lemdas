package apitesting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Request struct {
	Method string
	Url    string
	Body   any
	Engine *gin.Engine

	ResponseCode      int
	ResponseBody      map[string]interface{}
	BodyIgnoredFields []string
}

func AuthGroupTest(t *testing.T, r Request, jwt string) {
	// with and without jwt
	for auth := 0; auth < 2; auth++ {
		// with and without body
		for bod := 0; bod < 2; bod++ {
			fmt.Printf("Testing [authGroup] - auth: %v bod: %v\n", auth, bod)
			fmt.Printf("%v", r)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(r.Method, r.Url, nil)
			if bod == 1 && r.Body != nil {
				jsonBody, _ := json.Marshal(r.Body)
				req, _ = http.NewRequest(r.Method, r.Url, bytes.NewBuffer(jsonBody))
			}
			if r.Body == nil {
				bod = 1
			}
			if auth == 1 {
				req.Header.Set("Authorization", "Bearer "+jwt)
			}

			r.Engine.ServeHTTP(w, req)
			unmarshalledBody := make(map[string]interface{})
			json.Unmarshal(w.Body.Bytes(), &unmarshalledBody)

			if auth == 0 {
				assert.Equal(t, 401, w.Code)
				assert.Equal(t, "JWT token is missing or in an invalid format", unmarshalledBody["error"].(string))
			}

			if auth == 1 && bod == 0 && r.Body != nil && (r.ResponseCode == 200 || r.ResponseCode == 201 || r.ResponseCode == 204) {
				assert.Equal(t, 400, w.Code)
			}

			if auth == 1 && bod == 1 {
				assert.Equal(t, r.ResponseCode, w.Code)
				if r.BodyIgnoredFields != nil {
					for _, field := range r.BodyIgnoredFields {
						deleteKey(&unmarshalledBody, field)
						deleteKey(&r.ResponseBody, field)
					}
				}
				if r.ResponseBody != nil {
					assert.Equal(t, r.ResponseBody, unmarshalledBody)
				} else {
					assert.Equal(t, map[string]interface{}{}, unmarshalledBody)
				}
			}
		}
	}
}

func deleteKey(m *map[string]interface{}, key string) {
	for k, v := range *m {
		if k == key {
			delete(*m, k)
		} else {
			switch v.(type) {
			case map[string]interface{}:
				newMap := v.(map[string]interface{})
				deleteKey(&newMap, key)
				v = newMap
			case []interface{}:
				for _, i := range v.([]interface{}) {
					switch i.(type) {
					case map[string]interface{}:
						newMap := i.(map[string]interface{})
						deleteKey(&newMap, key)
						i = newMap
					}
				}
			}
		}
	}
}
