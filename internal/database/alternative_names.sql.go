// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: alternative_names.sql

package database

import (
	"context"
)

const createAlternativeName = `-- name: CreateAlternativeName :one
INSERT INTO alternative_names (
    id, created_at, updated_at, comment, game, name
) VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    $4
)
RETURNING id, created_at, updated_at, comment, game, name
`

type CreateAlternativeNameParams struct {
	ID      int32
	Comment string
	Game    int32
	Name    string
}

func (q *Queries) CreateAlternativeName(ctx context.Context, arg CreateAlternativeNameParams) (AlternativeName, error) {
	row := q.db.QueryRowContext(ctx, createAlternativeName,
		arg.ID,
		arg.Comment,
		arg.Game,
		arg.Name,
	)
	var i AlternativeName
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Comment,
		&i.Game,
		&i.Name,
	)
	return i, err
}

const deleteAllAlternativeNames = `-- name: DeleteAllAlternativeNames :exec
DELETE FROM alternative_names
`

func (q *Queries) DeleteAllAlternativeNames(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllAlternativeNames)
	return err
}
