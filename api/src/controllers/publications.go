package controllers

import (
	"api/src/authentication"
	"api/src/localDatabase"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"io"
	"net/http"
)

// CreatePublications insert a publication into the database
func CreatePublications(w http.ResponseWriter, r *http.Request) {
	userID, error := authentication.ExtractUserId(r)
	if error != nil {
		response.Error(w, http.StatusUnauthorized, error)
		return
	}

	bodyRequest, error := io.ReadAll(r.Body)
	if error != nil {
		response.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var publication models.Publication
	if error = json.Unmarshal(bodyRequest, &publication); error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	publication.AuthorID = userID

	db, error := localDatabase.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationRepository(db)
	publication.ID, error = repository.Create(publication)
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	response.JSON(w, http.StatusOK, publication)
}

// SearchPublications selects all publications from user and their followers
func SearchPublications(w http.ResponseWriter, r *http.Request) {}

// SearchPublication select a publication from database
func SearchPublication(w http.ResponseWriter, r *http.Request) {}

// UpdatePublications update publication values into the database
func UpdatePublications(w http.ResponseWriter, r *http.Request) {}

// DeletePublications delete publication from the database
func DeletePublications(w http.ResponseWriter, r *http.Request) {}
