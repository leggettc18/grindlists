-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users ORDER BY name;

-- name: CreateUser :one
INSERT INTO users (name, email, hashed_password, created_at, updated_at)
VALUES ($1, $2, $3, $4, $4)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET name = $2, email= $3, hashed_password = $4, updated_at = $5
WHERE id = $1
RETURNING *;

-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING *;

-- name: GetList :one
SELECT * FROM lists WHERE id = $1;

-- name: ListLists :many
SELECT * FROM lists ORDER BY name;

-- name: CreateList :one
INSERT INTO lists (name, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $3)
RETURNING *;

-- name: UpdateList :one
UPDATE lists
SET name = $2, user_id = $3, updated_at = $4
WHERE id = $1
RETURNING *;

-- name: DeleteList :one
DELETE FROM lists
WHERE id = $1
RETURNING *;

-- name: GetItem :one
SELECT * FROM items WHERE id = $1;

-- name: ListItems :many
SELECT * FROM items ORDER BY name;

-- name: CreateItem :one
INSERT INTO items (name, source, created_at, updated_at)
VALUES ($1, $2, $3, $3)
RETURNING *;

-- name: UpdateItem :one
UPDATE items
SET name = $2, source = $3, updated_at = $4
WHERE id = $1
RETURNING *;

-- name: DeleteItem :one
DELETE FROM items
WHERE id = $1
RETURNING *;

-- name: SetListItem :exec
INSERT INTO list_items (quantity, collected, list_id, item_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $5);

-- name: UpdateListItem :exec
UPDATE list_items
SET quantity = $2, collected = $3, updated_at = $4
WHERE id = $1;

-- name: UnsetListItem :exec
DELETE FROM list_items
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserLists :many
SELECT * FROM lists WHERE user_id = $1;