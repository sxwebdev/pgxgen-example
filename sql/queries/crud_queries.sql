-- name: CreateAuthor :one
INSERT INTO authors ("first_name", "last_name", "email", "phone", "is_active", "notifications", "created_at")
	VALUES ($1, $2, $3, $4, $5, $6, now())
	RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM authors WHERE "id"=$1;

-- name: FindAuthors :many
SELECT * FROM authors ORDER BY "created_at" DESC;

-- name: GetAuthorByID :one
SELECT * FROM authors WHERE "id"=$1;

-- name: TotalAuthors :one
SELECT count(*) as total FROM authors;

-- name: UpdateAuthor :one
UPDATE authors
	SET "first_name"=$1, "last_name"=$2, "email"=$3, "phone"=$4, "is_active"=$5, "notifications"=$6, 
		"updated_at"=now()
	WHERE "id"=$7
	RETURNING *;

-- name: CreateBook :one
INSERT INTO books ("name", "description", "genre", "release_date", "author_id", "created_at")
	VALUES ($1, $2, $3, $4, $5, now())
	RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books WHERE "id"=$1;

-- name: FindBooksByAuthorID :many
SELECT * FROM books WHERE "author_id"=$1 ORDER BY "created_at" DESC;

-- name: GetBookByID :one
SELECT * FROM books WHERE "id"=$1;

-- name: TotalBooks :one
SELECT count(*) as total FROM books;

-- name: UpdateBook :one
UPDATE books
	SET "name"=$1, "description"=$2, "genre"=$3, "release_date"=$4, "updated_at"=now()
	WHERE "id"=$5
	RETURNING *;

