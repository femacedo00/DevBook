package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"webapp/src/response"
)

// Login uses email and password to authenticate in the application
func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, error := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if error != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: error.Error()})
	}

	getResponse, error := http.Post("http://localhost:5000/login", "application/json", bytes.NewBuffer(user))
	if error != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorAPI{Error: error.Error()})
		return
	}

	token, _ := io.ReadAll(getResponse.Body)
	fmt.Println(getResponse.StatusCode, string(token))
}
