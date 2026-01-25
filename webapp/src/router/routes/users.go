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
		NeedAuth: true,
	},
	{
		URI:      "/users/{userId}",
		Method:   http.MethodGet,
		Function: controllers.LoadUserProfile,
		NeedAuth: true,
	},
	{
		URI:      "/users/{userId}/unfollow",
		Method:   http.MethodPost,
		Function: controllers.UnfollowUser,
		NeedAuth: true,
	},
	{
		URI:      "/users/{userId}/follow",
		Method:   http.MethodPost,
		Function: controllers.FollowUser,
		NeedAuth: true,
	},
	{
		URI:      "/profile",
		Method:   http.MethodGet,
		Function: controllers.LoadLoggedInProfilePage,
		NeedAuth: true,
	},
	{
		URI:      "/edit-user",
		Method:   http.MethodGet,
		Function: controllers.LoadEditProfilePage,
		NeedAuth: true,
	},
	{
		URI:      "/edit-user",
		Method:   http.MethodPut,
		Function: controllers.EditUser,
		NeedAuth: true,
	},
	{
		URI:      "/update-password",
		Method:   http.MethodGet,
		Function: controllers.LoadUpdatePasswordPage,
		NeedAuth: true,
	},
	{
		URI:      "/update-password",
		Method:   http.MethodPost,
		Function: controllers.UpdatePassword,
		NeedAuth: true,
	},
	{
		URI:      "/delete-user",
		Method:   http.MethodDelete,
		Function: controllers.DeletePassword,
		NeedAuth: true,
	},
}
