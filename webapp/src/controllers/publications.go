package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/request"
	"webapp/src/response"
)

// CreatePublication insert a publication into the database
func CreatePublications(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	publication, error := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	if error != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications", config.APIURL)
	getResponse, error := request.RequestWithAuth(r, http.MethodPost, url, bytes.NewBuffer(publication))

	if error != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorAPI{Error: error.Error()})
		return
	}
	defer getResponse.Body.Close()

	if getResponse.StatusCode >= 400 {
		response.HandleErrorStatusCode(w, getResponse)
		return
	}

	response.JSON(w, getResponse.StatusCode, nil)
}
