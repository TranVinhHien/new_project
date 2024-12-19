// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  username, password,full_name,create_at
) VALUES (
  $1, $2, $3, $4
)RETURNING username, password, full_name, is_active, create_at
`

type CreateUserParams struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
	FullName string    `json:"full_name"`
	CreateAt time.Time `json:"create_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (Users, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.Password,
		arg.FullName,
		arg.CreateAt,
	)
	var i Users
	err := row.Scan(
		&i.Username,
		&i.Password,
		&i.FullName,
		&i.IsActive,
		&i.CreateAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT username, password, full_name, is_active, create_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (Users, error) {
	row := q.db.QueryRow(ctx, getUser, username)
	var i Users
	err := row.Scan(
		&i.Username,
		&i.Password,
		&i.FullName,
		&i.IsActive,
		&i.CreateAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET password = COALESCE($1, password),
full_name = COALESCE($2, full_name),
is_active = COALESCE($3, is_active)
WHERE username = $4 
RETURNING username, password, full_name, is_active, create_at
`

type UpdateUserParams struct {
	Password pgtype.Text `json:"password"`
	FullName pgtype.Text `json:"full_name"`
	IsActive pgtype.Bool `json:"is_active"`
	Username string      `json:"username"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (Users, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.Password,
		arg.FullName,
		arg.IsActive,
		arg.Username,
	)
	var i Users
	err := row.Scan(
		&i.Username,
		&i.Password,
		&i.FullName,
		&i.IsActive,
		&i.CreateAt,
	)
	return i, err
}
