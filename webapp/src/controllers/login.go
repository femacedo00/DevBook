package controllers

import "net/http"

// LoadLoginPage loads login page
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login Page"))
}
