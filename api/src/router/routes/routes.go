package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents all API routes
type Route struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	NeedAuth bool
}

// Config puts all routes inside the router
func Config(r *mux.Router) *mux.Router {
	routes := UserRoutes
	routes = append(routes, loginRoutes)
	routes = append(routes, PublicationsRoutes...)

	for _, route := range routes {
		if route.NeedAuth {
			r.HandleFunc(
				route.URI,
				middlewares.Logger(middlewares.Atuthenticate(route.Function)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(
				route.URI,
				middlewares.Logger(route.Function),
			).Methods(route.Method)
		}
	}

	return r
}
