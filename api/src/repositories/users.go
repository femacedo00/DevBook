package repositories

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

// NewUserRepositorie create a user repository
func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

// Create inserts a new user into the database
func (u users) Create(users models.User) (uint64, error) {
	return 0, nil
}
