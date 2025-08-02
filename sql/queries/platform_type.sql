-- name: CreatePlatformType :one
INSERT INTO platform_type (
    id, created_at, updated_at, name
) VALUES (
    $1,
    NOW(),
    NOW(),
    $2
)
RETURNING *;

-- name: DeleteAllPlatformType :exec
DELETE FROM platform_type;
