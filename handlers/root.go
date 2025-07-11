package handlers

import (
	"koka_style/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Root(db *gorm.DB) gin.HandlerFunc {

	return func(context *gin.Context) {

		authHeader := context.GetHeader("Authorization")
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			username := authHeader[7:]
			var user models.User
			if err := db.Where("username = ?", username).First(&user).Error; err == nil {
				context.String(http.StatusOK, "Welcome, %s!", user.Username)
				return
			}
		}

		context.Redirect(http.StatusTemporaryRedirect, "/login")
	}
}
