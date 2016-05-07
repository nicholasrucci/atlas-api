package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"atlas-api/config/schema"
	"atlas-api/db"
	"atlas-api/helpers"
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
	var (
		data   AuthenticatePostData
		user   schema.User
		userID int
	)

	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		helper.UserResponse(rw, req, 500, user, err)
		return
	}

	err = req.Body.Close()
	if err != nil {
		helper.UserResponse(rw, req, 500, user, err)
		return
	}

	if err = json.Unmarshal(body, &data); err != nil {
		helper.UserResponse(rw, req, 500, user, err)
		return
	}

	database, err := db.Connection()
	if err != nil {
		helper.UserResponse(rw, req, 500, user, err)
		return
	}

	rows, err := database.Query("SELECT * FROM users WHERE email=$1", data.Email)
	if err != nil {
		helper.UserResponse(rw, req, 500, user, err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&userID, &user.FirstName, &user.LastName, &user.Email, &user.PasswordHash, &user.PasswordSalt, &user.Disabled)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = database.Close()
	if err != nil {
		helper.UserResponse(rw, req, 500, user, err)
		return
	}

	if err = helper.Compare(data.Password, user); err != nil {
		helper.UserResponse(rw, req, 500, user, err)
		return
	}

	helper.UserResponse(rw, req, 500, user, err)
}
