package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID    uint    `gorm:"index"`
	ProductID uint    `gorm:"index"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int
	Price     float64
}
