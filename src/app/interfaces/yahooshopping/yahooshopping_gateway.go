package yahooshopping

import "net/http"

type YahooShoppingGateway struct {
	YahooShoppingHandler YahooShoppingHandler
}

func (gateway *YahooShoppingGateway) GetProduct(jan string) (*http.Response, error) {
	return gateway.YahooShoppingHandler.GetProduct(jan)
}
