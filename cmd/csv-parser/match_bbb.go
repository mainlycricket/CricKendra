package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
)

func insertBBB(filePath string, matchInfo *models.Match, team1Info, team2Info TeamInfo) error {
	var err error

	fp, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer fp.Close()

	tx, err := DB_POOL.Begin(context.Background())
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback(context.Background())
			log.Println("rolling back inserts...")
		} else {
			tx.Commit(context.Background())
		}
	}()

	reader := csv.NewReader(fp)

	if _, err := reader.Read(); err != nil {
		return err
	}

	var inningsCount int64

	innings := models.Innings{MatchId: matchInfo.Id}
	inningsQuery := url.Values{
		"match_id":       []string{fmt.Sprintf("%v", matchInfo.Id.Int64)},
		"innings_number": []string{"1"},
		"limit":          []string{"1"},
	}

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		inningsNum, _ := strconv.ParseInt(row[4], 10, 64)

		if inningsNum > inningsCount && inningsNum > 1 {
			// update innings
			if err = dbutils.UpdateInnings(context.Background(), tx, &innings); err != nil {
				return err
			}
		}

		if inningsNum > inningsCount {
			innings.SetDefaultScore()

			innings.InningsNumber = pgtype.Int8{Int64: inningsNum, Valid: true}

			if row[6] == team1Info.Name {
				innings.BattingTeamId = pgtype.Int8{Int64: team1Info.Id, Valid: true}
				innings.BowlingTeamId = pgtype.Int8{Int64: team2Info.Id, Valid: true}
			} else {
				innings.BattingTeamId = pgtype.Int8{Int64: team2Info.Id, Valid: true}
				innings.BowlingTeamId = pgtype.Int8{Int64: team1Info.Id, Valid: true}
			}

			if err = dbutils.InsertInnings(context.Background(), tx, &innings); err != nil {
				return err
			}

			inningsQuery["innings_number"] = []string{fmt.Sprintf("%d", inningsNum)}
			dbResponse, err := dbutils.ReadInnings(context.Background(), tx, inningsQuery)
			if err != nil {
				return err
			}

			innings.Id = dbResponse.Innings[0].Id
			inningsCount++
		}

		var delivery = models.Delivery{InningsId: innings.Id}

		ballNumber, _ := strconv.ParseFloat(row[5], 64)
		delivery.BallNumber = pgtype.Float8{Float64: ballNumber, Valid: true}

		striker := getPlayerFromCache(&team1Info, &team2Info, row[8])
		delivery.BatterId = striker.Id
		delivery.IsBatterRHB = striker.IsRHB

		nonStriker := getPlayerFromCache(&team1Info, &team2Info, row[9])
		delivery.NonStrikerId = nonStriker.Id
		delivery.IsNonStrikerRHB = nonStriker.IsRHB

		bowler := getPlayerFromCache(&team1Info, &team2Info, row[10])
		delivery.BowlerId = bowler.Id
		delivery.BowlingStyle = bowler.PrimaryBowlingStyle

		batterRuns, _ := strconv.ParseInt(row[11], 10, 64)
		delivery.BatterRuns = pgtype.Int8{Int64: batterRuns, Valid: true}

		extras, _ := strconv.ParseInt(row[12], 10, 64)
		delivery.TotalExtras = pgtype.Int8{Int64: extras, Valid: true}

		wides, _ := strconv.ParseInt(row[13], 10, 64)
		delivery.Wides = pgtype.Int8{Int64: wides, Valid: true}

		noballs, _ := strconv.ParseInt(row[14], 10, 64)
		delivery.Noballs = pgtype.Int8{Int64: noballs, Valid: true}

		byes, _ := strconv.ParseInt(row[15], 10, 64)
		delivery.Byes = pgtype.Int8{Int64: byes, Valid: true}

		legbyes, _ := strconv.ParseInt(row[16], 10, 64)
		delivery.Legbyes = pgtype.Int8{Int64: legbyes, Valid: true}

		penalty, _ := strconv.ParseInt(row[17], 10, 64)
		delivery.Penalty = pgtype.Int8{Int64: penalty, Valid: true}

		if row[18] != "" {
			delivery.Player1DismissalType = pgtype.Text{String: row[18], Valid: true}
			dismissedPlayer1 := getPlayerFromCache(&team1Info, &team2Info, row[19])
			delivery.Player1DismissedId = dismissedPlayer1.Id
		}

		if row[20] != "" {
			delivery.Player2DismissalType = pgtype.Text{String: row[18], Valid: true}
			dismissedPlayer2 := getPlayerFromCache(&team1Info, &team2Info, row[19])
			delivery.Player2DismissedId = dismissedPlayer2.Id
		}

		if err = dbutils.InsertDelivery(context.Background(), tx, &delivery); err != nil {
			return err
		}
	}

	if err = dbutils.UpdateInnings(context.Background(), tx, &innings); err != nil {
		return err
	}

	return nil
}

func getPlayerFromCache(team1Info, team2Info *TeamInfo, playerName string) responses.AllPlayers {
	var cacheKey PlayerKey

	if cricsheetId, ok := team1Info.Players[playerName]; ok {
		cacheKey.CricsheetId = cricsheetId
	} else if cricsheetId, ok := team2Info.Players[playerName]; ok {
		cacheKey.CricsheetId = cricsheetId
	}

	return PlayersCache[cacheKey]
}
