package controllers

import (
	"fmt"
	"net/http"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome")
}
