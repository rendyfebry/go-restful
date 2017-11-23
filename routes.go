package main

import (
	"net/http"
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
		HandlerIndex,
	},
	Route{
		"Todos",
		"GET",
		"/todos",
		HandlerTodos,
	},
	Route{
		"Todos Create",
		"POST",
		"/todos",
		HandlerTodosCreate,
	},
	Route{
		"Todos Single",
		"GET",
		"/todos/{todoId}",
		HandlerTodosSingle,
	},
	Route{
		"Persons",
		"GET",
		"/persons",
		HandlerPersons,
	},
}
