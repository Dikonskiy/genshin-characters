-- +goose Up
-- +goose StatementBegin
CREATE UNIQUE INDEX idx_unique_character ON Characters (
    CharacterName,
    Rarity,
    Region,
    Vision,
    WeaponType,
    ReleaseDate,
    Model
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE Characters DROP INDEX idx_unique_character;
-- +goose StatementEnd