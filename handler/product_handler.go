package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllProducts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}
