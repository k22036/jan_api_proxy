package usecase

import (
	"app/domain"
	"context"
)

type ProductInteractor struct {
	ProductRepository ProductRepository
}

func (interactor *ProductInteractor) Add(ctx context.Context, product domain.Product) error {
	return interactor.ProductRepository.Set(ctx, product)
}

func (interactor *ProductInteractor) Get(ctx context.Context, jan string) (domain.Product, error) {
	return interactor.ProductRepository.Get(ctx, jan)
}
