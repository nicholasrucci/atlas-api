package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func init() {
	db, err := sql.Open("postgres", "user=nicholasrucci dbname=atlas sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

}
