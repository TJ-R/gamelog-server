-- name: CreateGameExpansionRel :one
INSERT INTO game_expansion_rel (
    game_id, expansion_id, created_at, updated_at
) VALUES (
    $1,
    $2,
    NOW(),
    NOW()
)
RETURNING *;

-- name: DeleteAllGameExpansionRel :exec
DELETE FROM game_expansion_rel;
