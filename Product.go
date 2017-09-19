package grappos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var productBaseURL = "http://www.grappos.com/api2/subscriber.php?1=1&format=json"

type product struct {
	ProductID   string `json:"ProductID"`
	Description string `json:"description"`
	BrandID     string `json:"BrandID"`
	IsToppick   int    `json:"is_toppick"`
}

// ProductAPIResponse Holds the API response.
type ProductAPIResponse struct {
	Products []product `json:"products"`
}

// GetProducts Returns a list of Products.
func GetProducts(u string) (*ProductAPIResponse, error) {
	var s = new(ProductAPIResponse)
	queryParams := fmt.Sprintf("&uid=%s", u)

	err := productDataRetriever(s, productBaseURL+queryParams)

	if len(s.Products) == 0 {
		err = fmt.Errorf("Product not found with id: %s", u)
	}

	return s, err
}
