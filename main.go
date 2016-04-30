package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	FirstName    string `gorm:"size:255"`
	LastName     string `gorm:"size:255"`
	Email        string `gorm:"type:varchar(100);unique"`
	PasswordHash string `gorm:"size:255"`
	PasswordSalt string `gorm:"size:255"`
	Disabled     bool
}

func main() {
	db, err := gorm.Open("mysql", "root:@/atlas?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Print(err)
	}

	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
	}
	db.Create(&User{FirstName: "Nick", Email: "none"})

	// salt := make([]byte, PW_SALT_BYTES)
	// _, err := io.ReadFull(rand.Reader, salt)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// texthash, err := scrypt.Key([]byte(other_password), salt, 1<<14, 8, 1, PW_HASH_BYTES)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// hash, err := scrypt.Key([]byte(password), salt, 1<<14, 8, 1, PW_HASH_BYTES)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// if user.hash == hex.EncodeToString(texthash) {
	// 	 fmt.Println("Works")
	// } else {
	// 	 fmt.Println("Doesn't work")
	// }
	//
	// fmt.Println(user.hash)
	// fmt.Print(hex.EncodeToString(texthash))

}
