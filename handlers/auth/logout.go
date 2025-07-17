package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Logout godoc
//
//	@Summary	User logout
//	@Tags		auth
//	@Produce	json
//	@Success	200
//	@Router		/logout [post]
func Logout(db *gorm.DB) gin.HandlerFunc {

	return func(context *gin.Context) {

		context.SetCookie("token", "", -1, "/", "", false, true)

		context.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
	}
}
