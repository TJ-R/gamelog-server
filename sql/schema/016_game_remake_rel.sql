-- +goose Up
CREATE TABLE game_remake_rel (
    game_id INTEGER NOT NULL REFERENCES games(id),
    remake_id INTEGER NOT NULL REFERENCES games(id),
    PRIMARY KEY (game_id, remake_id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE game_remake_rel;
