package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID     uint
	Products   []CartProduct   `gorm:"foreignKey:CartID"`
	Amount     decimal.Decimal `gorm:"type:decimal(23,5)"`
	IsCheckout bool
}

func (c *Cart) GetProductIds() []uint {
	var ids []uint
	for _, product := range c.Products {
		ids = append(ids, product.ProductID)
	}
	return ids
}

type CartProduct struct {
	gorm.Model
	CartID    uint
	ProductID uint
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  uint
}

func (c *Cart) toOrderProducts() []OrderProduct {
	var oPs []OrderProduct
	for _, product := range c.Products {
		oPs = append(oPs, OrderProduct{
			ProductID: product.ProductID,
			Quantity:  product.Quantity,
		})
	}
	return oPs
}
func (c *Cart) ToOrder() *Order {
	return &Order{
		UserID:   c.UserID,
		Products: c.toOrderProducts(),
		Amount:   c.Amount,
	}
}
func (c *Cart) CalculateAmount() {
	c.Amount = c.getAmount()
}
func (c *Cart) getAmount() decimal.Decimal {
	var amount decimal.Decimal
	for _, product := range c.Products {
		amount = amount.Add(product.Product.Price.Mul(decimal.NewFromInt(int64(product.Quantity))))
	}
	return amount
}

func (c *Cart) CartHasOrderProduct() bool {
	if len(c.Products) != 0 {
		return true
	}
	return false
}

type AddProduct struct {
	ProductID uint
	Quantity  uint
}
