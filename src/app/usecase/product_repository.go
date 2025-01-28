package usecase

import (
	"app/domain"
	"context"
)

type ProductRepository interface {
	Set(ctx context.Context, product domain.Product) error
	Get(ctx context.Context, jan string) (domain.Product, error)
	MGet(ctx context.Context, keys ...string) ([]string, error)
	AllKeys(ctx context.Context) ([]string, error)
}
