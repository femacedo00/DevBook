package router

import "github.com/gorilla/mux"

// Generate returns a router with all configured routes
func Generate() *mux.Router {
	return mux.NewRouter()
}
