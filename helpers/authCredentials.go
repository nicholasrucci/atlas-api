package helper

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"

	"atlas-api/config/schema"
	"golang.org/x/crypto/scrypt"
)

// Credentials is a struct that will be returned from the methods below
// containing the salt and hash of a password
type Credentials struct {
	Salt string
	Hash string
}

var (
	// PasswordSaltBytes is the number of bytes a salt will be
	PasswordSaltBytes = 32

	// PasswordHashBytes is the number of bytes a hash will be
	PasswordHashBytes = 64
)

// CreateCredentials will hash the password and return the hash
// and the salt used when creating a user
func CreateCredentials(password string) (Credentials, error) {
	var credentials Credentials
	salt := make([]byte, PasswordSaltBytes)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return credentials, err
	}

	hash, err := scrypt.Key([]byte(password), salt, 1<<14, 8, 1, PasswordHashBytes)
	if err != nil {
		return credentials, err
	}

	credentials.Hash = hex.EncodeToString(hash)
	credentials.Salt = hex.EncodeToString(salt)

	return credentials, err
}

// Compare will hash the password and then compare it to the
// credentials that were passed down with it
func Compare(password string, user schema.User) error {
	userPassword, err := hex.DecodeString(user.PasswordSalt)
	if err != nil {
		return err
	}

	hash, err := scrypt.Key([]byte(password), userPassword, 1<<14, 8, 1, PasswordHashBytes)
	if err != nil {
		return err
	}

	if hex.EncodeToString(hash) == user.PasswordHash {
		return nil
	}
	return errors.New("Invalid credentials")
}
