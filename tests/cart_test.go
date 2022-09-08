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
		Products: []model.OrderProduct{
			{ProductID: 1, Quantity: 1},
			{ProductID: 2, Quantity: 3},
		},
		Amount: decimal.NewFromInt(70)})
	suite.currentCartShouldBe(model.Cart{
		UserID: 1,
		Products: []model.OrderProduct{
			{ProductID: 1, Quantity: 1},
			{ProductID: 2, Quantity: 3},
		},
		Amount: decimal.NewFromInt(70),
	})

}

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
		Products: []model.OrderProduct{
			{ProductID: 1, Quantity: 1},
			{ProductID: 2, Quantity: 3},
		},
		Amount: decimal.NewFromInt(70)})

	addProduct := model.Cart{
		UserID: 1,
		Products: []model.OrderProduct{
			{ProductID: 1, Quantity: 3},
			{ProductID: 2, Quantity: 1},
		},
	}
	suite.givenUpdateCartReq(addProduct)
	suite.responseStatusShouldBe(http.StatusOK)
	suite.currentCartShouldBe(model.Cart{
		UserID: 1,
		Products: []model.OrderProduct{
			{ProductID: 1, Quantity: 3},
			{ProductID: 2, Quantity: 1},
		},
		Amount: decimal.NewFromInt(50),
	})
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
	for i := 0; i < len(expectedCart.Products); i++ {
		assert.Equal(suite.T(), expectedCart.Products[i].ProductID, responseCart.Products[i].ProductID)
	}
	assert.Equal(suite.T(), expectedCart.Amount, responseCart.Amount)
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
