package db

import (
	"database/sql"
	"log"

	"github.com/DavidHuie/gomigrate"
	"github.com/Sirupsen/logrus"
	_ "github.com/lib/pq"
)

func init() {
	db, err := sql.Open("postgres", "user=nicholasrucci dbname=atlas sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	migrator, err := gomigrate.NewMigratorWithLogger(db, gomigrate.Postgres{}, "./config/db/migrations", logrus.New())
	if err != nil {
		log.Fatal(err)
	}

	err = migrator.Migrate()
	if err != nil {
		log.Fatal(err)
	}
}
