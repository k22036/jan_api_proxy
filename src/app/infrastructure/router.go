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

	router.POST("/products", func(c *gin.Context) { productController.Add(c) })
	router.GET("/products/:jan", func(c *gin.Context) { productController.Get(c) })
	Router = router
}
