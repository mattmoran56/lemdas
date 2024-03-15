package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/apitesting"
	"github.com/mattmoran/fyp/api/pkg/web/middleware"
	"testing"
)

func TestHandleGetAttributes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	group := router.Group("/", middleware.JWTAuthMiddleware(), middleware.CheckDatasetAccess())
	group.GET("dataset/:datasetId/attribute", HandleGetAttributes)

	apitesting.AttributeGroupTest(t, apitesting.Request{
		Method:       "GET",
		Url:          "/dataset/testdataset/attribute",
		Body:         nil,
		Engine:       router,
		ResponseCode: 200,
		ResponseBody: map[string]interface{}{"attributes": []interface{}{
			map[string]interface{}{"id": "testattribute1", "updated_at": 100, "created_at": 100, "dataset_id": "testdataset", "attribute_name": "test1", "attribute_value": "value1"},
			map[string]interface{}{"id": "testattribute2", "updated_at": 100, "created_at": 100, "dataset_id": "testdataset", "attribute_name": "test2", "attribute_value": "value2"},
		}},
		BodyIgnoredFields: nil,
	})
}
