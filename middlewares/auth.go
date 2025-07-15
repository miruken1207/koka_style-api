package middlewares

import (
	"net/http"

	config "koka_style/config"
	model "koka_style/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JWTAuthMiddleware() gin.HandlerFunc {

	return func(context *gin.Context) {

		tokenString := context.GetHeader("Authorization")

		if tokenString == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			context.Abort()
			return
		}

		tokenString = tokenString[len("Bearer "):]

		claims := &model.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return config.JwtSecret, nil
		})

		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			context.Abort()
			return
		}

		context.Set("username", claims.Username)

		context.Next()
	}
}
