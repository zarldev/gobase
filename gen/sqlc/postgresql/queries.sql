-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: CreateUser :exec
INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4);

-- name: UpdateUserEmail :exec
UPDATE users SET email = $2 WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;