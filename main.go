package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	var port = flag.String("p", "3000", "Server port")
	flag.Parse()

	fmt.Println("Server listening on:", *port)

	r := NewRouter()
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":"+*port, loggedRouter)
}
