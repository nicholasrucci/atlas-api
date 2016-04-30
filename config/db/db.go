package db

import (
	"github.com/jinzhu/gorm"

	// dialect for gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Connection will open the applications database and return
// it and a possible error
func Connection() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:@/atlas?charset=utf8&parseTime=True&loc=Local")
	return db, err
}

// Migrate will go through each of the tables and migrate
// them if they have not yet been migrated
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
