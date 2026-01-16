package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var publicationsRoutes = []Route{
	{
		URI:      "/publications",
		Method:   http.MethodPost,
		Function: controllers.CreatePublications,
		NeedAuth: true,
	},
	{
		URI:      "/publications/{publicationID}/like",
		Method:   http.MethodPost,
		Function: controllers.LikePublication,
		NeedAuth: true,
	},
	{
		URI:      "/publications/{publicationID}/dislike",
		Method:   http.MethodPost,
		Function: controllers.LikePublication,
		NeedAuth: true,
	},
	{
		URI:      "/publications/{publicationID}/update",
		Method:   http.MethodGet,
		Function: controllers.LoadUpdatePublicationPage,
		NeedAuth: true,
	},
	{
		URI:      "/publications/{publicationID}",
		Method:   http.MethodPut,
		Function: controllers.UpdatePublication,
		NeedAuth: true,
	},
}
