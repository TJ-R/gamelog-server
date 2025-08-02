-- name: CreatePlatform :one
INSERT INTO platforms (
    id, created_at, updated_at, abbreviation,
    generation, name, slug, summary, url
) VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING *;

-- name: DeleteAllPlatforms :exec
DELETE FROM platforms;
