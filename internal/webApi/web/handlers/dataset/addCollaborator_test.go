package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/apitesting"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/web/middleware"
	"testing"
)

func TestHandleAddCollaborator(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	group := router.Group("/", middleware.JWTAuthMiddleware(), middleware.CheckDatasetAccess())
	group.POST("dataset/:datasetId/collaborators", HandleAddCollaborator)

	apitesting.DatasetGroupTest(t, apitesting.Request{
		Method:       "POST",
		Url:          "/dataset/testdataset/collaborators",
		Body:         map[string]interface{}{"user_id": "testusercollaborator"},
		Engine:       router,
		ResponseCode: 201,
		ResponseBody: map[string]interface{}{"id": "", "updated_at": 0, "created_at": 0, "dataset_id": "testdataset", "user_id": "testusercollaborator", "user": map[string]interface{}{
			"id":         "testusercollaborator",
			"updated_at": float64(100),
			"created_at": float64(100),
			"email":      "test@test.com",
			"first_name": "test",
			"last_name":  "testson",
		}},
		BodyIgnoredFields: []string{"updated_at", "created_at", "id"},
	})
	database.DatasetCollaboratorRepo.RemoveCollaborator("testdataset", "testusercollaborator")
}
