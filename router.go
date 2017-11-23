package main

import (
	"log"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func NewRouter() *mux.Router {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiPrefix := os.Getenv("API_PREFIX")
	apiVersion := os.Getenv("API_VERSION")

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			PathPrefix(apiPrefix + apiVersion).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
