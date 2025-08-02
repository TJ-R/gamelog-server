-- name: CreateTheme :one
INSERT INTO themes (
    id, created_at, updated_at, name, slug, url
) VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    $4
)
RETURNING *;

-- name: DeleteAllThemes :exec
DELETE FROM themes;
