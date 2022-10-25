package authors

import (
	"context"

	"github.com/google/uuid"
	"github.com/sxwebdev/pgxgen-example/internal/store"
)

func (s *Service) Create(ctx context.Context, req store.CreateAuthorParams) (*store.Author, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return s.store.CreateAuthor(ctx, req)
}

func (s *Service) Update(ctx context.Context, req store.UpdateAuthorParams) (*store.Author, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return s.store.UpdateAuthor(ctx, req)
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return store.ErrEmptyAuthorID
	}

	// Remove all books
	if err := s.store.DeleteBooksByAuthorID(ctx, id); err != nil {
		return err
	}

	return s.store.DeleteAuthor(ctx, id)
}

func (s *Service) Total(ctx context.Context) (int64, error) {
	return s.store.TotalAuthors(ctx)
}

func (s *Service) Get(ctx context.Context, id uuid.UUID) (*store.Author, error) {
	if id == uuid.Nil {
		return nil, store.ErrEmptyAuthorID
	}

	return s.store.GetAuthorByID(ctx, id)
}

func (s *Service) Find(ctx context.Context) ([]*store.Author, error) {
	return s.store.FindAuthors(ctx)
}
