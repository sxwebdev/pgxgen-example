package repos

import (
	"github.com/sxwebdev/pgxgen-example/internal/store/repos/repo_authors"
	"github.com/sxwebdev/pgxgen-example/internal/store/repos/repo_books"
	"github.com/tkcrm/modules/db/postgres"
)

type IRepos interface {
	Authors(opts ...Option) repo_authors.Querier
	Books(opts ...Option) repo_books.ICustomQuerier
}

type repos struct {
	authors *repo_authors.Queries
	books   *repo_books.CustomQuerier
}

func New(psql *postgres.PostgreSQL) IRepos {
	return &repos{
		authors: repo_authors.New(psql.DB),
		books:   repo_books.NewCustom(psql.DB),
	}
}

func (s *repos) Authors(opts ...Option) repo_authors.Querier {
	options := getOptions(opts...)

	if options.Tx != nil {
		return s.authors.WithTx(options.Tx)
	}

	return s.authors
}

func (s *repos) Books(opts ...Option) repo_books.ICustomQuerier {
	options := getOptions(opts...)

	if options.Tx != nil {
		return s.books.WithTx(options.Tx)
	}

	return s.books
}

var _ IRepos = (*repos)(nil)
