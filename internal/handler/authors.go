package handler

import (
	"context"

	"github.com/sxwebdev/pgxgen-example/internal/config"
	"github.com/sxwebdev/pgxgen-example/internal/models"
	"github.com/sxwebdev/pgxgen-example/internal/services/authors"
	"github.com/sxwebdev/pgxgen-example/internal/store"
	"github.com/tkcrm/modules/pkg/logger"
)

type AuthorsServer struct {
	logger         logger.Logger
	config         *config.Config
	store          store.IStore
	authorsService *authors.Service
}

func newAuthorsServer(l logger.Logger, c *config.Config, st store.IStore, authorsService *authors.Service) *AuthorsServer {
	return &AuthorsServer{
		logger:         l,
		config:         c,
		store:          st,
		authorsService: authorsService,
	}
}

func (s *AuthorsServer) Create(ctx context.Context, req authors.CreateParams) (*models.Author, error) {
	return s.authorsService.Create(ctx, req)
}
