package yahooshopping

import "net/http"

type YahooShoppingHandler interface {
	GetProduct(jan string) (*http.Response, error)
}
