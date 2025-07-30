-- +goose Up
CREATE TABLE platforms_platform_family_rel (
    platform_id INTEGER NOT NULL REFERENCES platforms(id) ON DELETE CASCADE,
    platform_family_id INTEGER NOT NULL REFERENCES platform_family(id) ON DELETE CASCADE,
    PRIMARY KEY (platform_id, platform_family_id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE platforms_platform_family_rel;
