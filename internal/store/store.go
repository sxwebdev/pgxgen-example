package store

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sxwebdev/pgxgen-example/internal/store/repos"
	"github.com/tkcrm/modules/db/postgres"
)

type IStore interface {
	repos.IRepos
	DBConn() *pgxpool.Pool
}

type store struct {
	repos.IRepos
	psql *postgres.PostgreSQL
}

func New(psql *postgres.PostgreSQL) IStore {
	return &store{
		IRepos: repos.New(psql),
		psql:   psql,
	}
}

func (s *store) DBConn() *pgxpool.Pool {
	return s.psql.DB
}
