package infrastructure

import (
	"github.com/gin-gonic/gin"

	"app/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	LoadEnv()

	redisHandler := NewRedisHandler()
	geminiHandler := NewGeminiHandler()
	yahooShoppingHandler := NewYahooShoppingHandler()
	productController := controllers.NewProductController(redisHandler, geminiHandler, yahooShoppingHandler)

	api_v1 := router.Group("/api/v1")
	api_v1.POST("/product", func(c *gin.Context) { productController.Add(c) })
	api_v1.GET("/product/:jan", func(c *gin.Context) { productController.Get(c) })
	api_v1.GET("/products", func(c *gin.Context) { productController.GetAll(c) })
	api_v1.DELETE("/product/:jan", func(c *gin.Context) { productController.Delete(c) })

	Router = router
}
