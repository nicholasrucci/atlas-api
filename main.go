package main

import (
	"log"
	"net/http"

	"atlas-api/config/db"
	"atlas-api/config/schema"
	"atlas-api/route"
)

func main() {
	err := db.InitializeConnection()
	if err != nil {
		log.Fatal(err)
	}

	db.DB.AutoMigrate(&schema.User{}, &schema.Organization{}, &schema.Page{}, &schema.Platform{}, &schema.Task{}, &schema.Group{})

	router := route.NewRouter()

	log.Fatal(http.ListenAndServe(":3000", router))
}
