package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// Before the first run we need to get:
// go get github.com/gorilla/mux
func main() {
	fmt.Println("Online!")
	r := router.Generate()

	log.Fatal(http.ListenAndServe(":5000", r))
}
