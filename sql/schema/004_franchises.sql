-- +goose Up
CREATE TABLE franchises (
    id INTEGER PRIMARY KEY,
    create_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    slug TEXT NOT NULL,
    url TEXT
);

-- +goose Down
DROP TABLE franchises;
