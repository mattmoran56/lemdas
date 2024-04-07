package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/apitesting"
	"github.com/mattmoran/fyp/api/pkg/web/middleware"
	"testing"
)

func TestHandleUpdateAttribute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	group := router.Group("/", middleware.JWTAuthMiddleware(), middleware.CheckDatasetAccess())
	group.PUT("dataset/:datasetId/attribute/:datasetAttributeId", HandleUpdateAttribute)

	apitesting.AttributeGroupTest(t, apitesting.Request{
		Method:            "PUT",
		Url:               "/dataset/testdataset/attribute/testattribute1",
		Body:              map[string]interface{}{"attribute_name": "testupdate1", "attribute_value": "valueupdate1"},
		Engine:            router,
		ResponseCode:      201,
		ResponseBody:      map[string]interface{}{"id": "", "created_at": "", "updated_at": "", "dataset_id": "testdataset", "attribute_name": "testupdate1", "attribute_value": "valueupdate1"},
		BodyIgnoredFields: []string{"id", "created_at", "updated_at"},
	})
}
