package model

type ProductRepository interface {
	List() []Product
	AllExist(productIds []uint) (Exist bool)
}
type CartRepository interface {
	GetByUserID(userId int) (cart Cart)
	Update(cart Cart) error
}
