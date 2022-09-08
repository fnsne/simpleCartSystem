package repository

import (
	"gorm.io/gorm"
	"shopline-question/model"
)

var CART *CartRepo
var PRODUCT model.ProductRepository

func Initial(db *gorm.DB) {
	CART = NewCartRepo(db)
	PRODUCT = NewProductRepo(db)
}
