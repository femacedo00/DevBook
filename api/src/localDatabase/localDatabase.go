package localDatabase

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver
)

// Connect opens and returns the database connection
func Connect() (*sql.DB, error) {
	db, error := sql.Open("mysql", config.ConnectionStringDB)
	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		db.Close()
		return nil, error
	}

	return db, nil
}
