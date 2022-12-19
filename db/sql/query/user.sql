-- name: CreateUser :one
INSERT INTO users (username, password, full_name, email)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE username = $1 LIMIT 1;

-- name: ListUser :many
SELECT * FROM users
ORDER BY username
LIMIT $1
OFFSET $2;

-- name: UpdateUserPassword :one
UPDATE users
SET password = $2
WHERE username = $1 RETURNING *;

-- name: DeleteUser :exec
DELETE
FROM users
WHERE username = $1;