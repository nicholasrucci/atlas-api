package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"

	"golang.org/x/crypto/scrypt"
)

const (
	PW_SALT_BYTES = 32
	PW_HASH_BYTES = 64

	password       = "hello"
	other_password = "hello"
)

type User struct {
	hash string
	salt string
}

func main() {

	salt := make([]byte, PW_SALT_BYTES)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		log.Fatal(err)
	}

	texthash, err := scrypt.Key([]byte(other_password), salt, 1<<14, 8, 1, PW_HASH_BYTES)
	if err != nil {
		log.Fatal(err)
	}

	hash, err := scrypt.Key([]byte(password), salt, 1<<14, 8, 1, PW_HASH_BYTES)
	if err != nil {
		log.Fatal(err)
	}

	user := User{
		hash: hex.EncodeToString(hash),
		salt: string(salt),
	}

	if user.hash == hex.EncodeToString(texthash) {
		fmt.Println("Works")
	} else {
		fmt.Println("Doesn't work")
	}

	// fmt.Println(user.hash)
	// fmt.Print(hex.EncodeToString(texthash))

}
