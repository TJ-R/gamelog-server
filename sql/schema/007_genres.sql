-- +goose Up
CREATE TABLE genres (
    id INTEGER PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    slug TEXT,
    url TEXT
);

-- +goose Down
DROP TABLE genres;
