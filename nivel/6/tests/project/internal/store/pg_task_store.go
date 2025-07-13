package store

import (
	"rocketseat-go/internal/store/pgstore"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PGTaskStore struct {
	Queries *pgstore.Queries
	Pool    *pgxpool.Pool
}

func NewPGTaskStore(pool *pgxpool.Pool) PGTaskStore {
	return PGTaskStore{
		Queries: pgstore.New(pool),
		Pool:    pool,
	}
}
