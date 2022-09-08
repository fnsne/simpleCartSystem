package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID   uint
	Products []OrderProduct
	Amount   decimal.Decimal `gorm:"type:decimal(23,5)"`
}

type OrderProduct struct {
	gorm.Model
	CartID    uint
	ProductID uint
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  uint
}
