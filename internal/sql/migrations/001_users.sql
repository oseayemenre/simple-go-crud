-- +goose Up
CREATE TABLE users(
    id UUID PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE users;