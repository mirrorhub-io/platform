package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var conn *gorm.DB

func Connection() *gorm.DB {
	if conn == nil {
		db_init()
	}
	return conn
}

func db_init() {
	db, _ := gorm.Open("sqlite3", "/tmp/gorm.db")
	db.AutoMigrate(&Mirror{})
	db.AutoMigrate(&Service{})
	db.AutoMigrate(&Contact{})
	conn = db
}
