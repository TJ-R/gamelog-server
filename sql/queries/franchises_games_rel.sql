-- name: CreateFranchiseGameRel :one
INSERT INTO franchises_games_rel (
    franchise_id, games_id, created_at, updated_at
) VALUES (
    $1,
    $2,
    NOW(),
    NOW()
)
RETURNING *;

-- name: DeleteAllFranchiseGameRel :exec
DELETE FROM franchises_games_rel;
