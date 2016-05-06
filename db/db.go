package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=nicholasrucci dbname=atlas sslmode=disable")
	if err != nil {
		defer db.Close()
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		defer db.Close()
		return db, err
	}
	return db, err
}
