package main

import (
	"net/http"

	"github.com/rendyfebry/go-restful/controllers"
)

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
		controllers.HandlerIndex,
	},
	Route{
		"Persons",
		"GET",
		"/persons",
		controllers.HandlerPersons,
	},
	Route{
		"Persons",
		"GET",
		"/persons/{email}",
		controllers.HandlerPersonsSingle,
	},
}
