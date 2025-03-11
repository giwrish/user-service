-- +goose Up
CREATE TABLE users (
    id TEXT PRIMARY KEY,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT(timezone('utc', now())),
    updated_at TIMESTAMP DEFAULT(timezone('utc', now()))
);

-- +goose Down
DROP TABLE users;