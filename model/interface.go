package model

type ProductRepository interface {
	List() []Product
}
