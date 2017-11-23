package utils

import (
	"encoding/json"
	"net/http"

	"github.com/rendyfebry/go-restful/models"
)

func SendJsonResponse(w http.ResponseWriter, error interface{}, message interface{}, data interface{}) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	response := &models.ResponseObj{
		Error:   error,
		Data:    data,
		Message: message,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
