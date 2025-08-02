-- name: CreateFranchise :one
INSERT INTO franchises (
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

-- name: DeleteAllFranchises :exec
DELETE FROM franchises;
