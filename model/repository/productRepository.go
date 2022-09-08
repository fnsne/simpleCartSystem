package repository

import (
	"gorm.io/gorm"
	"shopline-question/model"
)

type ProductRepo struct {
	db *gorm.DB
}

func (r *ProductRepo) List() (products []model.Product) {
	r.db.Model(&model.Product{}).Find(&products)
	return products
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) AllExist(productIds []uint) (Exist bool) {
	var count int64
	r.db.Model(&model.Product{}).Where("id in (?)", productIds).Count(&count)
	Exist = int(count) == len(productIds)
	return Exist
}
