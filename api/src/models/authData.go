package models

// AuthData contains the ID and token of an authenticated user
type AuthData struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
