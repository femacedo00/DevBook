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
}
