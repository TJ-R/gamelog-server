-- name: CreateGameGenreRel :one
INSERT INTO game_genre_rel (
    game_id, genre_id, created_at, updated_at
) VALUES (
    $1,
    $2,
    NOW(),
    NOW()
)
RETURNING *;

-- name: DeleteAllGameGenreRel :exec
DELETE FROM game_genre_rel;
