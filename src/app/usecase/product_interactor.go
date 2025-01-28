package usecase

import (
	"app/domain"
	"app/interfaces/gemini"
	"app/interfaces/yahooshopping"
	geminiParser "app/usecase/gemini"
	yahooshoppingParser "app/usecase/yahooshopping"
	"context"
	"fmt"
)

type productInteractor interface {
	Add(ctx context.Context, product domain.Product) error
	Get(ctx context.Context, jan string) (domain.Product, error)
}

type ProductInteractor struct {
	ProductRepository    ProductRepository
	GeminiGateway        gemini.GeminiGateway
	YahooShoppingGateway yahooshopping.YahooShoppingGateway
}

func (interactor *ProductInteractor) Add(ctx context.Context, product domain.Product) error {
	return interactor.ProductRepository.Set(ctx, product)
}

func (interactor *ProductInteractor) Get(ctx context.Context, jan string) (domain.Product, error) {
	fmt.Println("Get product from Redis")
	product, _ := interactor.ProductRepository.Get(ctx, jan)
	if product.Name != "" {
		return product, nil
	}

	fmt.Println("Get product from Yahoo Shopping API")
	response, err := interactor.YahooShoppingGateway.GetProduct(jan)
	if err != nil {
		return domain.Product{}, err
	}
	names, err1 := yahooshoppingParser.ParseResponse(response)
	if err1 != nil {
		return domain.Product{}, err1
	}
	if len(names) == 0 {
		return domain.Product{}, fmt.Errorf("failed to get product")
	}

	fmt.Println("Extract product from Gemini API")
	prompt := geminiParser.ParseInput(names)
	response1, err2 := interactor.GeminiGateway.Request(ctx, prompt)
	if err2 != nil {
		return domain.Product{}, err2
	}
	name := geminiParser.ParseResponse(response1)

	if name == "" {
		return domain.Product{}, fmt.Errorf("failed to get product")
	}

	err3 := interactor.ProductRepository.Set(ctx, domain.Product{JAN: jan, Name: name})
	if err3 != nil {
		fmt.Println("Failed to set product to Redis")
	}

	return domain.Product{JAN: jan, Name: name}, nil
}
