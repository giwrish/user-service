-- +goose Up
ALTER TABLE users
RENAME COLUMN id to username;

-- +goose Down
ALTER TABLE users
RENAME COLUMN username to id;