CREATE TABLE IF NOT EXISTS highscores (
    id BIGSERIAL PRIMARY KEY,
    category VARCHAR(255) NOT NULL DEFAULT '',
    world VARCHAR(255) NOT NULL DEFAULT '',
    highscores_rank INT NOT NULL DEFAULT 0,
    highscores_value INT NOT NULL DEFAULT 0,
    char_name VARCHAR(255) NOT NULL DEFAULT '',
    vocation VARCHAR(255) NOT NULL DEFAULT '',
    char_level INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (category, world, highscores_rank)
);
