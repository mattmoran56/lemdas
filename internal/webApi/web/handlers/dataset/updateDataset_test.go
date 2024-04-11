package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/apitesting"
	"github.com/mattmoran/fyp/api/pkg/web/middleware"
	"testing"
)

func TestHandleUpdateDataset(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	group := router.Group("/", middleware.JWTAuthMiddleware(), middleware.CheckDatasetAccess())
	group.PUT("dataset/:datasetId", HandleUpdateDataset)

	if

	apitesting.DatasetGroupTest(t, apitesting.Request{
		Method:       "PUT",
		Url:          "/dataset/testdataset",
		Body:         map[string]interface{}{"dataset_name": "testnew", "is_public": false, "owner_id": "testuserowner"},
		Engine:       router,
		ResponseCode: 201,
		ResponseBody: map[string]interface{}{
			"id":           "testdataset",
			"created_at":   float64(100),
			"updated_at":   float64(100),
			"dataset_name": "testnew",
			"is_public":    false,
			"owner_id":     "testuserowner",
		},
		BodyIgnoredFields: []string{"created_at", "updated_at"},
	})
}
