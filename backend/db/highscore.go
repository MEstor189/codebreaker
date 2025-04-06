package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type HighscoreEntry struct {
	PlayerName string
	Score      int
}

func CreateHighscoreEntry(playername string, score int) HighscoreEntry {
	return HighscoreEntry{
		PlayerName: playername,
		Score:      score,
	}
}

func InsertHighscore(conn *pgx.Conn, highscoreEntry HighscoreEntry) error {
	_, err := conn.Exec(
		context.Background(),
		`INSERT INTO highscores (player_name, score) VALUES ($1, $2)`,
		highscoreEntry.PlayerName,
		highscoreEntry.Score,
	)
	if err != nil {
		return fmt.Errorf("Insert failed: %w", err)
	}
	return nil
}

func GetTop10Highscores(conn *pgx.Conn) ([]HighscoreEntry, error) {
	rows, err := conn.Query(
		context.Background(),
		`SELCET (layer_name, score) FROM highscores ORDERD BY score DESC LIMIT 10`,
	)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var highscores []HighscoreEntry

	for rows.Next() {
		var entry HighscoreEntry
		if err := rows.Scan(&entry.PlayerName, &entry.Score); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		highscores = append(highscores, entry)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("row iteration error: %w", rows.Err())
	}
	return highscores, nil
}
