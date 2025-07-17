package cart

import (
	"fmt"
	middleware "koka_style/middlewares"
	model "koka_style/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary		Remove product from cart
// @Tags		сart
// @Accept		json
// @Produce		json
// @Security	ApiKeyAuth
// @Param		product_id	path		uint	true	"Product ID"
// @Success		200
// @Failure		400
// @Failure		401
// @Failure		500
// @Router		/cart/{product_id} [delete]
func RemoveFromCart(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		userID, err := middleware.GetUserIDFromContext(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		productIDStr := context.Param("product_id")
		var productID uint
		if _, err := fmt.Sscanf(productIDStr, "%d", &productID); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}

		if err := db.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&model.Cart{}).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not remove product from cart"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Product removed from cart"})
	}
}
