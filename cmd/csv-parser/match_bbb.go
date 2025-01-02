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
			mainError = fmt.Errorf(`%s: %v`, matchInfo.CricsheetId.String, mainError)
		}
		channel <- mainError
	}()

	fp, err := os.Open(filePath)
	if err != nil {
		mainError = fmt.Errorf(`error while opening file: %v`, err)
	}
	defer fp.Close()

	tx, err := DB_POOL.Begin(context.Background())
	if err != nil {
		mainError = fmt.Errorf(`error while beginning transaction: %v`, err)
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
		mainError = fmt.Errorf(`error while reading headers: %v`, err)
		return
	}

	var inningsCount, deliveryNumber int64
	var maidenOverData = struct{ bowlerId, runs, balls int64 }{}

	innings := models.Innings{MatchId: matchInfo.Id}
	var battingScorecardEntries BattingScorecardEntries
	var bowlingScorecardEntries BowlingScorecardEntries

	team1PlayersId, err := team1Info.GetPlayersId()
	if err != nil {
		mainError = fmt.Errorf(`error while getting team1 players id from cache: %v`, err)
		return
	}

	team2PlayersId, err := team2Info.GetPlayersId()
	if err != nil {
		mainError = fmt.Errorf(`error while getting team2 players id from cache: %v`, err)
		return
	}

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			mainError = fmt.Errorf(`error while reading file: %v`, err)
			return
		}

		inningsNum, _ := strconv.ParseInt(row[4], 10, 64)

		if inningsNum > inningsCount && inningsNum > 1 {
			if err = wrapInnings(tx, innings, battingScorecardEntries, bowlingScorecardEntries); err != nil {
				mainError = fmt.Errorf(`failed to wrap innings: %v`, err)
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
				mainError = fmt.Errorf(`failed to insert innings: %v`, err)
				return
			}

			if inningsNum == 2 && innings.BattingTeamId == matchInfo.MatchWinnerId && matchInfo.PlayingFormat.String != "Test" && matchInfo.PlayingFormat.String != "first_class" {
				innings.InningsEnd = pgtype.Text{String: "target_reached", Valid: true}
			}

			if innings.BattingTeamId.Int64 == team1Info.Id {
				battingScorecardEntries.EnsurePlayers(inningsId, team1PlayersId)
				bowlingScorecardEntries.EnsurePlayers(inningsId, team2PlayersId)
			} else {
				bowlingScorecardEntries.EnsurePlayers(inningsId, team1PlayersId)
				battingScorecardEntries.EnsurePlayers(inningsId, team2PlayersId)
			}

			innings.Id = pgtype.Int8{Int64: inningsId, Valid: true}
			inningsCount++
			deliveryNumber = 0
		}

		deliveryNumber++
		var delivery = models.Delivery{
			InningsId:             innings.Id,
			InningsDeliveryNumber: pgtype.Int8{Int64: deliveryNumber, Valid: true},
		}

		ballNumber, _ := strconv.ParseFloat(row[5], 64)
		delivery.BallNumber = pgtype.Float8{Float64: ballNumber, Valid: true}
		delivery.OverNumber = pgtype.Int8{Int64: int64(ballNumber) + 1, Valid: true}

		striker := getPlayerFromCache(&team1Info, &team2Info, row[8])
		delivery.BatterId = striker.Id
		delivery.IsBatterRHB = striker.IsRHB

		nonStriker := getPlayerFromCache(&team1Info, &team2Info, row[9])
		delivery.NonStrikerId = nonStriker.Id
		delivery.IsNonStrikerRHB = nonStriker.IsRHB

		bowler := getPlayerFromCache(&team1Info, &team2Info, row[10])
		delivery.BowlerId = bowler.Id
		delivery.BowlingStyle = bowler.PrimaryBowlingStyle

		if bowler.Id.Int64 != maidenOverData.bowlerId {
			maidenOverData.runs = 0
			maidenOverData.balls = 0
			maidenOverData.bowlerId = bowler.Id.Int64
		}

		batterRuns, _ := strconv.ParseInt(row[11], 10, 64)
		delivery.BatterRuns = pgtype.Int8{Int64: batterRuns, Valid: true}
		if batterRuns == 4 {
			delivery.IsFour = pgtype.Bool{Bool: true, Valid: true}
			delivery.IsSix = pgtype.Bool{Bool: false, Valid: true}
		} else if batterRuns == 6 {
			delivery.IsSix = pgtype.Bool{Bool: true, Valid: true}
			delivery.IsFour = pgtype.Bool{Bool: false, Valid: true}
		} else {
			delivery.IsSix = pgtype.Bool{Bool: false, Valid: true}
			delivery.IsFour = pgtype.Bool{Bool: false, Valid: true}
		}

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
		maidenOverData.runs += bowlerRuns

		if wides == 0 && noballs == 0 {
			maidenOverData.balls++

			// if last ball
			if int64(ballNumber-float64(delivery.OverNumber.Int64-1))*10 == matchInfo.BallsPerOver.Int64 {
				if maidenOverData.balls == matchInfo.BallsPerOver.Int64 && maidenOverData.runs == 0 {
					bowlingScorecardEntries.AddBowlerMaiden(bowler.Id.Int64)
				}

				maidenOverData.runs = 0
				maidenOverData.balls = 0
			}
		}

		delivery.TotalRuns = pgtype.Int8{Int64: totalRuns, Valid: true}
		delivery.BowlerRuns = pgtype.Int8{Int64: bowlerRuns, Valid: true}

		innings.TotalRuns.Int64 += totalRuns
		if wides == 0 {
			innings.TotalBalls.Int64++
		}

		battingScorecardEntries.UpdateStrikerEntry(striker.Id.Int64, batterRuns, wides)

		isBowlerWicket := models.IsBowlerDismissal(row[18]) || models.IsBowlerDismissal(row[20])
		bowlingScorecardEntries.UpdateBowlerEntry(bowler.Id.Int64, bowlerRuns, batterRuns, wides, noballs, isBowlerWicket)

		if row[18] != "" {
			dismissalType, dismissedPlayerName := row[18], row[19]
			dismissedPlayer := getPlayerFromCache(&team1Info, &team2Info, dismissedPlayerName)

			delivery.Player1DismissalType = pgtype.Text{String: dismissalType, Valid: true}
			delivery.Player1DismissedId = dismissedPlayer.Id
		}

		if row[20] != "" {
			dismissalType, dismissedPlayerName := row[20], row[21]
			dismissedPlayer := getPlayerFromCache(&team1Info, &team2Info, dismissedPlayerName)

			delivery.Player2DismissalType = pgtype.Text{String: dismissalType, Valid: true}
			delivery.Player2DismissedId = dismissedPlayer.Id
		}

		deliveryId, err := dbutils.InsertDelivery(context.Background(), tx, &delivery)
		if err != nil {
			mainError = fmt.Errorf(`error while inserting delivery: %v`, err)
			return
		}

		if delivery.Player1DismissalType.String != "" {
			battingScorecardEntries.AddDismissalEntry(delivery.Player1DismissedId.Int64, bowler.Id.Int64, deliveryId, delivery.Player1DismissalType.String)

			if models.IsTeamDismissal(delivery.Player1DismissalType.String) {
				innings.TotalWkts.Int64++
			}
		}

		if delivery.Player2DismissalType.String != "" {
			battingScorecardEntries.AddDismissalEntry(delivery.Player2DismissedId.Int64, bowler.Id.Int64, deliveryId, delivery.Player2DismissalType.String)

			if models.IsTeamDismissal(delivery.Player2DismissalType.String) {
				innings.TotalWkts.Int64++
			}
		}
	}

	if err = wrapInnings(tx, innings, battingScorecardEntries, bowlingScorecardEntries); err != nil {
		mainError = fmt.Errorf(`failed to wrap innings: %v`, err)
		return
	}

	if innings.InningsNumber.Int64 == 1 {
		innings.SetDefaultScore()
		innings.BattingTeamId.Int64, innings.BowlingTeamId.Int64 = innings.BowlingTeamId.Int64, innings.BattingTeamId.Int64
		inningsId, err := dbutils.InsertInnings(context.Background(), tx, &innings)
		if err != nil {
			mainError = fmt.Errorf(`failed to insert innings: %v`, err)
			return
		}
		innings.Id = pgtype.Int8{Int64: inningsId, Valid: true}

		battingScorecardEntries = make(BattingScorecardEntries, 11)
		bowlingScorecardEntries = make(BowlingScorecardEntries, 11)
		if innings.BattingTeamId.Int64 == team1Info.Id {
			battingScorecardEntries.EnsurePlayers(inningsId, team1PlayersId)
			bowlingScorecardEntries.EnsurePlayers(inningsId, team2PlayersId)
		} else {
			battingScorecardEntries.EnsurePlayers(inningsId, team2PlayersId)
			bowlingScorecardEntries.EnsurePlayers(inningsId, team1PlayersId)
		}

		if err := wrapInnings(tx, innings, battingScorecardEntries, bowlingScorecardEntries); err != nil {
			mainError = fmt.Errorf(`failed to wrap innings: %v`, err)
			return
		}
	}

	if err = dbutils.SetMatchBBBDone(context.Background(), tx, matchInfo.Id.Int64); err != nil {
		mainError = fmt.Errorf(`error while setting bbb done: %v`, err)
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
	if innings.TotalWkts.Int64 == 10 {
		innings.InningsEnd = pgtype.Text{String: "all_out", Valid: true}
	}

	if err := dbutils.UpdateInnings(context.Background(), tx, &innings); err != nil {
		return fmt.Errorf(`failed to update innings: %v`, err)
	}

	for batterId, entry := range battingScorecardEntries {
		if _, err := dbutils.InsertBattingScorecardEntry(context.Background(), tx, &entry); err != nil {
			return fmt.Errorf(`failed to insert batting scorecard entry of player %d: %v`, batterId, err)
		}
	}

	for bowlerId, entry := range bowlingScorecardEntries {
		if _, err := dbutils.InsertBowlingScorecardEntry(context.Background(), tx, &entry); err != nil {
			return fmt.Errorf(`failed to insert bowling scorecard entry of player %d: %v`, bowlerId, err)
		}
	}

	return nil
}
