-- name: CreateGame :one
INSERT INTO games (
    id, created_at, updated_at, first_release_date, name, parent_game,
    slug, storyline, summary, url, version_parent, version_title
) VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    NULL,
    $4,
    $5,
    $6,
    $7,
    NULL,
    $8
)
RETURNING *;

-- name: GetGameById :one
SELECT *
FROM games
WHERE id = $1;

-- name: UpdateParentAndVersionParent :exec
UPDATE games
SET parent_game = $2, version_parent = $3
WHERE id = $1
RETURNING *;

-- name: DeleteAllGames :exec
DELETE FROM games;
