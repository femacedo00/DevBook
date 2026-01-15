package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/request"
	"webapp/src/response"
	"webapp/src/utils"
)

// LoadLoginPage loads login page
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteHtmlTemplate(w, "login.html", nil)
}

// LoadUserRegisterPage loads user register page
func LoadUserRegisterPage(w http.ResponseWriter, r *http.Request) {
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
		response.HandleErrorStatusCode(w, getResponse)
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
