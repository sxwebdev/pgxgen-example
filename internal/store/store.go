package store

import "github.com/tkcrm/modules/db/postgres"

type IStore interface {
	// Postgres
	Querier
}

type Store struct {
	*Queries
}

func NewStore(db *postgres.PostgreSQL) IStore {
	return &Store{
		Queries: New(db.DB),
	}
}
