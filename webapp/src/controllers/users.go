package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
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

// EditUser calls an API to edit a user in the database
func EditUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, error := json.Marshal(map[string]string{
		"name":  r.FormValue("name"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
	})
	if error != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: error.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)
	getResponse, error := request.RequestWithAuth(r, http.MethodPut, url, bytes.NewBuffer(user))
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

// UpdatePassword calls an API to update a user's password in the database
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	passwords, error := json.Marshal(map[string]string{
		"current": r.FormValue("current"),
		"new":     r.FormValue("new"),
	})
	if error != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: error.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d/update-password", config.APIURL, userID)
	getResponse, error := request.RequestWithAuth(r, http.MethodPost, url, bytes.NewBuffer(passwords))
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

// DeletePassword calls an API to delete the user in the database
func DeletePassword(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)
	getResponse, error := request.RequestWithAuth(r, http.MethodDelete, url, nil)
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
