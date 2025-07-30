-- +goose Up
CREATE TABLE platform_family (
    id INTEGER PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    slug TEXT
);

-- +goose Down
DROP TABLE platform_family;
