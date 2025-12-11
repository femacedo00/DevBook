package routes

import (
	"api/src/router/controllers"
	"net/http"
)

var UserRoutes = []Route{
	{
		URI:      "/usuarios",
		Method:   http.MethodPost,
		Function: controllers.CreateUser,
		NeedAuth: false,
	},
	{
		URI:      "/usuarios",
		Method:   http.MethodGet,
		Function: controllers.SearchUsers,
		NeedAuth: false,
	},
	{
		URI:      "/usuarios/{userId}",
		Method:   http.MethodGet,
		Function: controllers.SearchUser,
		NeedAuth: false,
	},
	{
		URI:      "/usuarios/{userId}",
		Method:   http.MethodPut,
		Function: controllers.UpdateUser,
		NeedAuth: false,
	},
	{
		URI:      "/usuarios/{userId}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteUser,
		NeedAuth: false,
	},
}
