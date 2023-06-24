package core

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// var DB *gorm.DB

func LoadSqliteDBSettings() *gorm.DB {
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "file::memory:?cache=shared"
	} else {
		dbName = dbName + ".db"
	}

	var err error
	var db *gorm.DB
	db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})

	if err != nil {
		panic("Failed to set connection to database")
	}

	return db
}
