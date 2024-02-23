package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"net/http"
)

func CheckDatasetAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		datasetID := c.Param("datasetId")

		dataset, err := database.DatasetRepo.GetDatasetByID(datasetID)
		if err != nil {
			if err.Error() == "record not found" {
				c.JSON(http.StatusNotFound, gin.H{"error": "Dataset not found"})
				c.Abort()
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			c.Abort()
			return
		}

		if dataset.OwnerID == c.GetString("userID") {
			c.Next()
			return
		} else {
			if dataset.IsPublic {
				if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
					c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
					c.Abort()
					return
				}
				c.Next()
				return
			} else {
				c.JSON(http.StatusNotFound, gin.H{"error": "Dataset not found"})
				c.Abort()
				return
			}
		}
	}
}
