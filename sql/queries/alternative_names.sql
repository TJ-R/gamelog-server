-- name: CreateAlternativeName :one
INSERT INTO alternative_names (
    id, created_at, updated_at, comment, game, name
) VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    $4
)
RETURNING *;

-- name: DeleteAllAlternativeNames :exec
DELETE FROM alternative_names;
