package books

import (
	"context"

	"github.com/google/uuid"
	"github.com/sxwebdev/pgxgen-example/internal/models"
	"github.com/sxwebdev/pgxgen-example/internal/store/repos/repo_books"
	"github.com/sxwebdev/pgxgen-example/internal/store/storecmn"
)

func (s *Service) Create(ctx context.Context, req CreateParams) (*models.Book, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return s.store.Books().Create(ctx, repo_books.CreateParams(req))
}

func (s *Service) Update(ctx context.Context, req UpdateParams) (*models.Book, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return s.store.Books().Update(ctx, repo_books.UpdateParams(req))
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return storecmn.ErrEmptyBookID
	}

	return s.store.Books().Delete(ctx, id)
}

func (s *Service) Total(ctx context.Context) (int64, error) {
	return s.store.Books().Total(ctx)
}

func (s *Service) Get(ctx context.Context, id uuid.UUID) (*models.Book, error) {
	if id == uuid.Nil {
		return nil, storecmn.ErrEmptyBookID
	}

	return s.store.Books().GetBookByID(ctx, id)
}

func (s *Service) Find(ctx context.Context, params FindParams) ([]*models.Book, error) {
	return s.store.Books().Find(ctx, repo_books.FindParams(params))
}
