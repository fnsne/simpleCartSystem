package repository

import "gorm.io/gorm"

var CART *CartRepo
var PRODUCT *ProductRepo

func Initial(db *gorm.DB) {
	//CART=NewCartRepo(db)
	PRODUCT = NewProductRepo(db)
}
