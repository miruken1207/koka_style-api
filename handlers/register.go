package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	model "koka_style/models"
)

func Register(db *gorm.DB) gin.HandlerFunc {

	return func(context *gin.Context) {

		var creds model.Credentials
		if err := context.BindJSON(&creds); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		user := model.User{Username: creds.Username, Password: creds.Password}
		if err := db.Create(&user).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"status": "registered"})
	}
}
