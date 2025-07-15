package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	model "koka_style/models"
)

// SignUp godoc
//	@Summary	User registration
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		credentials	body	model.SignUp_input	true	"Registration data (username, email, password)"
//	@Success	200
//	@Failure	400
//	@Failure	500
//	@Router		/sign_up [post]
func SignUp(db *gorm.DB) gin.HandlerFunc {

	return func(context *gin.Context) {

		var signup_input model.SignUp_input
		if err := context.BindJSON(&signup_input); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"}) // 400
			return
		}

		var existing model.User
		if err := db.Where("username = ?", signup_input.Username).Or("email = ?", signup_input.Email).First(&existing).Error; err == nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"}) // 400
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signup_input.Password), bcrypt.DefaultCost)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"}) // 500
			return
		}

		user := model.User{
			Username: signup_input.Username,
			Password: string(hashedPassword),
			Email:    signup_input.Email,
		}

		if err := db.Create(&user).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"}) // 500
			return
		}

		context.JSON(http.StatusOK, gin.H{"status": "User registered successfully"}) // 200
	}
}
