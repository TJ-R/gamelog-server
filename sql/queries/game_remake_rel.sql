-- name: CreateGameRemakeRel :one
INSERT INTO game_remake_rel (
    game_id, remake_id, created_at, updated_at
) VALUES (
    $1, 
    $2,
    NOW(),
    NOW()
)
RETURNING *;

-- name: DeleteAllGameRemakeRel :exec
DELETE FROM game_remake_rel;
