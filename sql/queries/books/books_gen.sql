-- name: Create :one
INSERT INTO books (name, description, genre, release_date, author_id, created_at)
	VALUES ($1, $2, $3, $4, $5, now())
	RETURNING *;

-- name: Delete :exec
DELETE FROM books WHERE id=$1;

-- name: GetBookByID :one
SELECT * FROM books WHERE id=$1 LIMIT 1;

-- name: Total :one
SELECT count(*) as total FROM books;

-- name: Update :one
UPDATE books
	SET name=$1, description=$2, genre=$3, release_date=$4, updated_at=now()
	WHERE id=$5
	RETURNING *;

