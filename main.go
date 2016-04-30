package main

import (
	"log"
	"net/http"

	"atlas-api/config/db"
	"atlas-api/route"
)

func main() {
	database, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	db.Migrate(database)
	router := route.NewRouter()
	log.Fatal(http.ListenAndServe(":3000", router))
}
