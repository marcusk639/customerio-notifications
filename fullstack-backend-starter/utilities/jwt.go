package utilities

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

// GetSecretKey returns the JWT secret key from environment variables
func GetSecretKey() []byte {
	return []byte(os.Getenv("JWT_SECRET_KEY"))
}

// Create a signed JWT token with the given id and email
func CreateJWT(userID string, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"email":   email,
	})

	return token.SignedString(GetSecretKey())
}
