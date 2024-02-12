package web

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/web/middleware"
	"github.com/mattmoran/fyp/api/webApi/web/handlers/dataset"
	"go.uber.org/zap"
)

func InitiateServer() {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	authGroup := r.Group("/", middleware.JWTAuthMiddleware())
	{
		datasetGroup := authGroup.Group("/dataset")
		{
			datasetGroup.GET("", dataset.HandleGetDatasets)
			datasetGroup.GET("/", dataset.HandleGetDatasets)

			datasetGroup.GET("/:datasetId", dataset.HandleGetDataset)
			datasetGroup.GET("/:datasetId/", dataset.HandleGetDataset)

			datasetGroup.GET("/:datasetId/files", dataset.HandleGetFiles)
			datasetGroup.GET("/:datasetId/files/", dataset.HandleGetFiles)

			datasetGroup.GET("/:datasetId/attributes", dataset.HandleGetAttributes)
			datasetGroup.GET("/:datasetId/attributes/", dataset.HandleGetAttributes)

			datasetGroup.POST("", dataset.HandleCreateDataset)
			datasetGroup.POST("/", dataset.HandleCreateDataset)

			datasetGroup.POST("/:datasetId/attribute", dataset.HandleCreateAttribute)
			datasetGroup.POST("/:datasetId/attribute/", dataset.HandleCreateAttribute)
		}
	}

	err := r.Run(":8003")
	if err != nil {
		zap.S().Fatal("Couldn't start server")
	}

	zap.S().Info("Web API Server started on port 8003")
}
