package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopline-question/model/repository"
)

func GetAllProducts(ctx *gin.Context) {
	products := repository.PRODUCT.List()
	ctx.JSON(http.StatusOK, products)
}
