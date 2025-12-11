package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// Before the first run we need to get:
// go get github.com/gorilla/mux
func main() {
	config.Load()

	fmt.Println("Online!")
	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
