package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string
	Description string
	GroupID     int
	Group       TaskGroup `gorm:"costraint: OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type TaskGroup struct {
	gorm.Model
	Name string
}
