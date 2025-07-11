package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	model "koka_style/models"
)

// Login godoc
// @Summary      Авторизация пользователя
// @Description  Логин по username и password, возвращает Bearer-токен при успехе
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body   model.Credentials  true  "Данные для входа"
// @Success      200 {object} map[string]string "Bearer token"
// @Failure      400 {object} map[string]string "Invalid input"
// @Failure      401 {object} map[string]string "Invalid credentials"
// @Router       /login [post]
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
