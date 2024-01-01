package models

import (
	"backend/gen/todo"
	"time"
)

type Task struct {
	Id          string     `gorm:"primaryKey"`
	Title       string     `gorm:"not null"`
	Description string     `gorm:"not null;default ''"`
	Completed   bool       `gorm:"not null;default false"`
	CompletedAt *time.Time `gorm:"default null"`
	CompletedBy *string    `gorm:"default null"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	CreatedBy   string     `gorm:"not null"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime"`
}

func (t *Task) ToResponse() (*task.Task) {

}