package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Root godoc
//	@Summary	Root route
//	@Tags		root
//	@Produce	json
//	@Success	200
//	@Router		/ [get]
func Root(db *gorm.DB) gin.HandlerFunc {

	return func(context *gin.Context) {

		context.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Koka Style online store API!",
			"routes": []string{
				"POST /sign_up — register a new user",
				"POST /login — authenticate a user",
			},
		})
	}
}
