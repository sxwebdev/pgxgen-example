package handler

import (
	"context"

	"github.com/sxwebdev/pgxgen-example/internal/config"
	"github.com/sxwebdev/pgxgen-example/internal/models"
	"github.com/sxwebdev/pgxgen-example/internal/services/books"
	"github.com/sxwebdev/pgxgen-example/internal/store"
	"github.com/tkcrm/modules/logger"
)

type BooksServer struct {
	logger       logger.Logger
	config       *config.Config
	store        store.IStore
	booksService *books.Service
}

func newBooksServer(l logger.Logger, c *config.Config, st store.IStore, booksService *books.Service) *BooksServer {
	return &BooksServer{
		logger:       l,
		config:       c,
		store:        st,
		booksService: booksService,
	}
}

func (s *BooksServer) Create(ctx context.Context, req books.CreateParams) (*models.Book, error) {
	return s.booksService.Create(ctx, req)
}
