package repository

import (
	"gorm.io/gorm"
	"shopline-question/model"
)

var CART model.CartRepository
var PRODUCT model.ProductRepository

func Initial(db *gorm.DB) {
	CART = NewCartRepo(db)
	PRODUCT = NewProductRepo(db)
}
