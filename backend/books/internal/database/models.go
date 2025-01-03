// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"time"
)

type Author struct {
	ID        int32
	Name      string
	DeletedAt sql.NullTime
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Book struct {
	ID        int32
	Name      string
	AuthorID  int32
	DeletedAt sql.NullTime
	CreatedAt time.Time
	UpdatedAt time.Time
}
