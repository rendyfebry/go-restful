package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome")
}

func HandlerTodos(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write Presentation"},
		Todo{Name: "Host meetup"},
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func HandlerTodosSingle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprint(w, "Todo show:", todoId)
}