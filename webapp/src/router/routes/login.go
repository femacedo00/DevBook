package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routesLogin = []Route{
	{
		URI:      "/",
		Method:   http.MethodGet,
		Function: controllers.LoadLoginPage,
		NeedAuth: false,
	},
	{
		URI:      "/login",
		Method:   http.MethodGet,
		Function: controllers.LoadLoginPage,
		NeedAuth: false,
	},
}
