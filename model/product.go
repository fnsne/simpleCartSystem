package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name      string
	Price     decimal.Decimal `gorm:"type:decimal(23,5)"`
	Inventory uint            `gorm:"check:unsignedInventory, inventory >=0"`
}
