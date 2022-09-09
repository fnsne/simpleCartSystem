package model

type ProductRepository interface {
	List() []Product
	AllExist(productIds []uint) (Exist bool)
}
type CartRepository interface {
	GetByID(cartID int) (cart Cart)
	Update(cart Cart) error
	Checkout(cartID int) (orderID uint, err error)
}
