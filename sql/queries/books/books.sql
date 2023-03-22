-- name: DeleteByAuthorID :exec
DELETE FROM books WHERE "author_id"=$1;
