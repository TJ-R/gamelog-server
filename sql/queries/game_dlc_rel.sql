-- name: CreateGameDlcRel :one
INSERT INTO game_dlc_rel (
    game_id, dlc_id, created_at, updated_at
) VALUES (
    $1,
    $2,
    NOW(),
    NOW()
)
RETURNING *;

-- name: DeleteAllGameDlcRel :exec
DELETE FROM game_dlc_rel;
