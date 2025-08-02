-- name: CreatePlatformFamily :one
INSERT INTO platform_family (
    id, created_at, updated_at, name, slug
) VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3
)
RETURNING *;

-- name: DeleteAllPlatformFamilies :exec
DELETE FROM platform_family;
