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
	r := router.Generate()

	fmt.Printf("Online on port: %d!\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
