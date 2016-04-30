package schema

import (
	"github.com/jinzhu/gorm"
)

// User schema
type User struct {
	gorm.Model
	FirstName    string `gorm:"size:255"`
	LastName     string `gorm:"size:255"`
	Email        string `gorm:"type:varchar(100);unique"`
	PasswordHash string `gorm:"size:255"`
	PasswordSalt string `gorm:"size:255"`
	Disabled     bool
}
