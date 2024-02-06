package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/utils"
	"net/http"
	"strings"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")

		// Check if the Authorization header is empty or doesn't start with "Bearer "
		if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "JWT token is missing or in an invalid format"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "JWT token is missing"})
			c.Abort()
			return
		}

		claims, err := utils.VerifyJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Pass the userID to the request context for later use
		c.Set("userID", claims.UserId)

		c.Next()
	}
}
