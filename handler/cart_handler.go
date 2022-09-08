package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopline-question/model/repository"
)

func GetCart(ctx *gin.Context) {
	//todo:這邊先hardcode，等加上user system再從session/cookie中拿出userID
	cart := repository.CART.GetByUserID(1)
	ctx.JSON(http.StatusOK, cart)
}
