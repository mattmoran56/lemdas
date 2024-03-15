package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/apitesting"
	"github.com/mattmoran/fyp/api/pkg/web/middleware"
	"testing"
)

func TestHandleGetCollaborators(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	group := router.Group("/", middleware.JWTAuthMiddleware(), middleware.CheckDatasetAccess())
	group.GET("dataset/:datasetId/collaborators", HandleGetCollaborators)

	apitesting.DatasetGroupTest(
		t,
		apitesting.Request{
			Method:       "GET",
			Url:          "/dataset/testdataset/collaborators",
			Body:         nil,
			Engine:       router,
			ResponseCode: 200,
			ResponseBody: map[string]interface{}{
				"collaborators": []interface{}{
					map[string]interface{}{"id": "", "updated_at": 0, "created_at": 0, "dataset_id": "testdataset", "user": map[string]interface{}{
						"id":         "testusernotowner",
						"updated_at": 0,
						"created_at": 0,
						"email":      "test@test.com",
						"first_name": "test",
						"last_name":  "testson",
					}},
				},
			},
			BodyIgnoredFields: []string{"updated_at", "created_at", "id"},
		},
	)
}
