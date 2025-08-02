-- +goose Up
CREATE TABLE game_remaster_rel (
    game_id INTEGER NOT NULL REFERENCES games(id),
    remaster_id INTEGER NOT NULL REFERENCES games(id),
    PRIMARY KEY (game_id, remaster_id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE game_remaster_rel;
