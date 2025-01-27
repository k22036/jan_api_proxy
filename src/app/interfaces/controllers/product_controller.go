package controllers

import (
	"app/domain"
	"app/interfaces/redis"
	"app/usecase"
	"context"
)

type ProductController struct {
	ProductInteractor usecase.ProductInteractor
}

func NewProductController(redisHandler redis.RedisHandler) *ProductController {
	return &ProductController{
		ProductInteractor: usecase.ProductInteractor{
			ProductRepository: &redis.ProductRepository{
				RedisHandler: redisHandler,
			},
		},
	}
}

func (controller *ProductController) Add(c Context) {
	product := domain.Product{}
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(400, NewError(err))
		return
	}

	ctx := context.Background()
	err1 := controller.ProductInteractor.Add(ctx, product)
	if err1 != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, nil)
}

func (controller *ProductController) Get(c Context) {
	jan := c.Param("jan")
	ctx := context.Background()
	product, err := controller.ProductInteractor.Get(ctx, jan)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, product)
}
