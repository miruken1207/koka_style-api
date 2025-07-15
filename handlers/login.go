package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	config "koka_style/config"
	model "koka_style/models"
)

// Login godoc
//	@Summary	User login
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		credentials	body	model.Login_input	true	"User credentials (username and password)"
//	@Success	200
//	@Failure	400
//	@Failure	401
//	@Failure	500
//	@Router		/login [post]
func Login(db *gorm.DB) gin.HandlerFunc {

	return func(context *gin.Context) {

		var login_input model.Login_input
		if err := context.BindJSON(&login_input); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"}) // 400
			return
		}

		var user model.User
		if err := db.Where("username = ?", login_input.Username).First(&user).Error; err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"}) // 401
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login_input.Password)); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"}) // 401
			return
		}

		expirationTime := time.Now().Add(24 * time.Hour)
		claims := &model.Claims{
			Username: user.Username,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(config.JwtSecret)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"}) // 500
			return
		}

		context.JSON(http.StatusOK, gin.H{"token": tokenString}) // 200
	}
}
