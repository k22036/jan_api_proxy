package infrastructure

import (
	"fmt"
	"net/http"
	"os"
)

type YahooShoppingHandler struct {
	client *http.Client
}

func NewYahooShoppingHandler() *YahooShoppingHandler {
	return &YahooShoppingHandler{
		client: &http.Client{},
	}
}

func (h *YahooShoppingHandler) GetProduct(jan string) (*http.Response, error) {
	appId := os.Getenv("YAHOO_APP_ID")
	url := "https://shopping.yahooapis.jp/ShoppingWebService/V3/itemSearch?appid=" + appId + "&jan_code=" + jan
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	//res, err := h.client.Do(req)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get product: %s", res.Status)
	}

	return res, nil
}
