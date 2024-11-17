package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"net/url"
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

			if err = dbutils.InsertInnings(context.Background(), tx, &innings); err != nil {
				mainError = err
				return
			}

			inningsQuery["innings_number"] = []string{fmt.Sprintf("%d", inningsNum)}
			dbResponse, err := dbutils.ReadInnings(context.Background(), tx, inningsQuery)
			if err != nil {
				mainError = err
				return
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

		battingScorecardEntries.UpdateStrikerEntry(striker.Id.Int64, batterRuns, wides)

		isBowlerWicket := models.IsBowlerDismissal(row[18]) || models.IsBowlerDismissal(row[20])
		bowlingScorecardEntries.UpdateBowlerEntry(bowler.Id.Int64, bowlerRuns, batterRuns, wides, noballs, isBowlerWicket)

		if row[18] != "" {
			delivery.Player1DismissalType = pgtype.Text{String: row[18], Valid: true}
			dismissedPlayer1 := getPlayerFromCache(&team1Info, &team2Info, row[19])
			delivery.Player1DismissedId = dismissedPlayer1.Id
			battingScorecardEntries.EnsureEntry(innings.Id.Int64, dismissedPlayer1.Id.Int64)
			battingScorecardEntries.AddDismissalEntry(dismissedPlayer1.Id.Int64, bowler.Id.Int64, row[18])
			innings.TotalWkts.Int64++
		}

		if row[20] != "" {
			delivery.Player2DismissalType = pgtype.Text{String: row[20], Valid: true}
			dismissedPlayer2 := getPlayerFromCache(&team1Info, &team2Info, row[21])
			delivery.Player2DismissedId = dismissedPlayer2.Id
			battingScorecardEntries.EnsureEntry(innings.Id.Int64, dismissedPlayer2.Id.Int64)
			battingScorecardEntries.AddDismissalEntry(dismissedPlayer2.Id.Int64, bowler.Id.Int64, row[20])
			innings.TotalWkts.Int64++
		}

		if err = dbutils.InsertDelivery(context.Background(), tx, &delivery); err != nil {
			mainError = err
			return
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
		if err := dbutils.InsertBattingScorecardEntry(context.Background(), tx, &entry); err != nil {
			return fmt.Errorf(`failed to insert batting scorecard entry of player %d: %v`, batterId, err)
		}
	}

	for _, entry := range bowlingScorecardEntries {
		if err := dbutils.InsertBowlingScorecardEntry(context.Background(), tx, &entry); err != nil {
			return err
		}
	}

	return nil
}
