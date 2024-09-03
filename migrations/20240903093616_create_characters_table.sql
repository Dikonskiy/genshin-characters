-- +goose Up
-- +goose StatementBegin
CREATE TABLE Characters (
    id INT AUTO_INCREMENT PRIMARY KEY,
    CharacterName VARCHAR(255),
    Rarity VARCHAR(50),
    Region VARCHAR(100),
    Vision VARCHAR(50),
    Arkhe VARCHAR(50),
    WeaponType VARCHAR(50),
    ReleaseDate DATE,
    Model VARCHAR(255),
    Constellation VARCHAR(100),
    Birthday VARCHAR(20),
    SpecialDish VARCHAR(255),
    Affiliation VARCHAR(100),
    Limited VARCHAR(50),
    VoiceEng VARCHAR(100),
    VoiceCN VARCHAR(100),
    VoiceJP VARCHAR(100),
    VoiceKR VARCHAR(100),
    Ascension VARCHAR(255),
    AscensionSpecialty VARCHAR(255),
    AscensionBoss VARCHAR(255),
    AscensionMaterial02 VARCHAR(255),
    AscensionMaterial24 VARCHAR(255),
    AscensionMaterial46 VARCHAR(255),
    AscensionGem01 VARCHAR(255),
    AscensionGem13 VARCHAR(255),
    AscensionGem35 VARCHAR(255),
    AscensionGem56 VARCHAR(255),
    TalentMaterial VARCHAR(255),
    TalentBook12 VARCHAR(255),
    TalentBook23 VARCHAR(255),
    TalentBook34 VARCHAR(255),
    TalentBook45 VARCHAR(255),
    TalentBook56 VARCHAR(255),
    TalentBook67 VARCHAR(255),
    TalentBook78 VARCHAR(255),
    TalentBook89 VARCHAR(255),
    TalentBook910 VARCHAR(255),
    TalentWeekly VARCHAR(255)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Characters;
-- +goose StatementEnd