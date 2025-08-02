-- name: CreateGameAlternativeNameRel :one
INSERT INTO games_alternative_names_rel (
    game_id, alternative_name_id, created_at, updated_at
) VALUES (
    $1,
    $2,
    NOW(),
    NOW()
)
RETURNING *;

-- name: DeleteAllGameAlternativeNameRel :exec
DELETE FROM games_alternative_names_rel;
