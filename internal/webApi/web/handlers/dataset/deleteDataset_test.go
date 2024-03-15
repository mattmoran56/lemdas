package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/apitesting"
	"github.com/mattmoran/fyp/api/pkg/web/middleware"
	"testing"
)

func TestHandleDeleteDataset(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	group := router.Group("/", middleware.JWTAuthMiddleware(), middleware.CheckDatasetAccess())
	group.DELETE("dataset/:datasetId", HandleDeleteDataset)

	apitesting.DatasetGroupTest(t, apitesting.Request{
		Method:            "DELETE",
		Url:               "/dataset/testdataset",
		Body:              nil,
		Engine:            router,
		ResponseCode:      204,
		ResponseBody:      nil,
		BodyIgnoredFields: nil,
	})
}
