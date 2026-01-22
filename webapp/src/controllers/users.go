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

// UserRegister calls an API to register a user in the database
func UserRegister(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, error := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"nick":     r.FormValue("nick"),
		"password": r.FormValue("password"),
	})

	if error != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/users", config.APIURL)
	getResponse, error := http.Post(url, "application/json", bytes.NewBuffer(user))
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

// UnfollowUser calls an API to unfollow a user in the database
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, error := strconv.ParseUint(params["userId"], 10, 64)
	if error != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/unfollow", config.APIURL, userID)
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

// FollowUser calls an API to follow a user in the database
func FollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, error := strconv.ParseUint(params["userId"], 10, 64)
	if error != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/follow", config.APIURL, userID)
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
