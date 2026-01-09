package controllers

import (
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/request"
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
	response, error := request.RequestWithAuth(r, http.MethodGet, url, nil)
	fmt.Println(response, error)

	utils.ExecuteHtmlTemplate(w, "home.html", nil)
}
