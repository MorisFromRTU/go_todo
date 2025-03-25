package middleware

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"todo-app/auth"
)

var ErrUnauthorized = errors.New("unauthorized")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header required"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" || authHeader == tokenString {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid token format"})
			return
		}

		userID, err := auth.VerifyToken(tokenString)
		if err != nil {
            c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
            return
        }

        c.Set("user_id", userID)
        c.Next()
	}
}
