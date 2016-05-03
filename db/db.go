package db

import (
	"github.com/jinzhu/gorm"

	// dialect for gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB is the global database
var DB *gorm.DB

// InitializeConnection will open the applications database and return
// it and a possible error
func InitializeConnection() error {
	var err error
	DB, err = gorm.Open("mysql", "root:@/atlas?charset=utf8&parseTime=True&loc=Local")
	return err
}
