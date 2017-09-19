package grappos

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type url struct {
	Route       string
	QueryParams string
}

var newBaseURL = "http://www.grappos.com/api2/"

func (u *url) setModel(m string) {
	u.Route = m
}

func (u *url) buildURL() string {

	switch u.Route {
	case "locate":
		u.Route = "locate.php?1=1&format=json"
	case "subscriber":
		u.Route = "subscriber.php?1=1&format=json"
	case "retailer":
		u.Route = "retailer.php?1=1&format=json"
	}

	return newBaseURL + u.Route + u.QueryParams
}

func (u *url) addQueryParam(q string) {
	u.QueryParams += q
}

func (u *url) getData() error {

	res, err := http.Get(u.buildURL())
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal([]byte(body), &m)
	if err != nil {
		log.Println("whoops:", err)
	}

	return err
}

func locationDataRetriever(m *LocationAPIResponse, q string) error {
	res, err := http.Get(q)

	body, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal([]byte(body), &m)

	return err
}

func productDataRetriever(m *ProductAPIResponse, q string) error {
	res, err := http.Get(q)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal([]byte(body), &m)
	if err != nil {
		log.Println("whoops:", err)
	}

	return err
}

func retailerDataRetriever(m *RetailerAPIResponse, q string) error {
	res, err := http.Get(q)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal([]byte(body), &m)
	if err != nil {
		log.Println("whoops:", err)
	}

	return err
}
