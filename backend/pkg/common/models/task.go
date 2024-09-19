package models

type TaskGroup struct {
	ID    uint
	Name  string
	Tasks []Task
}

type Task struct {
	ID          uint
	Title       string
	Description string
	GroupID     int
}
