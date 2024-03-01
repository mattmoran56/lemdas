package dataset

import "github.com/gin-gonic/gin"

func HandleGetAccessLevel(c *gin.Context) {
	access := c.GetString("access")

	c.JSON(200, gin.H{"access": access})
	return
}
