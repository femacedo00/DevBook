package controllers

import "net/http"

// CreateUser insert a user into the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a user"))
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
