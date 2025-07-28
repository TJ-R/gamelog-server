-- +goose Up
CREATE TABLE games (
    id INTEGER PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    first_release_date BIGINT,
    name TEXT NOT NULL,
    parent_game INTEGER REFERENCES games(id),
    slug TEXT,
    storyline TEXT,
    summary TEXT,
    url TEXT,
    version_parent INTEGER REFERENCES games(id),
    version_title TEXT
);

-- +goose Down
DROP TABLE games;
