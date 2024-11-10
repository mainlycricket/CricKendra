package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"

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
	match_bbb_path := "/home/tushar/Desktop/Cricsheet/odis_male_csv2/64814.csv"
	match_cricsheet_id := strings.TrimSuffix(filepath.Base(match_bbb_path), ".csv")

	match, team1, team2, err := extractMatchInfo(match_info_path, "international", "ODI")

	if err != nil {
		log.Fatalf("error while extracting match info: %v", err)
	}

	if err := dbutils.InsertMatch(context.Background(), DB_POOL, &match); err != nil {
		log.Fatalf(`failed to insert match: %v`, err)
	}

	matchQuery := url.Values{
		"cricsheet_id": []string{match_cricsheet_id},
		"limit":        []string{"1"},
	}

	dbResponse, err := dbutils.ReadMatches(context.Background(), DB_POOL, matchQuery)
	if err != nil {
		return
	}
	match.Id = dbResponse.Matches[0].Id

	if err := insertBBB(match_bbb_path, &match, team1, team2); err != nil {
		log.Fatalf(`error while inserting BBB: %v`, err)
	}
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
