-- name: CreateChirp :one
INSERT INTO Chirps (id, created_at, updated_at, body, user_id)
VALUES (
    gen_random_uuid(), Now(), Now(), $1, $2
)
RETURNING *;

-- name: GetChirps :many
SELECT id, created_at, updated_at, body, user_id FROM Chirps
ORDER BY created_at ASC;

-- name: GetChirpById :one
SELECT id, created_at, updated_at, body, user_id FROM Chirps
WHERE id = $1;

-- name: DeleteChirpById :exec
DELETE FROM Chirps
WHERE id = $1;

-- name: GetChirpsByUserId :many
SELECT * FROM Chirps
WHERE user_id = $1;