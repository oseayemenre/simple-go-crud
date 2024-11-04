-- name: CreateUser :one
INSERT INTO users(id, name, password, createdAt, updatedAt)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;