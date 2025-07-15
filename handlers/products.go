package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetProducts godoc
//
//	@Summary	Get list of products
//	@Tags		products
//	@Produce	json
//	@Success	200
//	@Router		/products [get]
func GetProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "List of products"})
	}
}
