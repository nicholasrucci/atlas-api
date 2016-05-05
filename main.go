package main

import (
	"log"
	"net/http"

	_ "atlas-api/config/db"
	"atlas-api/route"
)

func main() {
	router := route.NewRouter()

	log.Fatal(http.ListenAndServe(":3000", router))
}
