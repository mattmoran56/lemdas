package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/apitesting"
	"github.com/mattmoran/fyp/api/pkg/web/middleware"
	"testing"
)

func TestHandleGetFiles(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	group := router.Group("/", middleware.JWTAuthMiddleware(), middleware.CheckDatasetAccess())
	group.GET("dataset/:datasetId/files", HandleGetFiles)

	apitesting.DatasetGroupTest(
		t,
		apitesting.Request{
			Method:       "GET",
			Url:          "/dataset/testdataset/files",
			Body:         nil,
			Engine:       router,
			ResponseCode: 200,
			ResponseBody: map[string]interface{}{"files": []interface{}{
				map[string]interface{}{"id": "testfile1", "created_at": float64(100), "updated_at": float64(100), "name": "Test file1.tif", "owner_id": "testuserowner", "dataset_id": "testdataset", "status": "processed"},
				map[string]interface{}{"id": "testfile2", "created_at": float64(100), "updated_at": float64(100), "name": "Test file2.tif", "owner_id": "testuserowner", "dataset_id": "testdataset", "status": "processed"},
			}},
			BodyIgnoredFields: nil,
		},
	)
}
