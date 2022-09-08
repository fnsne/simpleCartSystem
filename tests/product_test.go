package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
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
	engine *gin.Engine
}

func (suite *ProductTests) SetupTest() {
	db := config.NewUserDB()
	config.DB = db
	config.DropTestTable(&model.User{})
	model.Migrate(db)
	repository.Initial(db)
	suite.engine = router.NewRouter()
}
