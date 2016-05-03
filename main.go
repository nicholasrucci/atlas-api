package main

import (
	"log"
	"net/http"

	"atlas-api/db"
	"atlas-api/route"
)

func main() {
	err := db.InitializeConnection()
	if err != nil {
		log.Fatal(err)
	}

	router := route.NewRouter()

	log.Fatal(http.ListenAndServe(":3000", router))
}
