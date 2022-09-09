package tests

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"shopline-question/config"
	"shopline-question/model"
	"shopline-question/model/repository"
	"shopline-question/router"
	"testing"
)

func TestCartTestsSuite(t *testing.T) {
	suite.Run(t, new(CartTests))
}

type CartTests struct {
	suite.Suite
	engine   *gin.Engine
	request  *http.Request
	response *http.Response
}

func (suite *CartTests) SetupTest() {
	db := config.NewUserDB()
	config.DB = db
	config.DropTestTable(&model.Product{})
	config.DropTestTable(&model.Cart{})
	config.DropTestTable(&model.User{})
	config.DropTestTable(&model.CartProduct{})
	config.DropTestTable(&model.Order{})
	config.DropTestTable(&model.OrderProduct{})
	model.Migrate(db)
	repository.Initial(db)
	suite.engine = router.NewRouter()
}

func (suite *CartTests) Test_GetCart() {
	GivenProducts([]model.Product{
		{
			Model:     gorm.Model{ID: 1},
			Name:      "product 1",
			Price:     decimal.NewFromInt(10),
			Inventory: 3,
		},
		{
			Model:     gorm.Model{ID: 2},
			Name:      "product 2",
			Price:     decimal.NewFromInt(20),
			Inventory: 5,
		},
	})
	GivenCart(model.Cart{
		UserID: 1,
		Products: []model.CartProduct{
			{ProductID: 1, Quantity: 1},
			{ProductID: 2, Quantity: 3},
		},
		Amount: decimal.NewFromInt(70)})
	suite.currentCartShouldBe(model.Cart{
		UserID: 1,
		Products: []model.CartProduct{
			{ProductID: 1, Quantity: 1},
			{ProductID: 2, Quantity: 3},
		},
		Amount: decimal.NewFromInt(70),
	})

}

//todo get cart without existed cart will create one

