-- name: CreateUser :one
INSERT INTO users (id, provider, provider_id, username, email, created, updated)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: CreateUserWithPassword :one
INSERT INTO users (id, provider, provider_id, username, email, password_hash, created, updated)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByProviderId :one
SELECT * FROM users WHERE provider = $1 AND provider_id = $2;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;
