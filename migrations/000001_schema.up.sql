-- Schema compatible with internal/domain/character_type.go
-- Character, Death, Killer

CREATE TABLE IF NOT EXISTS characters (
    name              VARCHAR(255) PRIMARY KEY,
    sex               VARCHAR(50)  NOT NULL DEFAULT '',
    title             VARCHAR(255) NOT NULL DEFAULT '',
    vocation          VARCHAR(100) NOT NULL DEFAULT '',
    world             VARCHAR(100) NOT NULL DEFAULT '',
    residence         VARCHAR(255) NOT NULL DEFAULT '',
    account_status    VARCHAR(50)  NOT NULL DEFAULT '',
    unlocked_titles   INT          NOT NULL DEFAULT 0,
    level             INT          NOT NULL DEFAULT 0,
    achievement_points INT         NOT NULL DEFAULT 0,
    created_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS deaths (
    id             BIGSERIAL PRIMARY KEY,
    character_name VARCHAR(255) NOT NULL REFERENCES characters(name) ON DELETE CASCADE,
    level          INT         NOT NULL DEFAULT 0,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_deaths_character_name ON deaths(character_name);

CREATE TABLE IF NOT EXISTS killers (
    id        BIGSERIAL PRIMARY KEY,
    death_id  BIGINT     NOT NULL REFERENCES deaths(id) ON DELETE CASCADE,
    name      VARCHAR(255) NOT NULL DEFAULT '',
    is_player BOOLEAN   NOT NULL DEFAULT FALSE
);

CREATE INDEX IF NOT EXISTS idx_killers_death_id ON killers(death_id);
