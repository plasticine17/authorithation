package Network

import (
	"fmt"
	"net/http"
)

type Route struct {
	Name string

	Methot string

	Pattern string

	HandlerFunc http.HandlerFunc
}

type Routes []Route

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Products Rest API: ")
}

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"Register", "POST", "/v1/register", Register},
}
