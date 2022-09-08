package tests

import (
	"gorm.io/gorm"
	"shopline-question/config"
	"shopline-question/model"
)

func GivenProducts(products []model.Product) *gorm.DB {
	return config.DB.Create(&products)
}

func GivenCart(cart model.Cart) {
	config.DB.Create(&cart)
}
