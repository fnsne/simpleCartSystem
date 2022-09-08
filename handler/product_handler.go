package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopline-question/model/repository"
)

// @Summary Get Product list
// @description 獲得products list
// @Tags product
// @produce json
// @router /api/product/ [get]
// @success 200 {object} []model.Product
func GetAllProducts(ctx *gin.Context) {
	products := repository.PRODUCT.List()
	ctx.JSON(http.StatusOK, products)
}
