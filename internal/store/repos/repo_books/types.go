package repo_books

import (
	"github.com/google/uuid"
	"github.com/sxwebdev/pgxgen-example/internal/models"
)

type FindParams struct {
	Name     *string
	Genre    *models.BookType
	AuthorID *uuid.UUID
}
