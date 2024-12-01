package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
)

func insertBBB(filePath string, matchInfo *models.Match, team1Info, team2Info TeamInfo, channel chan<- error) {
	var mainError error

	defer func() {
		if mainError != nil {
			mainError = fmt.Errorf(`error while inserting BBB of %s: %v`, matchInfo.CricsheetId.String, mainError)
		}
		channel <- mainError
	}()

	fp, err := os.Open(filePath)
	if err != nil {
		mainError = err
	}
	defer fp.Close()

	tx, err := DB_POOL.Begin(context.Background())
	if err != nil {
		mainError = err
	}

	defer func() {
		if mainError != nil {
			tx.Rollback(context.Background())
		} else {
			tx.Commit(context.Background())
		}
	}()

	reader := csv.NewReader(fp)

	if _, err := reader.Read(); err != nil {
		mainError = err
		return
	}

	var inningsCount int64

	innings := models.Innings{MatchId: matchInfo.Id}
	var battingScorecardEntries BattingScorecardEntries
	var bowlingScorecardEntries BowlingScorecardEntries

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			mainError = err
			return
		}

		inningsNum, _ := strconv.ParseInt(row[4], 10, 64)

		if inningsNum > inningsCount && inningsNum > 1 {
			if err = wrapInnings(tx, innings, battingScorecardEntries, bowlingScorecardEntries); err != nil {
				mainError = err
				return
			}
			battingScorecardEntries = make(BattingScorecardEntries, 11)
			bowlingScorecardEntries = make(BowlingScorecardEntries, 11)
		}

		if inningsNum > inningsCount {
			innings.SetDefaultScore()

			innings.InningsNumber = pgtype.Int8{Int64: inningsNum, Valid: true}
			if inningsNum > 2 && matchInfo.PlayingFormat.String != "Test" && matchInfo.PlayingFormat.String != "first_class" {
				innings.IsSuperOver = pgtype.Bool{Bool: true, Valid: true}
			}

			if row[6] == team1Info.Name {
				innings.BattingTeamId = pgtype.Int8{Int64: team1Info.Id, Valid: true}
				innings.BowlingTeamId = pgtype.Int8{Int64: team2Info.Id, Valid: true}
			} else {
				innings.BattingTeamId = pgtype.Int8{Int64: team2Info.Id, Valid: true}
				innings.BowlingTeamId = pgtype.Int8{Int64: team1Info.Id, Valid: true}
			}

			inningsId, err := dbutils.InsertInnings(context.Background(), tx, &innings)
			if err != nil {
				mainError = err
				return
			}

			innings.Id = pgtype.Int8{Int64: inningsId, Valid: true}
			inningsCount++
		}

		var delivery = models.Delivery{InningsId: innings.Id}

		ballNumber, _ := strconv.ParseFloat(row[5], 64)
		delivery.BallNumber = pgtype.Float8{Float64: ballNumber, Valid: true}

		striker := getPlayerFromCache(&team1Info, &team2Info, row[8])
		delivery.BatterId = striker.Id
		delivery.IsBatterRHB = striker.IsRHB
		battingScorecardEntries.EnsureEntry(innings.Id.Int64, striker.Id.Int64)

		nonStriker := getPlayerFromCache(&team1Info, &team2Info, row[9])
		delivery.NonStrikerId = nonStriker.Id
		delivery.IsNonStrikerRHB = nonStriker.IsRHB
		battingScorecardEntries.EnsureEntry(innings.Id.Int64, nonStriker.Id.Int64)

		bowler := getPlayerFromCache(&team1Info, &team2Info, row[10])
		delivery.BowlerId = bowler.Id
		delivery.BowlingStyle = bowler.PrimaryBowlingStyle
		bowlingScorecardEntries.EnsureEntry(innings.Id.Int64, bowler.Id.Int64)

		batterRuns, _ := strconv.ParseInt(row[11], 10, 64)
		delivery.BatterRuns = pgtype.Int8{Int64: batterRuns, Valid: true}

		extras, _ := strconv.ParseInt(row[12], 10, 64)
		delivery.TotalExtras = pgtype.Int8{Int64: extras, Valid: true}

		wides, _ := strconv.ParseInt(row[13], 10, 64)
		delivery.Wides = pgtype.Int8{Int64: wides, Valid: true}
		innings.Wides.Int64 += wides

		noballs, _ := strconv.ParseInt(row[14], 10, 64)
		delivery.Noballs = pgtype.Int8{Int64: noballs, Valid: true}
		innings.Noballs.Int64 += noballs

		byes, _ := strconv.ParseInt(row[15], 10, 64)
		delivery.Byes = pgtype.Int8{Int64: byes, Valid: true}
		innings.Byes.Int64 += byes

		legbyes, _ := strconv.ParseInt(row[16], 10, 64)
		delivery.Legbyes = pgtype.Int8{Int64: legbyes, Valid: true}
		innings.Legbyes.Int64 += legbyes

		penalty, _ := strconv.ParseInt(row[17], 10, 64)
		delivery.Penalty = pgtype.Int8{Int64: penalty, Valid: true}
		innings.Penalty.Int64 += penalty

		totalRuns, bowlerRuns := batterRuns+extras, batterRuns+wides+noballs

		delivery.TotalRuns = pgtype.Int8{Int64: totalRuns, Valid: true}
		delivery.BowlerRuns = pgtype.Int8{Int64: bowlerRuns, Valid: true}

		innings.TotalRuns.Int64 += totalRuns
		if wides == 0 {
			innings.TotalBalls.Int64++
		}

		battingScorecardEntries.UpdateStrikerEntry(striker.Id.Int64, batterRuns, wides)

		isBowlerWicket := models.IsBowlerDismissal(row[18]) || models.IsBowlerDismissal(row[20])
		bowlingScorecardEntries.UpdateBowlerEntry(bowler.Id.Int64, bowlerRuns, batterRuns, wides, noballs, isBowlerWicket)

		var player1DismissedId, player2DismissedId int64
		var player1DismissalType, player2DismissalType string

		if row[18] != "" {
			player1DismissalType = row[18]
			dismissedPlayer1 := getPlayerFromCache(&team1Info, &team2Info, row[19])
			player1DismissedId = dismissedPlayer1.Id.Int64

			innings.TotalWkts.Int64++
			battingScorecardEntries.EnsureEntry(innings.Id.Int64, player1DismissedId)
		}

		if row[20] != "" {
			player2DismissalType = row[20]
			dismissedPlayer2 := getPlayerFromCache(&team1Info, &team2Info, row[21])
			player2DismissedId = dismissedPlayer2.Id.Int64

			innings.TotalWkts.Int64++
			battingScorecardEntries.EnsureEntry(innings.Id.Int64, player2DismissedId)
		}

		deliveryId, err := dbutils.InsertDelivery(context.Background(), tx, &delivery)
		if err != nil {
			mainError = err
			return
		}

		if player1DismissalType != "" {
			battingScorecardEntries.AddDismissalEntry(player1DismissedId, bowler.Id.Int64, deliveryId, player1DismissalType)
		}

		if player2DismissalType != "" {
			battingScorecardEntries.AddDismissalEntry(player2DismissedId, bowler.Id.Int64, deliveryId, player2DismissalType)
		}
	}

	if err = wrapInnings(tx, innings, battingScorecardEntries, bowlingScorecardEntries); err != nil {
		mainError = err
		return
	}

	matchInfo.IsBBBDone = pgtype.Bool{Bool: true, Valid: true}
	if err = dbutils.UpdateMatch(context.Background(), tx, matchInfo); err != nil {
		mainError = err
		return
	}
}

func getPlayerFromCache(team1Info, team2Info *TeamInfo, playerName string) responses.AllPlayers {
	var cacheKey PlayerKey

	if cricsheetId, ok := team1Info.Players[playerName]; ok {
		cacheKey.CricsheetId = cricsheetId
	} else if cricsheetId, ok := team2Info.Players[playerName]; ok {
		cacheKey.CricsheetId = cricsheetId
	}

	player, _ := PlayersCache.Get(cacheKey)
	return player
}

func wrapInnings(tx pgx.Tx, innings models.Innings, battingScorecardEntries BattingScorecardEntries, bowlingScorecardEntries BowlingScorecardEntries) error {
	if err := dbutils.UpdateInnings(context.Background(), tx, &innings); err != nil {
		return err
	}

	for batterId, entry := range battingScorecardEntries {
		if _, err := dbutils.InsertBattingScorecardEntry(context.Background(), tx, &entry); err != nil {
			return fmt.Errorf(`failed to insert batting scorecard entry of player %d: %v`, batterId, err)
		}
	}

	for _, entry := range bowlingScorecardEntries {
		if _, err := dbutils.InsertBowlingScorecardEntry(context.Background(), tx, &entry); err != nil {
			return err
		}
	}

	return nil
}
