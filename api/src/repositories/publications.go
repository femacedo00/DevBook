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
