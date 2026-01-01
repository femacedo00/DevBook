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
		NeedAuth: true,
	},
	{
		URI:      "/users/{userId}",
		Method:   http.MethodGet,
		Function: controllers.SearchUser,
		NeedAuth: true,
	},
	{
		URI:      "/users/{userId}",
		Method:   http.MethodPut,
		Function: controllers.UpdateUser,
		NeedAuth: true,
	},
	{
		URI:      "/users/{userId}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteUser,
		NeedAuth: true,
	},
	{
		URI:      "/users/{userId}/follow",
		Method:   http.MethodPost,
		Function: controllers.FollowUser,
		NeedAuth: true,
	},
	{
		URI:      "/users/{userId}/unfollow",
		Method:   http.MethodPost,
		Function: controllers.UnfollowUser,
		NeedAuth: true,
	},
	{
		URI:      "/users/{userId}/followers",
		Method:   http.MethodGet,
		Function: controllers.SearchFollowers,
		NeedAuth: true,
	},
	{
		URI:      "/users/{userId}/following",
		Method:   http.MethodGet,
		Function: controllers.SearchFollowing,
		NeedAuth: true,
	},
}
