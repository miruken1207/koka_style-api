package handlers

import (
	"net/http"

	model "koka_style/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetProducts godoc
//
//	@Summary	Get list of products
//	@Tags		products
//	@Produce	json
//	@Success	200
//	@Failure	401
//	@Failure	500
//  @Security	ApiKeyAuth
//	@Router		/products [get]
func GetProducts(db *gorm.DB) gin.HandlerFunc {

	return func(context *gin.Context) {

		var products []model.Product
		if err := db.Find(&products).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Unable retrieve products"}) // 500
		}

		context.JSON(http.StatusOK, gin.H{"List of products": products}) // 200
	}
}
