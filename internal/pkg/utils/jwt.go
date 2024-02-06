package utils

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type JWTPayload struct {
	UserId    string `json:"user_id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	jwt.StandardClaims
}

func CreateJWT(claims JWTPayload) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))

	claims.ExpiresAt = time.Now().Add(24 * time.Hour).Unix()
	claims.IssuedAt = time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (JWTPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTPayload{}, func(token *jwt.Token) (interface{}, error) {
		secretKey := []byte(os.Getenv("JWT_SECRET"))
		return secretKey, nil
	})

	if err != nil {
		return JWTPayload{}, err
	}

	if claims, ok := token.Claims.(*JWTPayload); ok && token.Valid {
		// Check if the token has expired
		if time.Now().Unix() > claims.ExpiresAt {
			return JWTPayload{}, jwt.NewValidationError("token has expired", jwt.ValidationErrorExpired)
		}
		return *claims, nil
	}

	return JWTPayload{}, jwt.NewValidationError("invalid token", jwt.ValidationErrorMalformed)
}
