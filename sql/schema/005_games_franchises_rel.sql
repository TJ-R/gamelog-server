-- +goose Up
CREATE TABLE franchises_games_rel (
    franchise_id INTEGER NOT NULL REFERENCES franchises(id) ON DELETE CASCADE,
    games_id INTEGER NOT NULL REFERENCES games(id) ON DELETE CASCADE,
    create_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE franchises_games_rel;
