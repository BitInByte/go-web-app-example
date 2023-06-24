package model

import "gorm.io/gorm"

type Status string

const (
	Created    Status = "created"
	InProgress        = "in progress"
	Done              = "done"
)

type Todo struct {
	gorm.Model
	Title string
	Body  string
	// Status Status `sql:"type:ENUM('created', 'in progress', 'done')"`
	// Progress Status `gorm:"type:enum('created', 'in progress', 'done')"`
	Status Status `gorm:"type:check(status in ('created', 'in progress', 'done'))"`
	UserID uint
}
