package pgstore

import (
	"context"
	"rocketseat-go/internal/store"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PGTaskStore struct {
	Queries *Queries
	Pool    *pgxpool.Pool
}

func NewPGTaskStore(pool *pgxpool.Pool) PGTaskStore {
	return PGTaskStore{
		Queries: New(pool),
		Pool:    pool,
	}
}

func (pgs *PGTaskStore) CreateTask(
	ctx context.Context,
	title, description string,
	priority int32,
) (store.Task, error) {
	task, err := pgs.Queries.CreateTask(ctx, CreateTaskParams{
		Title:       title,
		Description: description,
		Priority:    priority,
	})
	if err != nil {
		return store.Task{}, err
	}
	return store.Task{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Priority:    task.Priority,
		CreatedAt:   task.CreatedAt.Time,
		UpdatedAt:   task.UpdatedAt.Time,
	}, nil
}

func (pgs *PGTaskStore) GetTaskByID(ctx context.Context, id int) (Task, error) {
	panic("Implement GetTaskByID method")
}

func (pgs *PGTaskStore) UpdateTask(ctx context.Context, task Task) error {
	panic("Implement UpdateTask method")
}

func (pgs *PGTaskStore) DeleteTask(ctx context.Context, id int) error {
	panic("Implement DeleteTask method")
}

func (pgs *PGTaskStore) ListTasks(ctx context.Context) ([]Task, error) {
	panic("Implement ListTasks method")
}
