package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
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

	url := fmt.Sprintf("%s/login", config.APIURL)
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

	var authData models.AuthData
	if error = json.NewDecoder(getResponse.Body).Decode(&authData); error != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErrorAPI{Error: error.Error()})
		return
	}

	if error = cookies.Save(w, authData.ID, authData.Token); error != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErrorAPI{Error: error.Error()})
		return
	}

	response.JSON(w, http.StatusOK, nil)
}
