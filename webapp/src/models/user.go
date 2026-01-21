package models

import "time"

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
