package store

import "time"

type Task struct {
	Id          int32              `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Priority    int32              `json:"priority"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type TaskStore interface {
	CreateTask(title string, description string, priority int32) (Task, error)
	GetTaskByID(id int32) (Task, error)
	ListTasks() ([]Task, error)
	UpdateTask(
		id int32,
		title string,
		description string,
		priority int32,
	) (Task, error)
	DeleteTask(id int32) error
}
