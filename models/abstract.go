package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

var conn *gorm.DB

func Connection() *gorm.DB {
	if conn == nil {
		db_init()
	}
	return conn
}

func db_init() {
	path := "/tmp/mirrorhub.db"
	if os.Getenv("DB_PATH") != "" {
		path = os.Getenv("DB_PATH")
	}
	db, _ := gorm.Open("sqlite3", path)
	db.AutoMigrate(&Mirror{})
	db.AutoMigrate(&Service{})
	db.AutoMigrate(&Contact{})
	conn = db
}
