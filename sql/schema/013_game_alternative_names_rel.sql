-- +goose Up
CREATE TABLE games_alternative_names_rel (
    game_id INTEGER NOT NULL REFERENCES games(id),
    alternative_name_id INTEGER NOT NULL REFERENCES alternative_names(id),
    PRIMARY KEY (game_id, alternative_name_id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE games_alternative_names_rel;
