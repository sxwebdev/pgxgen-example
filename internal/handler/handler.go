package handler

import (
	"github.com/sxwebdev/pgxgen-example/internal/config"
	"github.com/sxwebdev/pgxgen-example/internal/services/authors"
	"github.com/sxwebdev/pgxgen-example/internal/services/books"
	"github.com/sxwebdev/pgxgen-example/internal/store"
	"github.com/tkcrm/modules/logger"
)

type Handler struct {
	logger logger.Logger
	config *config.Config

	Authors *AuthorsServer
	Books   *BooksServer
}

func New(l logger.Logger, c *config.Config, st store.IStore) *Handler {
	return &Handler{
		logger:  l,
		config:  c,
		Authors: newAuthorsServer(l, c, st, authors.New(l, c, st)),
		Books:   newBooksServer(l, c, st, books.New(l, c, st)),
	}
}
