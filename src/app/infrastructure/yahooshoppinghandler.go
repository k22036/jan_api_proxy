package infrastructure

import (
	"encoding/json"
	"fmt"
	"io"
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

type Response struct {
	TotalResultsAvailable int     `json:"totalResultsAvailable"`
	TotalResultsReturned  int     `json:"totalResultsReturned"`
	FirstResultPosition   int     `json:"firstResultPosition"`
	Request               Request `json:"request"`
	Hits                  []Hit   `json:"hits"`
}

type Request struct {
	Query string `json:"query"`
}

type Hit struct {
	Index                 int             `json:"index"`
	Name                  string          `json:"name"`
	Description           string          `json:"description"`
	HeadLine              string          `json:"headLine"`
	URL                   string          `json:"url"`
	InStock               bool            `json:"inStock"`
	Code                  string          `json:"code"`
	Condition             string          `json:"condition"`
	ImageID               string          `json:"imageId"`
	Image                 Image           `json:"image"`
	Review                Review          `json:"review"`
	AffiliateRate         float64         `json:"affiliateRate"`
	Price                 int             `json:"price"`
	PremiumPrice          int             `json:"premiumPrice"`
	PremiumPriceStatus    bool            `json:"premiumPriceStatus"`
	PremiumDiscountType   interface{}     `json:"premiumDiscountType"`
	PremiumDiscountRate   interface{}     `json:"premiumDiscountRate"`
	PriceLabel            PriceLabel      `json:"priceLabel"`
	Point                 Point           `json:"point"`
	Shipping              Shipping        `json:"shipping"`
	GenreCategory         GenreCategory   `json:"genreCategory"`
	ParentGenreCategories []GenreCategory `json:"parentGenreCategories"`
	Brand                 Brand           `json:"brand"`
	ParentBrands          []Brand         `json:"parentBrands"`
	JanCode               string          `json:"janCode"`
	ReleaseDate           interface{}     `json:"releaseDate"`
	Seller                Seller          `json:"seller"`
	Delivery              Delivery        `json:"delivery"`
	Payment               string          `json:"payment"`
}

type Image struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
}

type Review struct {
	Rate  float64 `json:"rate"`
	Count int     `json:"count"`
	URL   string  `json:"url"`
}

type PriceLabel struct {
	Taxable         interface{} `json:"taxable"`
	DefaultPrice    int         `json:"defaultPrice"`
	DiscountedPrice interface{} `json:"discountedPrice"`
	FixedPrice      interface{} `json:"fixedPrice"`
	PremiumPrice    interface{} `json:"premiumPrice"`
	PeriodStart     interface{} `json:"periodStart"`
	PeriodEnd       interface{} `json:"periodEnd"`
}

type Point struct {
	Amount        int `json:"amount"`
	Times         int `json:"times"`
	PremiumAmount int `json:"premiumAmount"`
	PremiumTimes  int `json:"premiumTimes"`
}

type Shipping struct {
	Code int    `json:"code"`
	Name string `json:"name"`
}

type GenreCategory struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Depth int    `json:"depth"`
}

type Brand struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Seller struct {
	SellerID     string `json:"sellerId"`
	Name         string `json:"name"`
	URL          string `json:"url"`
	IsBestSeller bool   `json:"isBestSeller"`
	Review       Review `json:"review"`
	ImageID      string `json:"imageId"`
}

type Delivery struct {
	Area     string      `json:"area"`
	DeadLine interface{} `json:"deadLine"`
	Day      interface{} `json:"day"`
}

func (h *YahooShoppingHandler) GetProduct(jan string) (*Response, error) {
	appId := os.Getenv("YAHOO_APP_ID")
	url := "https://shopping.yahooapis.jp/ShoppingWebService/V3/itemSearch?appid=" + appId + "&jan_code=" + jan
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)
	if res.StatusCode != http.StatusOK {
		return nil, err
	}

	// レスポンスボディの文字列(json)を取得
	buffer := make([]byte, 1024)
	var jsonString string
	for {
		c, err := res.Body.Read(buffer)
		if err != nil {
			if err == io.EOF {
				jsonString += string(buffer[:c])
				break
			}
			print(err)
		}
		jsonString += string(buffer[:c])
	}

	// json文字列を構造体に変換
	var response Response
	err1 := json.Unmarshal([]byte(jsonString), &response)
	if err1 != nil {
		fmt.Println("Error decoding JSON:", err1)
		return nil, err1
	}

	//fmt.Printf("Parsed Response: %+v\n", response)

	return &response, nil
}
