package controllers

import (
	"api/src/localDatabase"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// CreateUser insert a user into the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	if error = user.Prepare(); error != nil {
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
	user.ID, error = repository.Create(user)
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	response.JSON(w, http.StatusCreated, user)
}

// SearchUsers select all users from the database
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, error := localDatabase.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repositories := repositories.NewUserRepository(db)
	users, error := repositories.Search(nameOrNick)
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
	}

	response.JSON(w, http.StatusOK, users)
}

// SearchUser select a user from the database
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching a user"))
}

// UpdateUser update user values in the database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating a user"))
}

// DeleteUser delete user from the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting a user"))
}
