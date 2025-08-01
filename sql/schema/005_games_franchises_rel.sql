-- +goose Up
CREATE TABLE franchises_games_rel (
    franchise_id INTEGER NOT NULL REFERENCES franchises(id),
    games_id INTEGER NOT NULL REFERENCES games(id),
    PRIMARY KEY (franchise_id, games_id), -- Composite Primary Key
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE franchises_games_rel;
