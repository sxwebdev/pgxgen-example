-- name: Create :one
INSERT INTO authors (first_name, last_name, email, phone, is_active, notifications, created_at)
	VALUES ($1, $2, $3, $4, $5, $6, now())
	RETURNING *;

-- name: Delete :exec
DELETE FROM authors WHERE id=$1;

-- name: Find :many
SELECT * FROM authors ORDER BY created_at DESC;

-- name: GetAuthorByID :one
SELECT * FROM authors WHERE id=$1 LIMIT 1;

-- name: Total :one
SELECT count(*) as total FROM authors;

-- name: Update :one
UPDATE authors
	SET first_name=$1, last_name=$2, email=$3, phone=$4, is_active=$5, notifications=$6, 
		updated_at=now()
	WHERE id=$7
	RETURNING *;

