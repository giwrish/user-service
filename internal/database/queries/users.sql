-- name: UserExists :one
SELECT EXISTS(SELECT 1 FROM users WHERE username = $1);

-- name: GetUser :one
SELECT username, created_at, updated_at FROM users 
WHERE username = $1;

-- name: CreateUser :one
INSERT INTO users (username, password, created_at, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING username, created_at, updated_at;

-- name: UpdateUserPassword :one
UPDATE users
SET password = $1, updated_at = $2
WHERE username = $3
RETURNING username;

-- name: DeleteUser :one
DELETE FROM users
WHERE username = $1
RETURNING username;