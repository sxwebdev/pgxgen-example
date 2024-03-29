// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package repo_authors

import (
	"context"

	"github.com/google/uuid"
	. "github.com/sxwebdev/pgxgen-example/internal/models"
)

type Querier interface {
	Create(ctx context.Context, arg CreateParams) (*Author, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Find(ctx context.Context) ([]*Author, error)
	GetAuthorByID(ctx context.Context, id uuid.UUID) (*Author, error)
	Total(ctx context.Context) (int64, error)
	Update(ctx context.Context, arg UpdateParams) (*Author, error)
}

var _ Querier = (*Queries)(nil)
