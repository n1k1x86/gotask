package models

import (
	"time"
)

type TaskGroup struct {
	ID   uint
	Name string
}

type Task struct {
	ID          uint
	Title       string
	Description string
	GroupID     int
	DueDate     time.Time
}
