-- name: CreateGameType :one
INSERT INTO game_types (
    id, created_at, updated_at, type
) VALUES (
    $1,
    NOW(),
    NOW(),
    $2
)
RETURNING *;

-- name: DeleteAllGameTypes :exec
DELETE FROM game_types;
