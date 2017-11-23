package controllers

import (
	"fmt"
	"net/http"
)

// GetIndexPage get index page
func GetIndexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome")
}
