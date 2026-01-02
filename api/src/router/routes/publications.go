package routes

import (
	"api/src/controllers"
	"net/http"
)

var PublicationsRoutes = []Route{
	{
		URI:      "/publications",
		Method:   http.MethodPost,
		Function: controllers.CreatePublications,
		NeedAuth: true,
	},
	{
		URI:      "/publications",
		Method:   http.MethodGet,
		Function: controllers.SearchPublications,
		NeedAuth: true,
	},
	{
		URI:      "/publications/{publicationId}",
		Method:   http.MethodGet,
		Function: controllers.SearchPublication,
		NeedAuth: true,
	},
	{
		URI:      "/publications/{publicationId}",
		Method:   http.MethodPut,
		Function: controllers.UpdatePublications,
		NeedAuth: true,
	},
	{
		URI:      "/publications/{publicationId}",
		Method:   http.MethodDelete,
		Function: controllers.DeletePublications,
		NeedAuth: true,
	},
}
