-- +goose Up
CREATE TABLE game_expansion_rel (
    game_id INTEGER NOT NULL REFERENCES games(id),
    expansion_id INTEGER NOT NULL REFERENCES games(id),
    PRIMARY KEY (game_id, expansion_id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE game_expansion_rel;
