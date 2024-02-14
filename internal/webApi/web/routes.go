package web

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/web/middleware"
	"github.com/mattmoran/fyp/api/webApi/web/handlers/dataset"
	"github.com/mattmoran/fyp/api/webApi/web/handlers/file"
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

			datasetGroup.POST("", dataset.HandleCreateDataset)
			datasetGroup.POST("/", dataset.HandleCreateDataset)

			datasetAttributeGroup := datasetGroup.Group("/:datasetId/attribute")
			{
				datasetAttributeGroup.GET("", dataset.HandleGetAttributes)
				datasetAttributeGroup.GET("/", dataset.HandleGetAttributes)

				datasetAttributeGroup.POST("", dataset.HandleCreateAttribute)
				datasetAttributeGroup.POST("/", dataset.HandleCreateAttribute)

				datasetAttributeGroup.PUT("", dataset.HandleUpdateAttribute)
				datasetAttributeGroup.PUT("/", dataset.HandleUpdateAttribute)

				datasetAttributeGroup.DELETE("/:datasetAttributeId", dataset.HandleDeleteAttribute)
				datasetAttributeGroup.DELETE("/:datasetAttributeId/", dataset.HandleDeleteAttribute)
			}
		}

		fileGroup := authGroup.Group("/file")
		{
			fileGroup.GET("/:fileId", file.HandleGetFile)
			fileGroup.GET("/:fileId/", file.HandleGetFile)

			fileAttributeGroup := fileGroup.Group("/:fileId/attribute")
			{
				fileAttributeGroup.GET("", file.HandleGetAttributes)
				fileAttributeGroup.GET("/", file.HandleGetAttributes)

				fileAttributeGroup.POST("", file.HandleCreateAttribute)
				fileAttributeGroup.POST("/", file.HandleCreateAttribute)

				fileAttributeGroup.PUT("", file.HandleUpdateAttribute)
				fileAttributeGroup.PUT("/", file.HandleUpdateAttribute)

				fileAttributeGroup.DELETE("/:fileAttributeId", file.HandleDeleteAttribute)
				fileAttributeGroup.DELETE("/:fileAttributeId/", file.HandleDeleteAttribute)
			}
		}
	}

	err := r.Run(":8003")
	if err != nil {
		zap.S().Fatal("Couldn't start server")
	}

	zap.S().Info("Web API Server started on port 8003")
}
