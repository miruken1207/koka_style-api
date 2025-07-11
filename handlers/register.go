package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	model "koka_style/models"
)

// Register godoc
// @Summary      Регистрация пользователя
// @Description  Регистрирует нового пользователя по username и password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body   model.Credentials  true  "Данные для регистрации"
// @Success      200 {object} map[string]string "registered"
// @Failure      400 {object} map[string]string "Invalid input или User already exists"
// @Router       /register [post]
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
