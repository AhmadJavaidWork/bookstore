-- name: CreateAuthor :one
INSERT INTO authors (name, created_at, updated_at)
VALUES ($1, NOW(), NOW())
RETURNING *;

-- name: GetAllAuthors :many
SELECT id, name
FROM authors
WHERE deleted_at IS NULL;

-- name: GetAuthorById :one
SELECT id, name
FROM authors
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetAuthorByName :many
SELECT id, name
FROM authors
WHERE name LIKE $1 AND deleted_at IS NULL;

-- name: DeleteAuthor :exec
UPDATE authors
SET deleted_at=NOW()
WHERE id=$1;
