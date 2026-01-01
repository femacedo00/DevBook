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

// SearchID returns a user matching an id
func (repository users) SearchID(userID uint64) (models.User, error) {
	lines, error := repository.db.Query(
		"select id, name, nick, email, createdIn from users where id = ?",
		userID,
	)

	if error != nil {
		return models.User{}, error
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if error = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedIn,
		); error != nil {
			return models.User{}, error
		}
	}

	return user, nil
}

// SearchEmail returns a user (only id and password) matching an email
func (repository users) SearchEmail(email string) (models.User, error) {
	lines, error := repository.db.Query(
		"select id, password from users where email = ?",
		email,
	)
	if error != nil {
		return models.User{}, error
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if error = lines.Scan(
			&user.ID,
			&user.Password,
		); error != nil {
			return models.User{}, error
		}
	}
	return user, nil
}

// Update changes a user informations
func (repository users) Update(userID uint64, user models.User) error {
	statement, error := repository.db.Prepare(
		"update users set name = ?, nick = ? , email = ? where id = ?",
	)
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(
		user.Name,
		user.Nick,
		user.Email,
		userID,
	); error != nil {
		return error
	}

	return nil
}

// Delete a user matching by an id
func (repository users) Delete(userID uint64) error {
	statement, error := repository.db.Prepare("delete from users where id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(userID); error != nil {
		return error
	}

	return nil
}

// Follow allows a user to follow another user
func (repository users) Follow(userID, followerID uint64) error {
	statement, error := repository.db.Prepare(
		"insert ignore into followers (user_id, follower_id) values (?,?)",
	)
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(userID, followerID); error != nil {
		return error
	}

	return nil
}

// Follow allows a user to unfollow another user
func (repository users) Unfollow(userID, followerID uint64) error {
	statement, error := repository.db.Prepare(
		"delete from followers where user_id = ? and follower_id = ?",
	)
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(userID, followerID); error != nil {
		return error
	}

	return nil
}

// SearchFollowers gets all user followers
func (repository users) SearchFollowers(userID uint64) ([]models.User, error) {
	lines, error := repository.db.Query(`
		select u.id, u.name, u.nick, u.email, u.createdIn
		from users u join followers f
		on  u.id = f.follower_id
		where f.user_id = ?
	`, userID)
	if error != nil {
		return nil, error
	}
	defer lines.Close()

	var followers []models.User

	for lines.Next() {
		var follower models.User

		if error = lines.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
			&follower.CreatedIn,
		); error != nil {
			return nil, error
		}

		followers = append(followers, follower)
	}

	return followers, nil
}

// SearchFollowing returns all users that a specific user is following
func (repository users) SearchFollowing(userID uint64) ([]models.User, error) {
	lines, error := repository.db.Query(`
		select u.id, u.name, u.nick, u.email, u.createdIn
		from users u join followers f
		on  u.id = f.user_id
		where f.follower_id = ?
	`, userID)
	if error != nil {
		return nil, error
	}
	defer lines.Close()

	var following []models.User

	for lines.Next() {
		var follower models.User

		if error = lines.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
			&follower.CreatedIn,
		); error != nil {
			return nil, error
		}

		following = append(following, follower)
	}

	return following, nil
}
