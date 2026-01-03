package repositories

import (
	"api/src/models"
	"database/sql"
)

// Publications represents the publication repository
type Publications struct {
	db *sql.DB
}

// NewPublicationRepository creates a new publication repository
func NewPublicationRepository(db *sql.DB) *Publications {
	return &Publications{db}
}

// Create inserts a publication into database
func (repository Publications) Create(publication models.Publication) (uint64, error) {
	statement, error := repository.db.Prepare(
		"insert into publications (title, content, author_id) values (?, ?, ?)",
	)
	if error != nil {
		return 0, error
	}
	defer statement.Close()

	result, error := statement.Exec(publication.Title, publication.Content, publication.AuthorID)
	if error != nil {
		return 0, error
	}

	lastId, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}

	return uint64(lastId), nil
}

// SearchID returns all publications from user and their followers
func (repository Publications) Search(userID uint64) ([]models.Publication, error) {
	lines, error := repository.db.Query(`
		select distinct p.*, u.nick
		from publications p 
		join users u
		on p.author_id = u.id
		left join followers f
		on u.id = f.user_id
		where u.id = ? or f.follower_id = ?
		order by 1 desc
	`, userID, userID)
	if error != nil {
		return nil, error
	}
	defer lines.Close()

	var publications []models.Publication

	for lines.Next() {
		var publication models.Publication

		if error = lines.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedIn,
			&publication.AuthorNick,
		); error != nil {
			return nil, error
		}

		publications = append(publications, publication)
	}
	return publications, nil
}

// SearchID returns a publication matching an id
func (repository Publications) SearchID(PublicationId uint64) (models.Publication, error) {
	lines, error := repository.db.Query(`
		select p.*, u.nick
		from publications p 
		join users u
		on p.author_id = u.id
		where p.id = ?
	`, PublicationId)
	if error != nil {
		return models.Publication{}, error
	}
	defer lines.Close()

	var publication models.Publication

	if lines.Next() {
		if error = lines.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedIn,
			&publication.AuthorNick,
		); error != nil {
			return models.Publication{}, error
		}
	}
	return publication, nil
}

// SearchUserID returns all publications matching a user id
func (repository Publications) SearchUserID(userId uint64) ([]models.Publication, error) {
	lines, error := repository.db.Query(`
		select p.*, u.nick
		from publications p 
		join users u
		on p.author_id = u.id
		where p.author_id = ?
		order by 1 desc
	`, userId)
	if error != nil {
		return nil, error
	}
	defer lines.Close()

	var publications []models.Publication

	for lines.Next() {
		var publication models.Publication

		if error = lines.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedIn,
			&publication.AuthorNick,
		); error != nil {
			return nil, error
		}

		publications = append(publications, publication)
	}
	return publications, nil
}

// Update changes a publication informations into database
func (repository Publications) Update(publicationID uint64, publication models.Publication) error {
	statement, error := repository.db.Prepare(
		"update publications set title = ?, content = ? where id = ?",
	)
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error := statement.Exec(publication.Title, publication.Content, publicationID); error != nil {
		return error
	}

	return nil
}

// Delete a publication matching by an id
func (repository Publications) Delete(publicationID uint64) error {
	statement, error := repository.db.Prepare(
		"delete from publications where id = ?",
	)
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error := statement.Exec(publicationID); error != nil {
		return error
	}

	return nil
}
