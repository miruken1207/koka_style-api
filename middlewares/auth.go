package middleware

import (
	"errors"
	config "koka_style/config"
	model "koka_style/models"
	"strings"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(context *gin.Context) {

		tokenString := context.GetHeader("Authorization")
		if tokenString != "" && strings.HasPrefix(tokenString, "Bearer ") {
			tokenString = tokenString[len("Bearer "):]
		} else {
			var err error
			tokenString, err = context.Cookie("token")
			if err != nil {
				context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
				return
			}
		}

		token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return config.JwtSecret, nil
		})

		if err != nil {
			switch {
			case err == jwt.ErrTokenMalformed:
				context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid token format"})
			case err == jwt.ErrTokenExpired:
				context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			default:
				context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			}
			return
		}

		claims, ok := token.Claims.(*model.Claims)
		if !ok {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		if claims.ExpiresAt != nil && claims.ExpiresAt.Unix() < time.Now().Unix() {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			return
		}

		context.Set("username", claims.Username)
		context.Set("userID", claims.ID)

		context.Next()
	}
}

func GetUserIDFromContext(c *gin.Context) (uint, error) {
	userID, exists := c.Get("userID")
	if !exists {
		return 0, errors.New("user ID not found in context")
	}
	if id, ok := userID.(uint); ok {
		return id, nil
	}
	return 0, errors.New("invalid user ID type")
}
