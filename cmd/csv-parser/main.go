package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/pkg/dotenv"
)

var DB_POOL *pgxpool.Pool

func main() {
	if err := initDB(); err != nil {
		log.Fatalf("db init failed: %v", err)
	}

	matchInfoChannel := make(chan MatchInfoResponse)
	bbbErrChannel := make(chan error)

	basePath := "/home/tushar/Desktop/Cricsheet/odis_male_csv2"
	playingFormat, playingLevel, isMale := "ODI", "international", true

	dirEntries, err := os.ReadDir(basePath)
	if err != nil {
		log.Fatalf("error while reading directory")
	}

	for _, dirEntry := range dirEntries {
		fileName := dirEntry.Name()
		if strings.HasSuffix(fileName, "_info.csv") {
			matchCricsheetId := strings.TrimSuffix(fileName, "_info.csv")

			match, _ := dbutils.ReadMatchByCricsheetId(context.Background(), DB_POOL, matchCricsheetId)
			if match.IsBBBDone.Bool {
				continue
			}

			matchInfoPath := filepath.Join(basePath, fileName)

			matchThreadCounter.increase()
			go extractMatchInfo(matchInfoPath, playingLevel, playingFormat, isMale, matchInfoChannel)
		}
	}

	if !matchThreadCounter.isZero() {
		for matchInfoResponse := range matchInfoChannel {
			if matchInfoResponse.Err != nil {
				log.Printf("error while extracting match info: %v", matchInfoResponse.Err)
				matchThreadCounter.decrease()
				if matchThreadCounter.isZero() {
					fmt.Println("closing match channel")
					close(matchInfoChannel)
				}

				continue
			}

			matchCricsheetId := matchInfoResponse.Match.CricsheetId.String
			match_bbb_path := filepath.Join(basePath, matchCricsheetId+".csv")

			bbbThreadCounter.increase()
			go insertBBB(match_bbb_path, &matchInfoResponse.Match, matchInfoResponse.Team1Info, matchInfoResponse.Team2Info, bbbErrChannel)
			matchThreadCounter.decrease()

			if matchThreadCounter.isZero() {
				fmt.Println("closing match channel")
				close(matchInfoChannel)
			}
		}
	}

	if !bbbThreadCounter.isZero() {
		for err := range bbbErrChannel {
			if err != nil {
				log.Printf("error while inserting BBB data: %v", err)
			}

			bbbThreadCounter.decrease()

			if bbbThreadCounter.isZero() {
				fmt.Println("closing bbb channel")
				close(bbbErrChannel)
			}
		}
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
