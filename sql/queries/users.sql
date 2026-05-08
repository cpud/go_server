-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password, is_chirpy_red)
VALUES (
    gen_random_uuid(), Now(), Now(), $1, $2, FALSE
)
RETURNING *;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: UpdateUserByID :one
UPDATE users
SET email = $2, hashed_password=$3, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: ChirpyRed :one
UPDATE users
SET is_chirpy_red = true
WHERE id = $1
RETURNING *;