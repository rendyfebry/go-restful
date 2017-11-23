package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HandlerIndex)
	router.HandleFunc("/todos", HandlerTodos)
	router.HandleFunc("/todos/{todoId}", HandlerTodosSingle)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome")
}

func HandlerTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Todos")
}

func HandlerTodosSingle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprint(w, "Todo show:", todoId)
}
