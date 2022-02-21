package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "test.db") // Open initialize a new db connection - requires a driver
	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&Book{}) // migrate the database schema for the Book model

	DB = database // populate with our database instance
}
