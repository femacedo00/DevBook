package controllers

import "net/http"

// CreatePublications insert a publication into the database
func CreatePublications(w http.ResponseWriter, r *http.Request) {}

// SearchPublications selects all publications from user and their followers
func SearchPublications(w http.ResponseWriter, r *http.Request) {}

// SearchPublication select a publication from database
func SearchPublication(w http.ResponseWriter, r *http.Request) {}

// UpdatePublications update publication values into the database
func UpdatePublications(w http.ResponseWriter, r *http.Request) {}

// DeletePublications delete publication from the database
func DeletePublications(w http.ResponseWriter, r *http.Request) {}
