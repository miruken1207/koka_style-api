package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	model "koka_style/models"
)

func Login(db *gorm.DB) gin.HandlerFunc {

	return func(context *gin.Context) {

		var creds model.Credentials
		if err := context.BindJSON(&creds); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		var user model.User
		if err := db.Where("username = ? AND password = ?", creds.Username, creds.Password).First(&user).Error; err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"token": "Bearer " + user.Username})
	}
}
