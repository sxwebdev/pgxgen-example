package repo_books

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/sxwebdev/pgxgen-example/internal/models"
)

type ICustomQuerier interface {
	Querier
	Find(ctx context.Context, params FindParams) ([]*models.Book, error)
}

type CustomQuerier struct {
	*Queries
	psql DBTX
}

func NewCustom(psql DBTX) *CustomQuerier {
	return &CustomQuerier{
		Queries: New(psql),
		psql:    psql,
	}
}

func (q *CustomQuerier) WithTx(tx pgx.Tx) *CustomQuerier {
	return &CustomQuerier{
		Queries: New(tx),
		psql:    tx,
	}
}

var _ ICustomQuerier = (*CustomQuerier)(nil)
