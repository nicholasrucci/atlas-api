package main

import (
	"net/http"
	"log"

	"atlas-api/route"
	"atlas-api/config/db"
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
