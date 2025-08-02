-- name: CreateGenre :one
INSERT INTO genres (
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

-- name: DeleteAllGenres :exec
DELETE FROM genres;
