package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

// Logger writes request information to the terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Authenticate checks if user is authenticated
func Atuthenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Atuthenticating...")
		next(w, r)
	}
}