func (suite *CartTests) Test_AddProductToCart() {
	GivenProducts([]model.Product{
		{
			Model:     gorm.Model{ID: 1},
			Name:      "product 1",
			Price:     decimal.NewFromInt(10),
			Inventory: 3,
		},
		{
			Model:     gorm.Model{ID: 2},
			Name:      "product 2",
			Price:     decimal.NewFromInt(20),
			Inventory: 5,
		},
	})
	GivenCart(model.Cart{
		UserID: 1,
		Products: []model.CartProduct{
			{ProductID: 1, Quantity: 1},
			{ProductID: 2, Quantity: 3},
		},
		Amount: decimal.NewFromInt(70)})

	addProduct := model.Cart{
		UserID: 1,
		Products: []model.CartProduct{
			{ProductID: 1, Quantity: 3},
			{ProductID: 2, Quantity: 1},
		},
	}
	suite.givenUpdateCartReq(addProduct)
	suite.responseStatusShouldBe(http.StatusOK)
	suite.currentCartShouldBe(model.Cart{
		UserID: 1,
		Products: []model.CartProduct{
			{ProductID: 1, Quantity: 3},
			{ProductID: 2, Quantity: 1},
		},
		Amount: decimal.NewFromInt(50),
	})
}
func (suite *CartTests) Test_AddProductToCart_withNotExistProduct() {
	GivenProducts([]model.Product{
		{
			Model:     gorm.Model{ID: 1},
			Name:      "product 1",
			Price:     decimal.NewFromInt(10),
			Inventory: 3,
		},
		{
			Model:     gorm.Model{ID: 2},
			Name:      "product 2",
			Price:     decimal.NewFromInt(20),
			Inventory: 5,
		},
	})
	GivenCart(model.Cart{
		UserID: 1,
		Products: []model.CartProduct{
			{ProductID: 1, Quantity: 1},
			{ProductID: 2, Quantity: 3},
		},
		Amount: decimal.NewFromInt(70)})

	addProduct := model.Cart{
		UserID: 1,
		Products: []model.CartProduct{
			{ProductID: 100, Quantity: 3},
		},
	}
	suite.givenUpdateCartReq(addProduct)
	suite.responseStatusShouldBe(http.StatusBadRequest)
	suite.currentCartShouldBe(model.Cart{
		UserID: 1,
		Products: []model.CartProduct{
			{ProductID: 1, Quantity: 1},
			{ProductID: 2, Quantity: 3},
		},
		Amount: decimal.NewFromInt(70),
	})
}
func (suite *CartTests) Test_AddProductToCart_withNotEnoughProduct() {
	GivenProducts([]model.Product{
		{
			Model:     gorm.Model{ID: 1},
			Name:      "product 1",
			Price:     decimal.NewFromInt(10),
			Inventory: 1,
		},
	})
	GivenCart(model.Cart{UserID: 1})
	addProduct := model.Cart{
		UserID: 1,
		Products: []model.CartProduct{
			{ProductID: 1, Quantity: 3},
		},
	}
	suite.givenUpdateCartReq(addProduct)
	suite.responseStatusShouldBe(http.StatusBadRequest)
	suite.currentCartShouldBe(model.Cart{
		UserID:   1,
		Products: []model.CartProduct{},
		Amount:   decimal.NewFromInt(0),
	})
}
func (suite *CartTests) Test_checkoutCart_when_allInventoryEnough() {
	GivenProducts([]model.Product{
		{
			Model:     gorm.Model{ID: 1},
			Name:      "product 1",
			Price:     decimal.NewFromInt(10),
			Inventory: 1,
		},
		{
			Model:     gorm.Model{ID: 2},
			Name:      "product 2",
			Price:     decimal.NewFromInt(20),
			Inventory: 2,
		},
	})
	GivenCart(model.Cart{UserID: 1,
		Products: []model.CartProduct{
			{ProductID: 1, Quantity: 1},
			{ProductID: 2, Quantity: 1},
		},
	})
	suite.givenCartCheckoutReq()
	suite.responseStatusShouldBe(http.StatusOK)
	suite.currentCartShouldBe(model.Cart{
		UserID:   0,
		Products: []model.CartProduct{},
		Amount:   decimal.NewFromInt(0),
	})
	suite.orderInDBShouldBe(1, model.Order{
		UserID: 1,
		Products: []model.OrderProduct{
			{OrderID: 1, ProductID: 1, Quantity: 1},
			{OrderID: 1, ProductID: 2, Quantity: 1},
		},
		Amount: decimal.NewFromInt(30),
	})
	suite.productInventoryShouldBe(1, 0)
	suite.productInventoryShouldBe(2, 1)
}

func (suite *CartTests) Test_checkoutCart_when_inventoryNotEnough() {
	GivenProducts([]model.Product{
		{
			Model:     gorm.Model{ID: 1},
			Name:      "product 1",
			Price:     decimal.NewFromInt(10),
			Inventory: 1,
		},
		{
			Model:     gorm.Model{ID: 2},
			Name:      "product 2",
			Price:     decimal.NewFromInt(20),
			Inventory: 2,
		},
	})
	GivenCart(model.Cart{UserID: 1,
		Products: []model.CartProduct{
			{ProductID: 1, Quantity: 1},
			{ProductID: 2, Quantity: 3},
		},
	})
	suite.givenCartCheckoutReq()
	suite.responseStatusShouldBe(http.StatusBadRequest)
	suite.currentCartShouldBe(model.Cart{
		UserID: 1,
		Products: []model.CartProduct{
			{ProductID: 1, Quantity: 1},
			{ProductID: 2, Quantity: 3},
		},
		Amount: decimal.NewFromInt(70),
	})
	suite.orderShouldNotExist(1)
	suite.productInventoryShouldBe(1, 1)
	suite.productInventoryShouldBe(2, 2)
}
func (suite *CartTests) Test_checkoutCart_when_noProductInCart() {
	GivenProducts([]model.Product{
		{
			Model:     gorm.Model{ID: 1},
			Name:      "product 1",
			Price:     decimal.NewFromInt(10),
			Inventory: 1,
		},
		{
			Model:     gorm.Model{ID: 2},
			Name:      "product 2",
			Price:     decimal.NewFromInt(20),
			Inventory: 2,
		},
	})
	GivenCart(model.Cart{UserID: 1, Products: []model.CartProduct{}})
	suite.givenCartCheckoutReq()
	suite.responseStatusShouldBe(http.StatusBadRequest)
	suite.currentCartShouldBe(model.Cart{
		UserID:   1,
		Products: []model.CartProduct{},
		Amount:   decimal.NewFromInt(0),
	})
	suite.orderShouldNotExist(1)
	suite.productInventoryShouldBe(1, 1)
	suite.productInventoryShouldBe(2, 2)
}

