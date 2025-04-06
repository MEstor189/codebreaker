package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() error {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		databaseUrl = "postgres://user:password@localhost:5432/dicodebreaker_db"
	}

	var err error
	DB, err = pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	return nil
}

func ConnectDB() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("Failed to connect: %w", err)
	}
	return conn, nil
}
