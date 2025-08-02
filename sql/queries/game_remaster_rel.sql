-- name: CreateGameRemasterRel :one
INSERT INTO game_remaster_rel (
    game_id, remaster_id, created_at, updated_at
) VALUES (
    $1,
    $2,
    NOW(),
    NOW()
)
RETURNING *;

-- name: DeleteAllGameRemasterRel :exec
DELETE FROM game_remaster_rel;
