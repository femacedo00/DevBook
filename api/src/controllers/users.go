package controllers

import (
	"api/src/localDatabase"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CreateUser insert a user into the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, error := io.ReadAll(r.Body)
	if error != nil {
		log.Fatal(error)
	}

	var user models.User
	if error = json.Unmarshal(bodyRequest, &user); error != nil {
		log.Fatal(error)
	}

	db, error := localDatabase.Connect()
	if error != nil {
		log.Fatal(error)
	}

	repository := repositories.NewUserRepository(db)
	userId, error := repository.Create(user)
	if error != nil {
		log.Fatal(error)
	}

	w.Write(fmt.Appendf(nil, "Id entered: %d", userId))
}

// SearchUsers select all users from the database
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching all users"))
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
