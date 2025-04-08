package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
)

var DB_POOL *pgxpool.Pool

func main() {
	if err := initDB(); err != nil {
		log.Fatalf("db init failed: %v", err)
	}

	var directories map[string]string

	if os.Getenv("ENV") == "DOCKER" {
		directories = map[string]string{
			"/cricsheet/odis_male_csv2":  "ODI",
			"/cricsheet/t20is_male_csv2": "T20I",
			"/cricsheet/ipl_male_csv2":   "T20",
		}
	} else {
		directories = map[string]string{
			"/home/tushar/Desktop/personal/Cricsheet/odis_male_csv2": "ODI",
			"/home/tushar/Desktop/personal/Cricsheet/t20s_male_csv2": "T20I",
			"/home/tushar/Desktop/personal/Cricsheet/ipl_male_csv2":  "T20",
		}
	}

	infoParseChannel := newChannelWrapper[info_parse_response](0)
	matchInitChannel := newChannelWrapper[match_init_response](0)
	bbbChannel := newChannelWrapper[error](0)

	var mainWg sync.WaitGroup
	mainWg.Add(3)

	go triggerParseInfo(directories, infoParseChannel)
	go triggerMatchInit(infoParseChannel, matchInitChannel)
	go triggerMatchBbb(matchInitChannel, bbbChannel)
	go receiveBbb(bbbChannel)

	go infoParseChannel.close(&mainWg)
	go matchInitChannel.close(&mainWg)
	go bbbChannel.close(&mainWg)

	mainWg.Wait()
	fmt.Println("all matches done")
}

func initDB() error {
	var err error

	ctx, DB_URL := context.Background(), os.Getenv("DB_URL")
	DB_POOL, err = dbutils.Connect(ctx, DB_URL)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	return nil
}
