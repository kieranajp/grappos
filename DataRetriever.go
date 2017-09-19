package grappos

import (
	"io/ioutil"
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

func (u *url) getData() ([]byte, error) {

	res, err := http.Get(u.buildURL())
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	return body, err
}
