package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	var port = flag.String("p", "3000", "Server port")
	flag.Parse()

	fmt.Println("Server listening on:", *port)
	log.Fatal(http.ListenAndServe(":"+*port, router))
}
