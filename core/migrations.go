package core

import (
	"github.com/BitInByte/web-app-example/model"
	"gorm.io/gorm"
)

func Migrations(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Todo{})
}
