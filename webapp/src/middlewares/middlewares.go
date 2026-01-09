package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Logger writes request information to the terminal
func Logger(next_func http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next_func(w, r)
	}
}

// Authenticate checks if cookies exist
func Authenticate(next_func http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, error := cookies.Read(r)
		if error != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}
		// fmt.Println(values, error)

		next_func(w, r)
	}
}
