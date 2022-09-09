package repository

import (
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"shopline-question/model"
)

type CartRepo struct {
	db *gorm.DB
}

func (r *CartRepo) NewCart(userID uint) model.Cart {
	cart := model.Cart{UserID: userID}
	r.db.Create(&cart)
	return cart
}

func (r *CartRepo) GetByID(cartID uint) (cart model.Cart, err error) {
	err = r.db.Preload("Products.Product").
		Model(&model.Cart{}).
		Where("id=?", cartID).
		Where("is_checkout=?", false).
		First(&cart).Error
	if err != nil {
		return cart, err
	}
	cart.CalculateAmount()
	return cart, nil
}

func (r *CartRepo) Update(cart model.Cart) error {
	productIds := cart.GetProductIds()
	if !PRODUCT.AllExist(productIds) {
		return errors.New("product not exist")
	}
	err, err2 := r.checkProductInventory(cart)
	if err2 != nil {
		return err2
	}
	err = r.db.Model(&model.Cart{Model: gorm.Model{ID: cart.ID}}).
		Association("Products").
		Replace(cart.Products)
	if err != nil {
		return err
	}
	amount := decimal.NewFromInt(0)
	for _, cartProduct := range cart.Products {
		productAmount := cartProduct.Product.Price.Mul(decimal.NewFromInt(int64(cartProduct.Quantity)))
		amount = amount.Add(productAmount)
	}
	err = r.db.Model(&model.Cart{}).Where("id=?", cart.ID).Update("amount", amount).Error
	return err
}

func (r *CartRepo) Checkout(cartID uint) (orderID uint, err error) {
	cart, err := r.GetByID(cartID)
	if err != nil {
		return 0, err
	}
	if !cart.CartHasOrderProduct() {
		return 0, errors.New("there should be product in cart")
	}
	tx := r.db.Begin()
	for _, product := range cart.Products {
		err := tx.Exec("UPDATE products SET inventory=inventory-? WHERE id=?",
			product.Quantity,
			product.ProductID).Error
		if err != nil {
			tx.Rollback()
			return 0, errors.New("there is some product inventory not enough")
		}
	}
	err = tx.Model(&model.Cart{}).Where("id=?", cartID).Update("is_checkout", true).Error
	if err != nil {
		tx.Rollback()
		return 0, errors.New("there is some error in checkout.Please check whether your order has been done or not.")
	}
	order := cart.ToOrder()
	err = tx.Save(order).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	return order.ID, nil
}

func (r *CartRepo) checkProductInventory(cart model.Cart) (error, error) {
	var productInfos []model.Product
	err := r.db.Model(&model.Product{}).
		Where("id in (?)", cart.GetProductIds()).
		Find(&productInfos).Error
	if err != nil {
		return nil, err
	}
	productInventoryMap := make(map[uint]uint)
	for _, info := range productInfos {
		productInventoryMap[info.ID] = info.Inventory
	}
	for _, product := range cart.Products {
		inventory, exist := productInventoryMap[product.ProductID]
		if inventory < product.Quantity || !exist {
			return nil, errors.New("inventory not enough")
		}
	}
	return err, nil
}

func (r *CartRepo) getOrderProductsBy(cartID uint) []model.CartProduct {
	var products []model.CartProduct
	r.db.Preload("Product").
		Model(&model.CartProduct{}).
		Where("cart_id=?", cartID).
		Find(&products)
	return products
}

func NewCartRepo(db *gorm.DB) *CartRepo {
	return &CartRepo{db: db}
}
