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

// UserReq ...
type UserReq struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

// CreateUser will accept a POST request and add a user to the database
func CreateUser(rw http.ResponseWriter, req *http.Request) {
	var user schema.User
	var userReq UserReq
	var credentials helper.Credentials

	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		helper.UserResponse(rw, req, 200, user, err)
		return
	}
	if err := req.Body.Close(); err != nil {
		helper.UserResponse(rw, req, 200, user, err)
		return
	}

	if err := json.Unmarshal(body, &userReq); err != nil {
		helper.UserResponse(rw, req, 200, user, err)
		return
	}

	credentials, err = helper.CreateCredentials(userReq.Password)
	if err != nil {
		log.Fatal(err)
	}

	user.FirstName = userReq.FirstName
	user.LastName = userReq.LastName
	user.Email = userReq.Email
	user.PasswordHash = credentials.Hash
	user.PasswordSalt = credentials.Salt

	database, err := db.Connection()
	if err != nil {
		helper.UserResponse(rw, req, 200, user, err)
		return
	}

	_, err = database.Query("INSERT INTO users(first_name, last_name, email, password_hash, password_salt, disabled) VALUES($1, $2, $3, $4, $5, $6)",
		user.FirstName,
		user.LastName,
		user.Email,
		user.PasswordHash,
		user.PasswordSalt,
		false,
	)
	if err != nil {
		helper.UserResponse(rw, req, 200, user, err)
		database.Close()
		return
	}

	err = database.Close()
	if err != nil {
		helper.UserResponse(rw, req, 200, user, err)
		return
	}

	helper.UserResponse(rw, req, 200, user, err)
}
