package utils

import (
	"github.com/gin-gonic/gin"
)

func HandleHandlerError(c *gin.Context, err error) {
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		c.Abort()
	}
}
