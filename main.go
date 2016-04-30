package main

import (
	"net/http"
	"log"

	"atlas-api/route"
)



func main() {
	router := route.NewRouter()
	log.Fatal(http.ListenAndServe(":3000", router))
}
