// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: authors_gen.sql

package repo_authors

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	. "github.com/sxwebdev/pgxgen-example/internal/models"
)

const create = `-- name: Create :one
INSERT INTO authors (first_name, last_name, email, phone, is_active, notifications, created_at)
	VALUES ($1, $2, $3, $4, $5, $6, now())
	RETURNING id, first_name, last_name, email, phone, is_active, notifications, created_at, updated_at
`

type CreateParams struct {
	FirstName     string              `db:"first_name" json:"first_name"`
	LastName      string              `db:"last_name" json:"last_name"`
	Email         string              `db:"email" json:"email"`
	Phone         pgtype.Text         `db:"phone" json:"phone"`
	IsActive      bool                `db:"is_active" json:"is_active"`
	Notifications AuthorNotifications `db:"notifications" json:"notifications"`
}

func (q *Queries) Create(ctx context.Context, arg CreateParams) (*Author, error) {
	row := q.db.QueryRow(ctx, create,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Phone,
		arg.IsActive,
		arg.Notifications,
	)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
		&i.IsActive,
		&i.Notifications,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const delete = `-- name: Delete :exec
DELETE FROM authors WHERE id=$1
`

func (q *Queries) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, delete, id)
	return err
}

const find = `-- name: Find :many
SELECT id, first_name, last_name, email, phone, is_active, notifications, created_at, updated_at FROM authors ORDER BY created_at DESC
`

func (q *Queries) Find(ctx context.Context) ([]*Author, error) {
	rows, err := q.db.Query(ctx, find)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Author{}
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Phone,
			&i.IsActive,
			&i.Notifications,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAuthorByID = `-- name: GetAuthorByID :one
SELECT id, first_name, last_name, email, phone, is_active, notifications, created_at, updated_at FROM authors WHERE id=$1 LIMIT 1
`

func (q *Queries) GetAuthorByID(ctx context.Context, id uuid.UUID) (*Author, error) {
	row := q.db.QueryRow(ctx, getAuthorByID, id)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
		&i.IsActive,
		&i.Notifications,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const total = `-- name: Total :one
SELECT count(*) as total FROM authors
`

func (q *Queries) Total(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, total)
	var total int64
	err := row.Scan(&total)
	return total, err
}

const update = `-- name: Update :one
UPDATE authors
	SET first_name=$1, last_name=$2, email=$3, phone=$4, is_active=$5, notifications=$6, 
		updated_at=now()
	WHERE id=$7
	RETURNING id, first_name, last_name, email, phone, is_active, notifications, created_at, updated_at
`

type UpdateParams struct {
	FirstName     string              `db:"first_name" json:"first_name"`
	LastName      string              `db:"last_name" json:"last_name"`
	Email         string              `db:"email" json:"email"`
	Phone         pgtype.Text         `db:"phone" json:"phone"`
	IsActive      bool                `db:"is_active" json:"is_active"`
	Notifications AuthorNotifications `db:"notifications" json:"notifications"`
	ID            uuid.UUID           `db:"id" json:"id"`
}

func (q *Queries) Update(ctx context.Context, arg UpdateParams) (*Author, error) {
	row := q.db.QueryRow(ctx, update,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Phone,
		arg.IsActive,
		arg.Notifications,
		arg.ID,
	)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
		&i.IsActive,
		&i.Notifications,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}
