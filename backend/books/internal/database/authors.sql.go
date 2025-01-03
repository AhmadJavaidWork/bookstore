// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: authors.sql

package database

import (
	"context"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (name, created_at, updated_at)
VALUES ($1, NOW(), NOW())
RETURNING id, name, deleted_at, created_at, updated_at
`

func (q *Queries) CreateAuthor(ctx context.Context, name string) (Author, error) {
	row := q.db.QueryRowContext(ctx, createAuthor, name)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.DeletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAuthor = `-- name: DeleteAuthor :exec
UPDATE authors
SET deleted_at=NOW()
WHERE id=$1
`

func (q *Queries) DeleteAuthor(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteAuthor, id)
	return err
}

const getAllAuthors = `-- name: GetAllAuthors :many
SELECT id, name
FROM authors
WHERE deleted_at IS NULL
`

type GetAllAuthorsRow struct {
	ID   int32
	Name string
}

func (q *Queries) GetAllAuthors(ctx context.Context) ([]GetAllAuthorsRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllAuthorsRow
	for rows.Next() {
		var i GetAllAuthorsRow
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAuthorById = `-- name: GetAuthorById :one
SELECT id, name
FROM authors
WHERE id = $1 AND deleted_at IS NULL
`

type GetAuthorByIdRow struct {
	ID   int32
	Name string
}

func (q *Queries) GetAuthorById(ctx context.Context, id int32) (GetAuthorByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getAuthorById, id)
	var i GetAuthorByIdRow
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getAuthorByName = `-- name: GetAuthorByName :many
SELECT id, name
FROM authors
WHERE name LIKE $1 AND deleted_at IS NULL
`

type GetAuthorByNameRow struct {
	ID   int32
	Name string
}

func (q *Queries) GetAuthorByName(ctx context.Context, name string) ([]GetAuthorByNameRow, error) {
	rows, err := q.db.QueryContext(ctx, getAuthorByName, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAuthorByNameRow
	for rows.Next() {
		var i GetAuthorByNameRow
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
