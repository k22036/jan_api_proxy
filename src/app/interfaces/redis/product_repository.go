package redis

import (
	"app/domain"
	"context"
)

type ProductRepository struct {
	RedisHandler RedisHandler
}

func (repo *ProductRepository) Set(ctx context.Context, product domain.Product) error {
	return repo.RedisHandler.Set(ctx, product.JAN, product.Name)
}

func (repo *ProductRepository) Get(ctx context.Context, jan string) (domain.Product, error) {
	name, err := repo.RedisHandler.Get(ctx, jan)
	if err != nil {
		return domain.Product{}, err
	}

	product := domain.Product{
		JAN:  jan,
		Name: name,
	}
	return product, nil
}

func (repo *ProductRepository) MGet(ctx context.Context, keys ...string) ([]string, error) {
	return repo.RedisHandler.MGet(ctx, keys...)
}

func (repo *ProductRepository) AllKeys(ctx context.Context) ([]string, error) {
	return repo.RedisHandler.AllKeys(ctx)
}

func (repo *ProductRepository) Delete(ctx context.Context, key string) error {
	return repo.RedisHandler.Delete(ctx, key)
}
