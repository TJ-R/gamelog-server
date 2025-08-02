-- +goose Up
CREATE TABLE game_genre_rel (
    game_id INTEGER NOT NULL REFERENCES games(id),
    genre_id INTEGER NOT NULL REFERENCES genres(id),
    PRIMARY KEY (game_id, genre_id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE game_genre_rel;
