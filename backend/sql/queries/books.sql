-- name: CreateBook :one
INSERT INTO books (name, author_id, created_at, updated_at)
VALUES ($1, $2, NOW(), NOW())
RETURNING *;

-- name: GetAllBooks :many
SELECT books.id, books.name, books.author_id, authors.name AS author_name
FROM books
JOIN authors
  ON books.author_id=authors.id
  AND books.deleted_at IS NULL;

-- name: GetBookById :one
SELECT books.id, books.name, books.author_id, authors.name AS author_name
FROM books
JOIN authors
  ON books.author_id=authors.id
  AND books.id = $1
  AND books.deleted_at IS NULL;

-- name: GetBookByName :many
SELECT books.id, books.name, books.author_id, authors.name AS author_name
FROM books
JOIN authors
  ON books.author_id=authors.id
  AND books.name LIKE $1
  AND books.deleted_at IS NULL;

-- name: DeleteBook :exec
UPDATE books
SET deleted_at=NOW()
WHERE id=$1;
