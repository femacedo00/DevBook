package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var userRoutes = []Route{
	{
		URI:      "/user-register",
		Method:   http.MethodGet,
		Function: controllers.LoadUserRegisterPage,
		NeedAuth: false,
	},
}
