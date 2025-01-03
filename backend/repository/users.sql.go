// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package repository

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, provider, provider_id, username, email)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, provider, provider_id, username, email, avatar, created, updated
`

type CreateUserParams struct {
	ID         uuid.UUID
	Provider   string
	ProviderID string
	Username   string
	Email      string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.Provider,
		arg.ProviderID,
		arg.Username,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Provider,
		&i.ProviderID,
		&i.Username,
		&i.Email,
		&i.Avatar,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, provider, provider_id, username, email, avatar, created, updated FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Provider,
		&i.ProviderID,
		&i.Username,
		&i.Email,
		&i.Avatar,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, provider, provider_id, username, email, avatar, created, updated FROM users WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Provider,
		&i.ProviderID,
		&i.Username,
		&i.Email,
		&i.Avatar,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getUserByProviderId = `-- name: GetUserByProviderId :one
SELECT id, provider, provider_id, username, email, avatar, created, updated FROM users WHERE provider = $1 AND provider_id = $2
`

type GetUserByProviderIdParams struct {
	Provider   string
	ProviderID string
}

func (q *Queries) GetUserByProviderId(ctx context.Context, arg GetUserByProviderIdParams) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByProviderId, arg.Provider, arg.ProviderID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Provider,
		&i.ProviderID,
		&i.Username,
		&i.Email,
		&i.Avatar,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getUserExists = `-- name: GetUserExists :one
SELECT 1 FROM users WHERE id = $1
`

func (q *Queries) GetUserExists(ctx context.Context, id uuid.UUID) (int32, error) {
	row := q.db.QueryRowContext(ctx, getUserExists, id)
	var column_1 int32
	err := row.Scan(&column_1)
	return column_1, err
}
