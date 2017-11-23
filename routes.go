package main

import (
	"net/http"

	"github.com/rendyfebry/go-restful/controllers"
)

// Route (Data Type)
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controllers.GetIndexPage,
	},
	Route{
		"Persons",
		"GET",
		"/persons",
		controllers.GetPersons,
	},
	Route{
		"Persons",
		"GET",
		"/persons/{email}",
		controllers.GetPersonByEmail,
	},
}
