-- name: CreateUser :one
INSERT INTO users (username, password, created_at, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING username, created_at, updated_at;

-- name: GetUser :one
SELECT username, created_at, updated_at FROM users 
WHERE username = $1;

-- name: DeleteUser :one
DELETE FROM users
WHERE username = $1
RETURNING username;