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
		"Get Persons",
		"GET",
		"/persons",
		controllers.GetPersons,
	},
	Route{
		"Create Person",
		"POST",
		"/persons",
		controllers.CreatePerson,
	},
	Route{
		"Get Single Person",
		"GET",
		"/persons/{email}",
		controllers.GetPersonByEmail,
	},
}
