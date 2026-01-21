package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/request"
)

// User represents a person using the application
type User struct {
	ID           uint64        `json:"id"`
	Name         string        `json:"name"`
	Nick         string        `json:"nick"`
	Email        string        `json:"email"`
	CreatedIn    time.Time     `json:"createdIn"`
	Followers    []User        `json:"followers"`
	Following    []User        `json:"following"`
	Publications []Publication `json:"publications"`
}

// SearchCompleteUser requests four different APIs to aggregate user information
func SearchCompleteUser(userID uint64, r *http.Request) (User, error) {
	userChannel := make(chan User)
	followersChannel := make(chan []User)
	followingChannel := make(chan []User)
	publicationsChannel := make(chan []Publication)

	go SearchUserData(userChannel, userID, r)
	go SearchFollowersData(followersChannel, userID, r)
	go SearchFollowingData(followingChannel, userID, r)
	go SearchPublicationsData(publicationsChannel, userID, r)

	var (
		user         User
		followers    []User
		following    []User
		publications []Publication
	)

	for i := 0; i < 4; i++ {
		select {
		case loadedUser := <-userChannel:
			if loadedUser.ID == 0 {
				return User{}, errors.New("Error searching user!")
			}

			user = loadedUser
		case loadedFollowers := <-followersChannel:
			if loadedFollowers == nil {
				return User{}, errors.New("Error searching followers!")
			}

			followers = loadedFollowers
		case loadedFollowing := <-followingChannel:
			if loadedFollowing == nil {
				return User{}, errors.New("Error searching who user is following!")
			}

			following = loadedFollowing
		case loadedPublications := <-publicationsChannel:
			if loadedPublications == nil {
				return User{}, errors.New("Error searching publications!")
			}

			publications = loadedPublications
		}
	}

	user.Followers = followers
	user.Following = following
	user.Publications = publications

	return user, nil
}

// SearchUserData retrieves the main user data from the API.
func SearchUserData(channel chan<- User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)
	getResponse, error := request.RequestWithAuth(r, http.MethodGet, url, nil)
	if error != nil {
		channel <- User{}
		return
	}
	defer getResponse.Body.Close()

	var user User
	if error = json.NewDecoder(getResponse.Body).Decode(&user); error != nil {
		channel <- User{}
		return
	}

	channel <- user
}

// SearchFollowersData retrieves the user followers data from the API.
func SearchFollowersData(channel chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.APIURL, userID)
	getResponse, error := request.RequestWithAuth(r, http.MethodGet, url, nil)
	if error != nil {
		channel <- nil
		return
	}
	defer getResponse.Body.Close()

	var followers []User
	if error = json.NewDecoder(getResponse.Body).Decode(&followers); error != nil {
		channel <- nil
		return
	}

	channel <- followers
}

// SearchFollowingData retrieves the user following data from the API.
func SearchFollowingData(channel chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/following", config.APIURL, userID)
	getResponse, error := request.RequestWithAuth(r, http.MethodGet, url, nil)
	if error != nil {
		channel <- nil
		return
	}
	defer getResponse.Body.Close()

	var following []User
	if error = json.NewDecoder(getResponse.Body).Decode(&following); error != nil {
		channel <- nil
		return
	}

	channel <- following
}

// SearchPublicationsData retrieves the user publications data from the API.
func SearchPublicationsData(channel chan<- []Publication, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/publications", config.APIURL, userID)
	getResponse, error := request.RequestWithAuth(r, http.MethodGet, url, nil)
	if error != nil {
		channel <- nil
		return
	}
	defer getResponse.Body.Close()

	var publications []Publication
	if error = json.NewDecoder(getResponse.Body).Decode(&publications); error != nil {
		channel <- nil
		return
	}

	channel <- publications
}
