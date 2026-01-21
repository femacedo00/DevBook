package models

import (
	"net/http"
	"time"
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
	channelUser := make(chan User)
	channelFollowers := make(chan []User)
	channelFollowing := make(chan []User)
	channelPublications := make(chan []Publication)

	go SearchUserData(channelUser, userID, r)
	go SearchFollowersData(channelFollowers, userID, r)
	go SearchFollowingData(channelFollowing, userID, r)
	go SearchPublicationsData(channelPublications, userID, r)

	return User{}, nil
}

func SearchUserData(channel <-chan User, userID uint64, r *http.Request) {}

func SearchFollowersData(channel <-chan []User, userID uint64, r *http.Request) {}

func SearchFollowingData(channel <-chan []User, userID uint64, r *http.Request) {}

func SearchPublicationsData(channel <-chan []Publication, userID uint64, r *http.Request) {}
