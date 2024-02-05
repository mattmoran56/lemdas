package utils

import (
	"github.com/dgrijalva/jwt-go"
	"os"
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
