-- name: CreatePlatformPlatformFamilyRel :one
INSERT INTO platforms_platform_family_rel (
    platform_id, platform_family_id, created_at, updated_at
) VALUES (
    $1,
    $2,
    NOW(),
    NOW()
)
RETURNING *;

-- name: DeleteAllPlatformPlatformFamilyRel :exec
DELETE FROM platforms_platform_family_rel;
