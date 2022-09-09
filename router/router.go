package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"shopline-question/handler"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router = setupRouter(router)
	return router
}
func setupRouter(engine *gin.Engine) *gin.Engine {
	engine = setupSwagger(engine)

	api := engine.Group("/api")
	productAPI := api.Group("/product")
	productAPI.GET("/", handler.GetAllProducts)

	cartAPI := api.Group("/cart")
	cartAPI.GET("/", handler.GetCart)
	cartAPI.PUT("/", handler.UpdateCart)
	cartAPI.POST("/checkout", handler.CheckoutCart)

	return engine
}
func setupSwagger(engine *gin.Engine) *gin.Engine {
	if mode := gin.Mode(); mode == gin.DebugMode {
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	return engine
}
