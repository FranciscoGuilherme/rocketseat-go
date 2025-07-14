package pgstore

import "github.com/jackc/pgx/v5/pgxpool"

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

func (store *PGTaskStore) CreateTask(
	title,
	description string,
	priority int,
) error {
	panic("Implement CreateTask method")
}

func (store *PGTaskStore) GetTaskByID(id int) (Task, error) {
	panic("Implement GetTaskByID method")
}

func (store *PGTaskStore) UpdateTask(task Task) error {
	panic("Implement UpdateTask method")
}

func (store *PGTaskStore) DeleteTask(id int) error {
	panic("Implement DeleteTask method")
}

func (store *PGTaskStore) ListTasks() ([]Task, error) {
	panic("Implement ListTasks method")
}
