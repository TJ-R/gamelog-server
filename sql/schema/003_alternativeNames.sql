-- +goose Up
CREATE TABLE alternative_names (
    id INTEGER PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    comment TEXT NOT NULL,
    game INTEGER NOT NULL REFERENCES games(id) ON DELETE CASCADE,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE alternative_names;
