package model

import "gorm.io/gorm"

type Status string

const (
	Created    Status = "created"
	InProgress Status = "in progress"
	Done       Status = "done"
)

func (s Status) Values() []string {
	return []string{
		"created",
		"in progress",
		"done",
	}
}

type Todo struct {
	gorm.Model
	// Title string `json:"title"`
	Body string `json:"body"`
	// Status Status `sql:"type:ENUM('created', 'in progress', 'done')"`
	// Progress Status `gorm:"type:enum('created', 'in progress', 'done')"`
	Status Status `json:"status" gorm:"type:check(status in ('created', 'in progress', 'done'))"`
	UserID uint
}
