package repository

import (
	"gorm.io/gorm"
	"shopline-question/model"
)

type ProductRepo struct {
	db *gorm.DB
}

func (r *ProductRepo) List() []model.Product {
	var products []model.Product
	r.db.Model(&model.Product{}).Find(&products)
	return products
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db: db}
}
