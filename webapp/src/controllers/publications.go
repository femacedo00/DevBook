package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/request"
	"webapp/src/response"

	"github.com/gorilla/mux"
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

// LikePublication saves a like in the database
func LikePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationID, error := strconv.ParseUint(params["publicationID"], 10, 64)

	if error != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d/like", config.APIURL, publicationID)
	getResponse, error := request.RequestWithAuth(r, http.MethodPost, url, nil)

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

// DislikePublication removes a like from the database
func DislikePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationID, error := strconv.ParseUint(params["publicationID"], 10, 64)

	if error != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d/dislike", config.APIURL, publicationID)
	getResponse, error := request.RequestWithAuth(r, http.MethodPost, url, nil)

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
