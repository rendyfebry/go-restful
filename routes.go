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
		"Todos",
		"GET",
		"/todos",
		handlers.HandlerTodos,
	},
	Route{
		"Todos Create",
		"POST",
		"/todos",
		handlers.HandlerTodosCreate,
	},
	Route{
		"Todos Single",
		"GET",
		"/todos/{todoId}",
		handlers.HandlerTodosSingle,
	},
	Route{
		"Persons",
		"GET",
		"/persons",
		handlers.HandlerPersons,
	},
}
