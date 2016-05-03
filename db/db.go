package db

import (
	"atlas-api/config/schema"

	"github.com/jinzhu/gorm"

	// dialect for gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

// InitializeConnection will open the applications database and return
// it and a possible error
func InitializeConnection() error {
	var err error
	DB, err = gorm.Open("mysql", "root:@/atlas?charset=utf8&parseTime=True&loc=Local")
	return err
}

// Migrate will go through each of the tables and migrate
// them if they have not yet been migrated
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&schema.User{})
}
