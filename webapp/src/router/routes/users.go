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
	{
		URI:      "/users",
		Method:   http.MethodPost,
		Function: controllers.UserRegister,
		NeedAuth: false,
	},
	{
		URI:      "/search-users",
		Method:   http.MethodGet,
		Function: controllers.LoadUsersPages,
		NeedAuth: false,
	},
}
