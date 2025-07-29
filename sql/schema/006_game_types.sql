-- +goose Up
CREATE TABLE game_types (
    id INTEGER PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    type TEXT NOT NULL
);

-- +goose Down
DROP TABLE game_types;
