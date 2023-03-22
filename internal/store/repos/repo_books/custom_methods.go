package repo_books

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/huandu/go-sqlbuilder"
	"github.com/sxwebdev/pgxgen-example/internal/models"
)

func (s *CustomQuerier) findBuilder(params FindParams, columns ...string) *sqlbuilder.SelectBuilder {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder().
		Select(columns...).
		From(TableNameBooks.String())

	if params.Name != nil {
		sb.Where(sb.Equal("name", *params.Name))
	}

	if params.Genre != nil {
		sb.Where(sb.Equal("genre", *params.Genre))
	}

	if params.AuthorID != nil {
		sb.Where(sb.Equal("author_id", params.AuthorID.String()))
	}

	// set order by
	sb.OrderBy("created_at")

	// set order direction
	sb.Desc()

	return sb
}

func (s *CustomQuerier) Find(ctx context.Context, params FindParams) ([]*models.Book, error) {
	// init builder
	sb := s.findBuilder(params, "*")

	// execute query
	items := []*models.Book{}
	sql, args := sb.Build()
	if err := pgxscan.Get(ctx, s.psql, items, sql, args...); err != nil {
		return nil, err
	}

	return items, nil
}
