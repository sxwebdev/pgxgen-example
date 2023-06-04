package authors

import (
	"context"

	"github.com/google/uuid"
	"github.com/sxwebdev/pgxgen-example/internal/models"
	"github.com/sxwebdev/pgxgen-example/internal/store/repos/repo_authors"
	"github.com/sxwebdev/pgxgen-example/internal/store/storecmn"
	"github.com/tkcrm/modules/pkg/utils"
)

func (s *Service) Create(ctx context.Context, req CreateParams) (*models.Author, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	var params repo_authors.CreateParams
	if err := utils.JsonToStruct(req, &params); err != nil {
		return nil, err
	}

	return s.store.Authors().Create(ctx, params)
}

func (s *Service) Update(ctx context.Context, req UpdateParams) (*models.Author, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	var params repo_authors.UpdateParams
	if err := utils.JsonToStruct(req, &params); err != nil {
		return nil, err
	}

	return s.store.Authors().Update(ctx, params)
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return storecmn.ErrEmptyAuthorID
	}

	// Remove all books
	if err := s.store.Books().DeleteByAuthorID(ctx, id); err != nil {
		return err
	}

	return s.store.Authors().Delete(ctx, id)
}

func (s *Service) Total(ctx context.Context) (int64, error) {
	return s.store.Authors().Total(ctx)
}

func (s *Service) Get(ctx context.Context, id uuid.UUID) (*models.Author, error) {
	if id == uuid.Nil {
		return nil, storecmn.ErrEmptyAuthorID
	}

	return s.store.Authors().GetAuthorByID(ctx, id)
}

func (s *Service) Find(ctx context.Context) ([]*models.Author, error) {
	return s.store.Authors().Find(ctx)
}
