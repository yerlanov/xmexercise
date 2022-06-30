package storage

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/yerlanov/xmexercise/internal/company"
)

type Store interface {
	company.Storage
}

type SQLStore struct {
	*Queries
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) Store {
	return &SQLStore{
		Queries: New(db),
		db:      db,
	}
}
