CREATE TABLE IF NOT EXISTS highscores (
    id SERIAL PRIMARY KEY,
    player_name TEXT NOT NULL,
    score INTEGER NOT NULL
);