package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/pkg/dotenv"
)

var DB_POOL *pgxpool.Pool

func main() {
	err := initDB()
	if err != nil {
		log.Fatalf("db init failed: %v", err)
	}

	match_info_path := "/home/tushar/Desktop/Cricsheet/odis_male_csv2/64814_info.csv"

	match, err := extractMatchInfo(match_info_path, "international", "odi")

	if err != nil {
		log.Fatalf("error while extracting match info: %v", err)
	}

	fmt.Println(match)
}

func initDB() error {
	basePath, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to read base path: %v", err)
	}

	dotEnvPath := filepath.Join(basePath, ".env")
	err = dotenv.ReadDotEnv(dotEnvPath)
	if err != nil {
		return fmt.Errorf("error while reading .env file: %v", err)
	}

	ctx, DB_URL := context.Background(), os.Getenv("DB_URL")
	DB_POOL, err = dbutils.Connect(ctx, DB_URL)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	return nil
}
