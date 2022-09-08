package repository

import (
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"shopline-question/model"
)

type CartRepo struct {
	db *gorm.DB
}

func (r *CartRepo) GetByUserID(userId int) (cart model.Cart) {
	r.db.Preload("Products.Product").Model(&model.Cart{}).Where("user_id=?", userId).First(&cart)
	return cart
}

func (r *CartRepo) Update(cart model.Cart) error {
	if !PRODUCT.AllExist(cart.GetProductIds()) {
		return errors.New("product not exist")
	}

	err := r.db.Model(&model.Cart{Model: gorm.Model{ID: cart.ID}}).
		Association("Products").
		Replace(cart.Products)
	if err != nil {
		return err
	}
	products := r.getOrderProductsBy(cart.ID)
	amount := decimal.NewFromInt(0)
	for _, orderProduct := range products {
		productAmount := orderProduct.Product.Price.Mul(decimal.NewFromInt(int64(orderProduct.Quantity)))
		amount = amount.Add(productAmount)
	}
	err = r.db.Model(&model.Cart{}).Where("id=?", cart.ID).Update("amount", amount).Error
	return err
}

func (r *CartRepo) getOrderProductsBy(cartID uint) []model.OrderProduct {
	var products []model.OrderProduct
	r.db.Preload("Product").
		Model(&model.OrderProduct{}).
		Where("cart_id=?", cartID).
		Find(&products)
	return products
}

func NewCartRepo(db *gorm.DB) *CartRepo {
	return &CartRepo{db: db}
}
