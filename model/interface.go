package model

type ProductRepository interface {
	List() []Product
	AllExist(productIds []uint) (Exist bool)
}
type CartRepository interface {
	GetByID(cartID uint) (cart Cart, err error)
	Update(cart Cart) error
	Checkout(cartID uint) (orderID uint, err error)
	NewCart(userID uint) Cart
}
