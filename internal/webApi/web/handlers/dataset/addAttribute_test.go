package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/apitesting"
	"github.com/mattmoran/fyp/api/pkg/web/middleware"
	"testing"
)

func TestHandleCreateAttribute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	group := router.Group("/", middleware.JWTAuthMiddleware(), middleware.CheckDatasetAccess())
	group.POST("dataset/:datasetId/attribute", HandleCreateAttribute)

	apitesting.DatasetGroupTest(t, apitesting.Request{
		Method:            "POST",
		Url:               "/dataset/testdataset/attribute",
		Body:              map[string]interface{}{"attribute_name": "testadded1", "attribute_value": "valueadded1"},
		Engine:            router,
		ResponseCode:      201,
		ResponseBody:      map[string]interface{}{"id": "", "created_at": "", "updated_at": "", "dataset_id": "testdataset", "attribute_name": "testadded1", "attribute_value": "valueadded1"},
		BodyIgnoredFields: []string{"id", "created_at", "updated_at"},
	})
}
