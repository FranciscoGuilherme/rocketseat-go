package store

import "github.com/jackc/pgx/v5/pgtype"

type Task struct {
	ID          int32              `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Priority    int32              `json:"priority"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
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
