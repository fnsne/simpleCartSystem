package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID   uint
	Products []OrderProduct  `gorm:"foreignKey:OrderID"`
	Amount   decimal.Decimal `gorm:"type:decimal(23,5)"`
}

type OrderProduct struct {
	gorm.Model
	OrderID   uint
	ProductID uint
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  uint
}
