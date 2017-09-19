package grappos

import (
	"errors"
	"fmt"
	"log"
)

var baseURL = "http://www.grappos.com/api2/locate.php?1=1&format=json"

type location struct {
	DisplayName string `json:"displayName"`
	Latitude    string `json:"lat"`
	Longitude   string `json:"lon"`
	ZipCode     string `json:"zip"`
}

// LocationAPIResponse Holds the API response.
type LocationAPIResponse struct {
	Locations []location `json:"locations"`
}

// GetLocations Returns all locations.
func GetLocations(n int) (*LocationAPIResponse, error) {

	var s = new(LocationAPIResponse)

	queryParams := ""

	if n >= 0 {
		queryParams = fmt.Sprintf("&limit=%d", n)
	} else {
		return s, errors.New("Limit should be a positive int")
	}

	err := locationDataRetriever(s, baseURL+queryParams)
	if err != nil {
		log.Fatalf("Searching for location went wrong: %s", err)
	}

	return s, err
}

// SearchForLocation Postal Code or City Name (ex: “13066”, “San Francisco”).
func SearchForLocation(l string) (*LocationAPIResponse, error) {
	queryParams := ""

	var s = new(LocationAPIResponse)

	if len(l) == 5 {
		queryParams = fmt.Sprintf("&locate=%s", l)
	} else {
		return s, errors.New("Invalid Postal Code")
	}

	err := locationDataRetriever(s, baseURL+queryParams)
	if err != nil {
		log.Fatalf("Searching for location went wrong: %s", err)
	}

	return s, err
}
