package database

import (
	"database/sql"
)

func CreateDBConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=postgres password=somepass dbname=projects sslmode=disable")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
