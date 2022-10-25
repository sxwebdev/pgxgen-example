package books

import (
	"context"

	"github.com/google/uuid"
	"github.com/sxwebdev/pgxgen-example/internal/store"
)

func (s *Service) Create(ctx context.Context, req store.CreateBookParams) (*store.Book, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return s.store.CreateBook(ctx, req)
}

func (s *Service) Update(ctx context.Context, req store.UpdateBookParams) (*store.Book, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return s.store.UpdateBook(ctx, req)
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return store.ErrEmptyBookID
	}

	return s.store.DeleteBook(ctx, id)
}

func (s *Service) Total(ctx context.Context) (int64, error) {
	return s.store.TotalBooks(ctx)
}

func (s *Service) Get(ctx context.Context, id uuid.UUID) (*store.Book, error) {
	if id == uuid.Nil {
		return nil, store.ErrEmptyBookID
	}

	return s.store.GetBookByID(ctx, id)
}

func (s *Service) FindByBooksByAuthorID(ctx context.Context, authorId uuid.UUID) ([]*store.Book, error) {
	if authorId == uuid.Nil {
		return nil, store.ErrEmptyAuthorID
	}

	return s.store.FindBooksByAuthorID(ctx, authorId)
}
