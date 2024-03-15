package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/apitesting"
	"github.com/mattmoran/fyp/api/pkg/web/middleware"
	"testing"
)

func TestHandleGetDataset(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	group := router.Group("/", middleware.JWTAuthMiddleware(), middleware.CheckDatasetAccess())
	group.GET("dataset/:datasetId", HandleGetDataset)

	apitesting.DatasetGroupTest(t, apitesting.Request{
		Method:       "GET",
		Url:          "/dataset/testdataset",
		Body:         nil,
		Engine:       router,
		ResponseCode: 200,
		ResponseBody: map[string]interface{}{
			"id":           "testdataset",
			"created_at":   float64(100),
			"updated_at":   float64(100),
			"dataset_name": "test",
			"is_public":    false,
			"owner_id":     "testuserowner",
			"owner_name":   "test testson",
		},
		BodyIgnoredFields: []string{"is_public"},
	})
}
