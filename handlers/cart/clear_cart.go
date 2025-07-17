package cart

import (
	middleware "koka_style/middlewares"
	model "koka_style/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary		Clear user's cart
// @Tags		сart
// @Accept		json
// @Produce		json
// @Security	ApiKeyAuth
// @Success		200
// @Failure		401
// @Failure		500
// @Router			/cart [delete]
func ClearCart(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		userID, err := middleware.GetUserIDFromContext(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		if err := db.Where("user_id = ?", userID).Delete(&model.Cart{}).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not clear cart"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Cart cleared"})
	}
}
