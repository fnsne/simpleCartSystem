package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/goccy/go-json"
	"github.com/shopspring/decimal"
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

func TestProductTestsSuite(t *testing.T) {
	suite.Run(t, new(ProductTests))
}

type ProductTests struct {
	suite.Suite
	engine   *gin.Engine
	request  *http.Request
	response *http.Response
}

func (suite *ProductTests) SetupTest() {
	db := config.NewUserDB()
	config.DB = db
	//config.DropTestTable()
	model.Migrate(db)
	repository.Initial(db)
	suite.engine = router.NewRouter()
}

func (suite *ProductTests) Test_GetProductionList() {
	suite.givenProducts([]model.Product{
		{
			Model:     gorm.Model{ID: 1},
			Name:      "product1",
			Price:     decimal.NewFromInt(10),
			Inventory: 2,
		},
		{
			Model:     gorm.Model{ID: 2},
			Name:      "product2",
			Price:     decimal.NewFromInt(20),
			Inventory: 5,
		},
	})

	suite.givenGetProductListReq()
	suite.responseStatusShouldBe(http.StatusOK)
	suite.responseProductListShouldBe([]model.Product{
		{Name: "product1", Price: decimal.NewFromInt(10), Inventory: 2},
		{Name: "product2", Price: decimal.NewFromInt(20), Inventory: 5},
	})
}

func (suite *ProductTests) givenProducts(products []model.Product) *gorm.DB {
	return config.DB.Create(&products)
}

func (suite *ProductTests) responseProductListShouldBe(expectedProducts []model.Product) {
	var products []model.Product
	err := json.NewDecoder(suite.response.Body).Decode(&products)
	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), len(expectedProducts), len(products))
	for i := 0; i < len(expectedProducts); i++ {
		expected := expectedProducts[i]
		product := products[i]
		assert.Equal(suite.T(), expected.Name, product.Name)
		assert.Equal(suite.T(), expected.Price, product.Price)
		assert.Equal(suite.T(), expected.Inventory, product.Inventory)
	}
}

func (suite *ProductTests) responseStatusShouldBe(status int) {
	w := httptest.NewRecorder()
	suite.engine.ServeHTTP(w, suite.request)
	suite.response = w.Result()
	assert.Equal(suite.T(), status, suite.response.StatusCode)
}

func (suite *ProductTests) givenGetProductListReq() {
	suite.request = httptest.NewRequest(http.MethodGet, "/api/product/", nil)
}
