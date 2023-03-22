package books

import (
	"time"

	"github.com/google/uuid"
	"github.com/sxwebdev/pgxgen-example/internal/models"
)

type FindParams struct {
	Name     *string
	Genre    *models.BookType
	AuthorID *uuid.UUID
}

type CreateParams struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Genre       models.BookType `json:"genre"`
	ReleaseDate time.Time       `json:"release_date"`
	AuthorID    uuid.UUID       `json:"author_id"`
}

type UpdateParams struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Genre       models.BookType `json:"genre"`
	ReleaseDate time.Time       `json:"release_date"`
	ID          uuid.UUID       `json:"id"`
}
