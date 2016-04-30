package main

import (
	"log"
	"net/http"

	"atlas-api/db"
	"atlas-api/route"
)

func main() {
	database, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}

	db.Migrate(database)

	err = database.Close()
	if err != nil {
		log.Fatal(err)
	}

	router := route.NewRouter()

	log.Fatal(http.ListenAndServe(":3000", router))
}
