package routes

import "net/http"

// Route represents all API routes
type Route struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	NeedAuth bool
}
