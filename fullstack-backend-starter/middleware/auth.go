package middleware

import (
	"net/http"
	"strings"

	"github.com/customerio/fullstack-backend-homework/utilities"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Middleware responsible for parsing the Auth token if present on the Authorization header.
// Once parsed, it injects the claims into the Gin context.
//
// If no token is present, it will simply pass the request to the next handler.
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")

		if authorizationHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		tokenString := strings.Split(authorizationHeader, "Bearer ")[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}

			return utilities.GetSecretKey(), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("claims", claims)
			if userID, ok := claims["user_id"].(string); ok {
				c.Set("user_id", userID)
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}

		c.Next()
	}
}
