// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: books_gen.sql

package repo_books

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	. "github.com/sxwebdev/pgxgen-example/internal/models"
)

const create = `-- name: Create :one
INSERT INTO books (name, description, genre, release_date, author_id, created_at)
	VALUES ($1, $2, $3, $4, $5, now())
	RETURNING id, name, description, genre, release_date, author_id, created_at, updated_at
`

type CreateParams struct {
	Name        string           `db:"name" json:"name"`
	Description string           `db:"description" json:"description"`
	Genre       BookType         `db:"genre" json:"genre"`
	ReleaseDate pgtype.Timestamp `db:"release_date" json:"release_date"`
	AuthorID    uuid.UUID        `db:"author_id" json:"author_id"`
}

func (q *Queries) Create(ctx context.Context, arg CreateParams) (*Book, error) {
	row := q.db.QueryRow(ctx, create,
		arg.Name,
		arg.Description,
		arg.Genre,
		arg.ReleaseDate,
		arg.AuthorID,
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Genre,
		&i.ReleaseDate,
		&i.AuthorID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const delete = `-- name: Delete :exec
DELETE FROM books WHERE id=$1
`

func (q *Queries) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, delete, id)
	return err
}

const getBookByID = `-- name: GetBookByID :one
SELECT id, name, description, genre, release_date, author_id, created_at, updated_at FROM books WHERE id=$1 LIMIT 1
`

func (q *Queries) GetBookByID(ctx context.Context, id uuid.UUID) (*Book, error) {
	row := q.db.QueryRow(ctx, getBookByID, id)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Genre,
		&i.ReleaseDate,
		&i.AuthorID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const total = `-- name: Total :one
SELECT count(*) as total FROM books
`

func (q *Queries) Total(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, total)
	var total int64
	err := row.Scan(&total)
	return total, err
}

const update = `-- name: Update :one
UPDATE books
	SET name=$1, description=$2, genre=$3, release_date=$4, updated_at=now()
	WHERE id=$5
	RETURNING id, name, description, genre, release_date, author_id, created_at, updated_at
`

type UpdateParams struct {
	Name        string           `db:"name" json:"name"`
	Description string           `db:"description" json:"description"`
	Genre       BookType         `db:"genre" json:"genre"`
	ReleaseDate pgtype.Timestamp `db:"release_date" json:"release_date"`
	ID          uuid.UUID        `db:"id" json:"id"`
}

func (q *Queries) Update(ctx context.Context, arg UpdateParams) (*Book, error) {
	row := q.db.QueryRow(ctx, update,
		arg.Name,
		arg.Description,
		arg.Genre,
		arg.ReleaseDate,
		arg.ID,
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Genre,
		&i.ReleaseDate,
		&i.AuthorID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}
