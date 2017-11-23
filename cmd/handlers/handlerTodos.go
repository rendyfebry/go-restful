package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rendyfebry/go-restful/cmd/utils"
)

func HandlerTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(utils.TodoList); err != nil {
		panic(err)
	}
}
