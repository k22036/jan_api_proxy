package infrastructure

import (
	gin "github.com/gin-gonic/gin"

	"app/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	LoadEnv()
	productController := controllers.NewProductController(NewRedisHandler())
	router.POST("/products", func(c *gin.Context) { productController.Add(c) })
	router.GET("/products/:jan", func(c *gin.Context) { productController.Get(c) })
	Router = router
}
