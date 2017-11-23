package main

import (
	"net/http"

	"github.com/rendyfebry/go-restful/cmd/handlers"
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
		handlers.HandlerIndex,
	},
	Route{
		"Persons",
		"GET",
		"/persons",
		handlers.HandlerPersons,
	},
	Route{
		"Persons",
		"GET",
		"/persons/{email}",
		handlers.HandlerPersonsSingle,
	},
}
