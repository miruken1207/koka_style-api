package cart

import (
	"fmt"
	middleware "koka_style/middlewares"
	model "koka_style/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary		Add product to cart
// @Tags		сart
// @Accept		json
// @Produce		json
// @Param		product_id	path		uint	true	"ID of the product to add"
// @Success		200
// @Failure		400
// @Failure		401
// @Failure		404
// @Failure		500
// @Security		ApiKeyAuth
// @Router			/cart/{product_id} [post]
func AddToCart(db *gorm.DB) gin.HandlerFunc {
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

		var product model.Product
		if err := db.First(&product, productID).Error; err != nil {
			log.Printf("Error fetching product with ID %d: %v", productID, err)
			context.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		var cartItem model.Cart
		if err := db.Where("user_id = ? AND product_id = ?", userID, productID).First(&cartItem).Error; err == nil {
			cartItem.Quantity++
			db.Save(&cartItem)
		} else {
			db.Create(&model.Cart{
				UserID:    userID,
				ProductID: productID,
				Quantity:  1,
				Price:     product.Price,
			})
		}

		context.JSON(http.StatusOK, gin.H{"message": "Product added to cart"})
	}
}
