-- +goose Up
CREATE TABLE platforms (
    id INTEGER PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    abbreviation TEXT,
    generation INTEGER,
    name TEXT NOT NULL,
    slug TEXT,
    summary TEXT,
    url TEXT
);

-- +goose Down
DROP TABLE platforms;
