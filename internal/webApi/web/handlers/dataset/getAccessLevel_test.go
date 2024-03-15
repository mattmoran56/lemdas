package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/apitesting"
	"github.com/mattmoran/fyp/api/pkg/web/middleware"
	"testing"
)

func TestHandleGetAccessLevel(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	group := router.Group("/", middleware.JWTAuthMiddleware(), middleware.CheckDatasetAccess())
	group.GET("dataset/:datasetId/access", HandleGetAccessLevel)

	apitesting.DatasetGroupTest(t, apitesting.Request{
		Method:            "GET",
		Url:               "/dataset/testdataset/access",
		Body:              nil,
		Engine:            router,
		ResponseCode:      200,
		ResponseBody:      map[string]interface{}{"access": "value"},
		BodyIgnoredFields: nil,
	})
}
