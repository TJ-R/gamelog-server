-- +goose Up
CREATE TABLE game_dlc_rel (
    game_id INTEGER NOT NULL REFERENCES games(id),
    dlc_id INTEGER NOT NULL REFERENCES games(id),
    PRIMARY KEY (game_id, dlc_id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE game_dlc_rel;
