package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"atlas-api/config/db"
	"atlas-api/middleware"
)

// AuthenticatePostData will hold the email and password of the request
// that was sent up by the clientf
type AuthenticatePostData struct {
	Email    string
	Password string
}

// Authenticate - POST
// Will accept an Email and Password. Query the database for Email
// and grab salt and hash from user in the database with the same
// Email. It will then hash the requested password with the existing
// salt and compare the two.
func Authenticate(rw http.ResponseWriter, req *http.Request) {
	var data AuthenticatePostData
	var user db.User

	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		log.Fatal(err)
	}

	err = req.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
		rw.WriteHeader(422)
		// TODO: Handle error correctly
		return
	}

	database, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}

	if err = database.Where("email = ?", data.Email).Find(&user).Error; err != nil {
		rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
		rw.WriteHeader(404)
		database.Close()
		// TODO: Handle error correctly
		return
	}

	middleware.JSONHandler(rw, req)
	err = json.NewEncoder(rw).Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}