func (suite *CartTests) orderShouldNotExist(cartID int) {
	var count int64
	err := config.DB.Model(&model.Order{}).Where("id=?", cartID).Count(&count).Error
	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), int64(0), count)
}

// todo: when no product in cart
func (suite *CartTests) productInventoryShouldBe(productID int, inventory uint) {
	var p model.Product
	err := config.DB.Model(&model.Product{}).Where("id=?", productID).First(&p).Error
	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), inventory, p.Inventory)
}

func (suite *CartTests) orderInDBShouldBe(orderID int, expectedOrder model.Order) {
	var order model.Order
	err := config.DB.Model(&model.Order{}).Preload("Products").Where("id=?", orderID).First(&order).Error
	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedOrder.UserID, order.UserID)
	for i := 0; i < len(expectedOrder.Products); i++ {
		actual := order.Products[i]
		expected := expectedOrder.Products[i]
		assert.Equal(suite.T(), expected.ProductID, actual.ProductID)
		assert.Equal(suite.T(), expected.Quantity, actual.Quantity)
	}
	assert.Equal(suite.T(), expectedOrder.Amount, order.Amount)
}

func (suite *CartTests) givenCartCheckoutReq() {
	suite.request = httptest.NewRequest(http.MethodPost, "/api/cart/checkout", nil)
}

func (suite *CartTests) givenUpdateCartReq(cart model.Cart) {
	b := &bytes.Buffer{}
	err := json.NewEncoder(b).Encode(&cart)
	require.NoError(suite.T(), err)
	suite.request = httptest.NewRequest(http.MethodPut, "/api/cart/", b)
}

func (suite *CartTests) responseCartShouldBe(expectedCart model.Cart) {
	var responseCart model.Cart
	err := json.NewDecoder(suite.response.Body).Decode(&responseCart)
	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedCart.UserID, responseCart.UserID)
	assert.Equal(suite.T(), len(expectedCart.Products), len(responseCart.Products))
	if len(expectedCart.Products) != 0 {
		for i := 0; i < len(expectedCart.Products); i++ {
			assert.Equal(suite.T(), expectedCart.Products[i].ProductID, responseCart.Products[i].ProductID)
		}
		assert.Equal(suite.T(), expectedCart.Amount, responseCart.Amount)
	}
}

func (suite *CartTests) responseStatusShouldBe(status int) {
	w := httptest.NewRecorder()
	suite.engine.ServeHTTP(w, suite.request)
	suite.response = w.Result()
	assert.Equal(suite.T(), status, suite.response.StatusCode)
}

func (suite *CartTests) giveGetCartReq() {
	suite.request = httptest.NewRequest(http.MethodGet, "/api/cart/", nil)
}
func (suite *CartTests) currentCartShouldBe(cart model.Cart) {
	suite.giveGetCartReq()
	suite.responseStatusShouldBe(http.StatusOK)
	suite.responseCartShouldBe(cart)
}
