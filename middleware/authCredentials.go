package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"

	"golang.org/x/crypto/scrypt"
)

var (
	// PasswordSaltBytes is the number of bytes a salt will be
	PasswordSaltBytes = 32

	// PasswordHashBytes is the number of bytes a hash will be
	PasswordHashBytes = 64
)

// CreateCredentials will hash the password and return the hash
// and the salt used when creating a user
func CreateCredentials(password string) (string, string, error) {
	salt := make([]byte, PasswordSaltBytes)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return "", "", err
	}

	hash, err := scrypt.Key([]byte(password), salt, 1<<14, 8, 1, PasswordHashBytes)
	if err != nil {
		return "", "", err
	}

	return hex.EncodeToString(salt), hex.EncodeToString(hash), err
}
