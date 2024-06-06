-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT * FROM users WHERE id = ?;

-- name: CreateUser :exec
INSERT INTO users (id, username, email, password) VALUES (?, ?, ?, ?);

-- name: UpdateUserEmail :exec
UPDATE users SET email = ? WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;