-- name: CreateUser :one
INSERT INTO users (username, hashed_password, full_name, email, password_changed_at, created_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE username = $1;
