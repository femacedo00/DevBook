package routes

import (
	"api/src/controllers"
	"net/http"
)

var UserRoutes = []Route{
	{
		URI:      "/users",
		Method:   http.MethodPost,
		Function: controllers.CreateUser,
		NeedAuth: false,
	},
	{
		URI:      "/users",
		Method:   http.MethodGet,
		Function: controllers.SearchUsers,
		NeedAuth: false,
	},
	{
		URI:      "/users/{userID}",
		Method:   http.MethodGet,
		Function: controllers.SearchUser,
		NeedAuth: false,
	},
	{
		URI:      "/users/{userID}",
		Method:   http.MethodPut,
		Function: controllers.UpdateUser,
		NeedAuth: false,
	},
	{
		URI:      "/users/{userID}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteUser,
		NeedAuth: false,
	},
}
