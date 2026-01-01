package models

// Password is the struct for a password change request
type Password struct {
	New     string `json:"new"`
	Current string `json:"current"`
}
