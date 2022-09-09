package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopline-question/model"
	"shopline-question/model/repository"
)

// @Summary Get Cart
// @description 獲得Cart
// @Tags cart
// @produce json
// @router /api/cart/ [get]
// @success 200 {object} model.Cart
func GetCart(ctx *gin.Context) {
	//todo:這邊先hardcode，等加上user system再從session/cookie中拿出CartID
	cart := repository.CART.GetByID(1)
	ctx.JSON(http.StatusOK, cart)
}

// @Summary Update Cart
// @description 更新Cart
// @Tags cart
// @produce json
// @router /api/cart/ [put]
// @param cart body model.Cart required "要更新的cart"
// @success 200
func UpdateCart(ctx *gin.Context) {
	var cart model.Cart
	err := ctx.ShouldBindJSON(&cart)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	//todo:這邊先hardcode，等加上user system再從session/cookie中拿出cartID
	cart.ID = 1
	for i := 0; i < len(cart.Products); i++ {
		cart.Products[i].CartID = cart.ID
	}
	err = repository.CART.Update(cart)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func CheckoutCart(ctx *gin.Context) {
	//todo:這邊先hardcode，等加上user system再從session/cookie中拿出cartID
	orderID, err := repository.CART.Checkout(1)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"orderID": orderID})
}
