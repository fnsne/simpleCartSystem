package repository

import (
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

func NewCartRepo(db *gorm.DB) *CartRepo {
	return &CartRepo{db: db}
}
