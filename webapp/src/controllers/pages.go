package controllers

import (
	"net/http"
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
