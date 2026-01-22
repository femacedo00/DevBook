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
	{
		URI:      "/users/{userId}",
		Method:   http.MethodGet,
		Function: controllers.LoadUserProfile,
		NeedAuth: false,
	},
	{
		URI:      "/users/{userId}/unfollow",
		Method:   http.MethodPost,
		Function: controllers.UnfollowUser,
		NeedAuth: false,
	},
	{
		URI:      "/users/{userId}/follow",
		Method:   http.MethodPost,
		Function: controllers.FollowUser,
		NeedAuth: false,
	},
}
