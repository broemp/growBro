package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type DBStore interface{ Querier }

type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

func NewStore(connPool *pgxpool.Pool) DBStore {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
