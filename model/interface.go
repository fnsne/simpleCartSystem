package model

type ProductRepository interface {
	List() []Product
}
type CartRepository interface {
	GetByUserID(userId int) (cart Cart)
	Update(cart Cart) error
}
