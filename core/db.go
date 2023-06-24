package core

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadSqliteDBSettings() {
    dbName := os.Getenv("DB_NAME")
    if dbName == "" {
        dbName = "file::memory:?cache=shared"
    } else {
        dbName = dbName + ".db"
    }

	var err error
	DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})

	if err != nil {
		panic("Failed to set connection to database")
	}
}
