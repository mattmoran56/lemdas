package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"net/http"
)

func CheckFileAttributeAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		fileID := c.Param("fileId")
		fileAttributeID := c.Param("fileAttributeId")

		fileAttribute, err := database.FileAttributeRepo.GetFileAttributeByID(fileAttributeID)
		if err != nil {
			if err.Error() == "record not found" {
				c.JSON(http.StatusNotFound, gin.H{"error": "Attribute not found"})
				c.Abort()
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			c.Abort()
			return
		}

		if fileAttribute.FileID == fileID {
			c.Next()
			return
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Attribute not found"})
			c.Abort()
			return
		}
	}
}
