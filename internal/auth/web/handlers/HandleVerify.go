package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/utils"
	"net/http"
)

func HandleVerify(c *gin.Context) {
	type TokenRequest struct {
		Token string `json:"token"`
	}

	var r TokenRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := utils.VerifyJWT(r.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Token is valid"})
	return
}
