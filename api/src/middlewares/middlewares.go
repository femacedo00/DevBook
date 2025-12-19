package middlewares

import (
	"api/src/authentication"
	"api/src/response"
	"log"
	"net/http"
)

// Logger writes request information to the terminal
func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}

// Authenticate checks if user is authenticated
func Atuthenticate(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if error := authentication.ValidateToken(r); error != nil {
			response.Error(w, http.StatusUnauthorized, error)
			return
		}
		nextFunc(w, r)
	}
}
