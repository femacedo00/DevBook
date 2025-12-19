package controllers

import (
	"api/src/authentication"
	"api/src/localDatabase"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Login authenticated a user
func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, error := io.ReadAll(r.Body)
	if error != nil {
		response.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error = json.Unmarshal(bodyRequest, &user); error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := localDatabase.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	userDB, error := repository.SearchEmail(user.Email)
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	if error := security.ValidatePassword(userDB.Password, user.Password); error != nil {
		response.Error(w, http.StatusUnauthorized, error)
		return
	}

	token, error := authentication.CreateToken(userDB.ID)
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	fmt.Println(token)
	w.Write([]byte(token))

}
