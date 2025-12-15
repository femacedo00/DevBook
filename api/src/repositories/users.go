package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

// NewUserRepositorie create a user repository
func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

// Create inserts a new user into the database
func (repository users) Create(user models.User) (uint64, error) {
	statement, error := repository.db.Prepare(
		"insert into users (name, nick, email, password) values(?, ?, ?, ?)",
	)
	if error != nil {
		return 0, error
	}
	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if error != nil {
		return 0, error
	}

	lastId, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}

	return uint64(lastId), nil
}

// Search returns users matching a given name or nickname
func (repository users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	lines, error := repository.db.Query(
		"select id, name, nick, email, createdIn from users where name like ? or nick like ?",
		nameOrNick,
		nameOrNick,
	)

	if error != nil {
		return nil, error
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if error = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedIn,
		); error != nil {
			return nil, error
		}

		users = append(users, user)
	}

	return users, nil
}
