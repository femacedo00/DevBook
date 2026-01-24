package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/request"
	"webapp/src/response"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

// LoadLoginPage loads login page
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}

	utils.ExecuteHtmlTemplate(w, "login.html", nil)
}

// LoadUserRegisterPage loads user register page
func LoadUserRegisterPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}

	utils.ExecuteHtmlTemplate(w, "register.html", nil)
}

// LoadHomePage loads home page
func LoadHomePage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publications", config.APIURL)
	getResponse, error := request.RequestWithAuth(r, http.MethodGet, url, nil)
	if error != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorAPI{Error: error.Error()})
		return
	}
	defer getResponse.Body.Close()

	if getResponse.StatusCode >= 400 {
		cookies.Delete(w)
		http.Redirect(w, r, "/login", 302)
		return
	}

	var publications []models.Publication
	if error = json.NewDecoder(getResponse.Body).Decode(&publications); error != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErrorAPI{Error: error.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecuteHtmlTemplate(w, "home.html", struct {
		Publications []models.Publication
		UserID       uint64
	}{
		Publications: publications,
		UserID:       userID,
	})
}

// LoadUpdatePublicationPage loads edit publication page
func LoadUpdatePublicationPage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationID, error := strconv.ParseUint(params["publicationID"], 10, 64)
	if error != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d", config.APIURL, publicationID)
	getResponse, error := request.RequestWithAuth(r, http.MethodGet, url, nil)
	if error != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorAPI{Error: error.Error()})
		return
	}
	defer getResponse.Body.Close()

	if getResponse.StatusCode >= 400 {
		http.Redirect(w, r, "/home", 302)
		return
	}

	var publications models.Publication
	if error = json.NewDecoder(getResponse.Body).Decode(&publications); error != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErrorAPI{Error: error.Error()})
		return
	}

	utils.ExecuteHtmlTemplate(w, "update-publication.html", publications)
}

// LoadUsersPages loads the page that displays users based on a search filter
func LoadUsersPages(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	url := fmt.Sprintf("%s/users?user=%s", config.APIURL, nameOrNick)

	getResponse, error := request.RequestWithAuth(r, http.MethodGet, url, nil)
	if error != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorAPI{Error: error.Error()})
		return
	}
	defer getResponse.Body.Close()

	if getResponse.StatusCode >= 400 {
		http.Redirect(w, r, "/home", 302)
		return
	}

	var users []models.User
	if error = json.NewDecoder(getResponse.Body).Decode(&users); error != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErrorAPI{Error: error.Error()})
		return
	}

	utils.ExecuteHtmlTemplate(w, "users.html", users)
}

// LoadUserProfile loads the user's profile page
func LoadUserProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, error := strconv.ParseUint(params["userId"], 10, 64)
	if error != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: error.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	loggedInUserID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if userId == loggedInUserID {
		http.Redirect(w, r, "/profile", 302)
		return
	}

	user, error := models.SearchCompleteUser(userId, r)
	if error != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorAPI{Error: error.Error()})
		return
	}

	utils.ExecuteHtmlTemplate(w, "user.html", struct {
		User           models.User
		LoggedInUserID uint64
	}{
		User:           user,
		LoggedInUserID: loggedInUserID,
	})
}

// LoadLoggedInProfilePage loads the logged in user's profile page
func LoadLoggedInProfilePage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	user, error := models.SearchCompleteUser(userID, r)
	if error != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorAPI{Error: error.Error()})
		return
	}

	utils.ExecuteHtmlTemplate(w, "profile.html", user)
}

// LoadEditProfilePage loads the page to edit the user's profile
func LoadEditProfilePage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	channel := make(chan models.User)
	go models.SearchUserData(channel, userID, r)
	user := <-channel

	if user.ID == 0 {
		response.JSON(w, http.StatusInternalServerError, response.ErrorAPI{Error: "Error searching user"})
	}

	utils.ExecuteHtmlTemplate(w, "edit-user.html", user)
}
