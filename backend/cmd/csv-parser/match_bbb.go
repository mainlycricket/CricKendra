package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mainlycricket/CricKendra/backend/internal/dbutils"
	"github.com/mainlycricket/CricKendra/backend/internal/models"
)

type battingScorecardEntries map[int64]models.BattingScorecard
type bowlingScorecardEntries map[int64]models.BowlingScorecard

type teamInnings struct {
	match                   *models.Match
	innings                 *models.Innings
	battingTeam             *teamInfo
	bowlingTeam             *teamInfo
	battingScorecardEntries battingScorecardEntries
	bowlingScorecardEntries bowlingScorecardEntries
	fallOfWickets           []models.FallOfWicket
	battingPartnerships     []models.BattingPartnership
	startPartnershipFlag    bool
	currentDelivery         *models.Delivery
	maidenOverData          struct {
		bowlerId    string
		balls, runs int64
	}
}

type scoringInput struct {
	ballNumber                                     float64
	strikerId, nonStrikerId, bowlerId              string // cricsheet
	batterRuns                                     int64
	extras, wides, noballs, byes, legbyes, penalty int64
}

type dismissalInput struct {
	dismissedPlayer1Id, dismissedPlayer2Id string // cricsheet
	dismissalType1, dismissalType2         string
}

func insertBBB(filePath string, matchInfo *matchInfo, channel chan<- error) {
	var mainError error

	defer func() {
		if mainError != nil {
			mainError = fmt.Errorf(`%s: %v`, matchInfo.match.CricsheetId.String, mainError)
		}
		channel <- mainError
	}()

	fp, err := os.Open(filePath)
	if err != nil {
		mainError = fmt.Errorf(`error while opening file: %v`, err)
		return
	}
	defer fp.Close()

	tx, err := DB_POOL.Begin(context.Background())
	if err != nil {
		mainError = fmt.Errorf(`error while beginning transaction: %v`, err)
		return
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

	var inningsCount int64

	teamInnings := &teamInnings{
		match:       matchInfo.match,
		battingTeam: matchInfo.team1Info,
		bowlingTeam: matchInfo.team2Info,
	}

	var rowNo int = 2

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
			if err = teamInnings.wrapInnings(tx, false); err != nil {
				mainError = fmt.Errorf(`failed to wrap innings %d: %v`, inningsNum-1, err)
				return
			}
		}

		if inningsNum > inningsCount {
			battingTeamId := matchInfo.team1Info.id
			if row[6] == matchInfo.team2Info.name {
				battingTeamId = matchInfo.team2Info.id
			}

			if err := teamInnings.startInnings(tx, inningsNum, battingTeamId); err != nil {
				mainError = fmt.Errorf(`failed to start innings %d: %v`, inningsNum, err)
				return
			}

			inningsCount++
		}

		scoringInput := &scoringInput{}
		dismissalInput := &dismissalInput{}

		scoringInput.ballNumber, _ = strconv.ParseFloat(row[5], 64)
		scoringInput.strikerId = teamInnings.battingTeam.players[row[8]]
		scoringInput.nonStrikerId = teamInnings.battingTeam.players[row[9]]
		scoringInput.bowlerId = teamInnings.bowlingTeam.players[row[10]]
		scoringInput.batterRuns, _ = strconv.ParseInt(row[11], 10, 64)
		scoringInput.extras, _ = strconv.ParseInt(row[12], 10, 64)
		scoringInput.wides, _ = strconv.ParseInt(row[13], 10, 64)
		scoringInput.noballs, _ = strconv.ParseInt(row[14], 10, 64)
		scoringInput.byes, _ = strconv.ParseInt(row[15], 10, 64)
		scoringInput.legbyes, _ = strconv.ParseInt(row[16], 10, 64)
		scoringInput.penalty, _ = strconv.ParseInt(row[17], 10, 64)

		dismissalInput.dismissalType1 = row[18]
		dismissalInput.dismissedPlayer1Id = teamInnings.battingTeam.players[row[19]]
		dismissalInput.dismissalType2 = row[20]
		dismissalInput.dismissedPlayer2Id = teamInnings.battingTeam.players[row[21]]

		if err := teamInnings.handleDelivery(tx, scoringInput, dismissalInput); err != nil {
			mainError = fmt.Errorf(`failed to insert delivery of row %d: %v`, rowNo, err)
			return
		}

		rowNo++
	}

	if err = teamInnings.wrapInnings(tx, false); err != nil {
		mainError = fmt.Errorf(`failed to wrap last innings: %v`, err)
		return
	}

	// extra innings if only one innings happend
	if teamInnings.innings.InningsNumber.Int64 == 1 {
		if err := teamInnings.startInnings(tx, -1, teamInnings.bowlingTeam.id); err != nil {
			mainError = fmt.Errorf(`failed to start extra innings: %v`, err)
			return
		}

		if err := teamInnings.wrapInnings(tx, false); err != nil {
			mainError = fmt.Errorf(`failed to wrap extra innings: %v`, err)
			return
		}
	}

	if err = dbutils.SetMatchBBBDone(context.Background(), tx, matchInfo.match.Id.Int64); err != nil {
		mainError = fmt.Errorf(`error while setting bbb done: %v`, err)
		return
	}
}

func (teamInnings *teamInnings) wrapInnings(tx pgx.Tx, isExtra bool) error {
	for batterId, entry := range teamInnings.battingScorecardEntries {
		if err := dbutils.InsertBattingScorecardEntry(context.Background(), tx, &entry); err != nil {
			return fmt.Errorf(`failed to insert batting scorecard entry of player %d: %v`, batterId, err)
		}
	}

	for bowlerId, entry := range teamInnings.bowlingScorecardEntries {
		if err := dbutils.InsertBowlingScorecardEntry(context.Background(), tx, &entry); err != nil {
			return fmt.Errorf(`failed to insert bowling scorecard entry of player %d: %v`, bowlerId, err)
		}
	}

	if err := dbutils.InsertFallOfWicketsEntries(context.Background(), tx, teamInnings.fallOfWickets); err != nil {
		return fmt.Errorf(`failed to insert fall of wickets entries: %v`, err)
	}

	teamInnings.endPartnership(teamInnings.currentDelivery, true)
	if err := dbutils.InsertBattingPartnershipEntries(context.Background(), tx, teamInnings.battingPartnerships); err != nil {
		return fmt.Errorf(`failed to insert batting partnerships entries: %v`, err)
	}

	if !isExtra {
		if teamInnings.innings.TotalWkts.Int64 == 10 {
			teamInnings.innings.InningsEnd = pgtype.Text{String: "all_out", Valid: true}
		}

		if err := dbutils.UpdateInnings(context.Background(), tx, teamInnings.innings); err != nil {
			return fmt.Errorf(`failed to update innings: %v`, err)
		}
	}

	return nil
}

func (teamInnings *teamInnings) startInnings(tx pgx.Tx, inningsNumber, battingTeamId int64) error {
	if battingTeamId == teamInnings.bowlingTeam.id {
		teamInnings.battingTeam, teamInnings.bowlingTeam = teamInnings.bowlingTeam, teamInnings.battingTeam
	}

	innings := models.NewInnings(teamInnings.match.Id.Int64, teamInnings.battingTeam.id, teamInnings.bowlingTeam.id)

	if inningsNumber != -1 {
		innings.InningsNumber = pgtype.Int8{Int64: inningsNumber, Valid: true}
	}

	is_limited_over_format := !slices.Contains([]string{"Test", "first_class"}, teamInnings.match.PlayingFormat.String)
	if is_limited_over_format && inningsNumber > 2 {
		innings.IsSuperOver = pgtype.Bool{Bool: true, Valid: true}
	}

	if is_limited_over_format && inningsNumber == 2 && innings.BattingTeamId == teamInnings.match.MatchWinnerId {
		innings.InningsEnd = pgtype.Text{String: "target_reached", Valid: true}
	}

	inningsId, err := dbutils.InsertInnings(context.Background(), tx, innings)
	if err != nil {
		return fmt.Errorf(`failed to insert innings: %v`, err)
	}

	innings.Id = pgtype.Int8{Int64: inningsId, Valid: true}
	teamInnings.innings = innings

	if err := teamInnings.initializeBatting(); err != nil {
		return fmt.Errorf(`error while initialing batting: %v`, err)
	}

	if err := teamInnings.initializeBowling(); err != nil {
		return fmt.Errorf(`error while initialing bowling: %v`, err)
	}

	teamInnings.currentDelivery = &models.Delivery{
		InningsId:             innings.Id,
		InningsDeliveryNumber: pgtype.Int8{Int64: 0, Valid: true},
	}

	teamInnings.startPartnershipFlag = true

	return nil
}

func (teamInnings *teamInnings) initializeBatting() error {
	battersId, err := teamInnings.battingTeam.getPlayersId()
	if err != nil {
		return err
	}

	teamInnings.battingScorecardEntries = make(battingScorecardEntries, 11)
	teamInnings.fallOfWickets = make([]models.FallOfWicket, 0, 10)
	teamInnings.battingPartnerships = make([]models.BattingPartnership, 0, 10)

	for _, batterId := range battersId {
		teamInnings.battingScorecardEntries[batterId] = models.BattingScorecard{
			InningsId:   teamInnings.innings.Id,
			BatterId:    pgtype.Int8{Int64: batterId, Valid: true},
			RunsScored:  pgtype.Int8{Int64: 0, Valid: true},
			BallsFaced:  pgtype.Int8{Int64: 0, Valid: true},
			FoursScored: pgtype.Int8{Int64: 0, Valid: true},
			SixesScored: pgtype.Int8{Int64: 0, Valid: true},
		}
	}

	return nil
}

func (teamInnings *teamInnings) initializeBowling() error {
	bowlersId, err := teamInnings.bowlingTeam.getPlayersId()
	if err != nil {
		return err
	}

	teamInnings.bowlingScorecardEntries = make(bowlingScorecardEntries, 11)

	for _, bowlerId := range bowlersId {
		teamInnings.bowlingScorecardEntries[bowlerId] = models.BowlingScorecard{
			InningsId:       teamInnings.innings.Id,
			BowlerId:        pgtype.Int8{Int64: bowlerId, Valid: true},
			WicketsTaken:    pgtype.Int8{Int64: 0, Valid: true},
			RunsConceded:    pgtype.Int8{Int64: 0, Valid: true},
			BallsBowled:     pgtype.Int8{Int64: 0, Valid: true},
			MaidenOvers:     pgtype.Int8{Int64: 0, Valid: true},
			FoursConceded:   pgtype.Int8{Int64: 0, Valid: true},
			SixesConceded:   pgtype.Int8{Int64: 0, Valid: true},
			WidesConceded:   pgtype.Int8{Int64: 0, Valid: true},
			NoballsConceded: pgtype.Int8{Int64: 0, Valid: true},
		}
	}

	return nil
}

func (teamInnings *teamInnings) setBatPosition(batterId int64) {
	batterEntry := teamInnings.battingScorecardEntries[batterId]
	if batterEntry.BattingPosition.Valid {
		return
	}

	var batPosition int64 = 1

	for _, entry := range teamInnings.battingScorecardEntries {
		if entry.BattingPosition.Valid {
			batPosition++
		}
	}

	batterEntry.HasBatted = pgtype.Bool{Bool: true, Valid: true}
	batterEntry.BattingPosition = pgtype.Int8{Int64: batPosition, Valid: true}
	teamInnings.battingScorecardEntries[batterId] = batterEntry
}

func (teamInnings *teamInnings) updateStrikerScores(batterId int64, scoringInput *scoringInput) {
	teamInnings.setBatPosition(batterId)
	updatedEntry := teamInnings.battingScorecardEntries[batterId]

	updatedEntry.RunsScored.Int64 += scoringInput.batterRuns

	if scoringInput.batterRuns == 4 {
		updatedEntry.FoursScored.Int64++
	} else if scoringInput.batterRuns == 6 {
		updatedEntry.SixesScored.Int64++
	}

	if scoringInput.wides == 0 {
		updatedEntry.BallsFaced.Int64++
	}

	teamInnings.battingScorecardEntries[batterId] = updatedEntry
}

func (teamInnings *teamInnings) startPartnership(delivery *models.Delivery, batter1Id, batter2Id int64) {
	batter1Entry := teamInnings.battingScorecardEntries[batter1Id]
	batter2Entry := teamInnings.battingScorecardEntries[batter2Id]

	partnershipEntry := models.BattingPartnership{
		InningsId:                  teamInnings.innings.Id,
		WicketNumber:               pgtype.Int8{Int64: teamInnings.innings.TotalWkts.Int64 + 1, Valid: true},
		StartInningsDeliveryNumber: delivery.InningsDeliveryNumber,
		Batter1Id:                  batter1Entry.BatterId,
		Batter1Runs:                batter1Entry.RunsScored,
		Batter1Balls:               batter1Entry.BallsFaced,
		Batter2Id:                  batter2Entry.BatterId,
		Batter2Runs:                batter2Entry.RunsScored,
		Batter2Balls:               batter2Entry.BallsFaced,
		StartTeamRuns:              teamInnings.innings.TotalRuns,
		StartBallNumber:            teamInnings.currentDelivery.BallNumber,
		IsUnbeaten:                 pgtype.Bool{Bool: true, Valid: true},
	}

	teamInnings.battingPartnerships = append(teamInnings.battingPartnerships, partnershipEntry)
	teamInnings.startPartnershipFlag = false
}

func (teamInnings *teamInnings) endPartnership(delivery *models.Delivery, isUnbeaten bool) {
	partnershipIdx := len(teamInnings.battingPartnerships) - 1
	if partnershipIdx < 0 {
		return
	}

	currentPartnership := teamInnings.battingPartnerships[partnershipIdx]

	if !currentPartnership.EndInningsDeliveryNumber.Valid {
		batter1Entry := teamInnings.battingScorecardEntries[currentPartnership.Batter1Id.Int64]
		batter2Entry := teamInnings.battingScorecardEntries[currentPartnership.Batter2Id.Int64]

		currentPartnership.Batter1Runs.Int64 = batter1Entry.RunsScored.Int64 - currentPartnership.Batter1Runs.Int64
		currentPartnership.Batter1Balls.Int64 = batter1Entry.BallsFaced.Int64 - currentPartnership.Batter1Balls.Int64
		currentPartnership.Batter2Runs.Int64 = batter2Entry.RunsScored.Int64 - currentPartnership.Batter2Runs.Int64
		currentPartnership.Batter2Balls.Int64 = batter2Entry.BallsFaced.Int64 - currentPartnership.Batter2Balls.Int64
		currentPartnership.EndInningsDeliveryNumber = delivery.InningsDeliveryNumber
		currentPartnership.EndTeamRuns = teamInnings.innings.TotalRuns
		currentPartnership.EndBallNumber = delivery.BallNumber
		currentPartnership.IsUnbeaten.Bool = isUnbeaten

		teamInnings.battingPartnerships[partnershipIdx] = currentPartnership
	}

	teamInnings.startPartnershipFlag = true
}

func (teamInnings *teamInnings) addDismissalEntry(delivery *models.Delivery, isFirst bool) {
	var batterId int64
	var dismissalType string

	if isFirst {
		batterId = delivery.Player1DismissedId.Int64
		dismissalType = delivery.Player1DismissalType.String
	} else {
		batterId = delivery.Player2DismissedId.Int64
		dismissalType = delivery.Player2DismissalType.String
	}

	teamInnings.setBatPosition(batterId)
	updatedEntry := teamInnings.battingScorecardEntries[batterId]

	updatedEntry.DismissalType = pgtype.Text{String: dismissalType, Valid: true}
	updatedEntry.DismissalBallNumber = delivery.InningsDeliveryNumber

	if isFirst {
		updatedEntry.Fielder1Id, updatedEntry.Fielder2Id = delivery.Fielder1Id, delivery.Fielder2Id
	}

	if models.IsBowlerDismissal(dismissalType) {
		updatedEntry.DismissedById = pgtype.Int8{Int64: delivery.BowlerId.Int64, Valid: true}
	}

	teamInnings.battingScorecardEntries[batterId] = updatedEntry

	if models.IsTeamDismissal(dismissalType) {
		teamInnings.innings.TotalWkts.Int64++
	}

	fowEntry := models.FallOfWicket{
		InningsId:             teamInnings.innings.Id,
		InningsDeliveryNumber: delivery.InningsDeliveryNumber,
		BatterId:              pgtype.Int8{Int64: batterId, Valid: true},
		BallNumber:            delivery.BallNumber,
		TeamRuns:              teamInnings.innings.TotalRuns,
		WicketNumber:          teamInnings.innings.TotalWkts,
		DismissalType:         pgtype.Text{String: dismissalType, Valid: true},
	}

	teamInnings.fallOfWickets = append(teamInnings.fallOfWickets, fowEntry)

	teamInnings.endPartnership(delivery, false)
}

func (teamInnings *teamInnings) setBowlPosition(bowlerId int64) {
	bowlerEntry := (teamInnings.bowlingScorecardEntries)[bowlerId]
	if bowlerEntry.BowlingPosition.Valid {
		return
	}

	var bowlPosition int64 = 1

	for _, entry := range teamInnings.bowlingScorecardEntries {
		if entry.BowlingPosition.Valid {
			bowlPosition++
		}
	}

	bowlerEntry.BowlingPosition = pgtype.Int8{Int64: bowlPosition, Valid: true}
	(teamInnings.bowlingScorecardEntries)[bowlerId] = bowlerEntry
}

func (teamInnings *teamInnings) updateBowlerEntry(bowlerId int64, scoringInput *scoringInput, dismissalInput *dismissalInput) {
	teamInnings.setBowlPosition(bowlerId)
	updatedEntry := (teamInnings.bowlingScorecardEntries)[bowlerId]

	bowlerRuns := scoringInput.batterRuns + scoringInput.wides + scoringInput.noballs
	updatedEntry.RunsConceded.Int64 += bowlerRuns
	updatedEntry.WidesConceded.Int64 += scoringInput.wides
	updatedEntry.NoballsConceded.Int64 += scoringInput.noballs

	if scoringInput.wides == 0 && scoringInput.noballs == 0 {
		updatedEntry.BallsBowled.Int64++
	}

	if scoringInput.batterRuns == 4 {
		updatedEntry.FoursConceded.Int64++
	} else if scoringInput.batterRuns == 6 {
		updatedEntry.SixesConceded.Int64++
	}

	if models.IsBowlerDismissal(dismissalInput.dismissalType1) || models.IsBowlerDismissal(dismissalInput.dismissalType2) {
		updatedEntry.WicketsTaken.Int64++
	}

	(teamInnings.bowlingScorecardEntries)[bowlerId] = updatedEntry
}

func (teamInnings *teamInnings) addBowlerMaiden(bowlerId int64) {
	teamInnings.setBowlPosition(bowlerId)
	updatedEntry := (teamInnings.bowlingScorecardEntries)[bowlerId]
	updatedEntry.MaidenOvers.Int64++
	(teamInnings.bowlingScorecardEntries)[bowlerId] = updatedEntry
}

func (teamInnings *teamInnings) handleDelivery(tx pgx.Tx, scoringInput *scoringInput, dismissalInput *dismissalInput) error {
	delivery := teamInnings.currentDelivery
	delivery.InningsDeliveryNumber.Int64++

	// ball number
	delivery.BallNumber = pgtype.Float8{Float64: scoringInput.ballNumber, Valid: true}
	delivery.OverNumber = pgtype.Int8{Int64: int64(scoringInput.ballNumber) + 1, Valid: true}

	// striker
	strikerValue, ok := cachedPlayers.get(playerKey{cricsheet_id: scoringInput.strikerId})
	if !ok {
		return fmt.Errorf(`striker %s not found in cache`, scoringInput.strikerId)
	}
	striker := strikerValue.data
	delivery.BatterId = striker.Id
	delivery.IsBatterRHB = striker.IsRHB
	teamInnings.setBatPosition(striker.Id.Int64)

	// non striker
	nonStrikerValue, ok := cachedPlayers.get(playerKey{cricsheet_id: scoringInput.nonStrikerId})
	if !ok {
		return fmt.Errorf(`nonStriker %s not found in cache`, scoringInput.nonStrikerId)
	}
	nonStriker := nonStrikerValue.data
	delivery.NonStrikerId = nonStriker.Id
	delivery.IsNonStrikerRHB = nonStriker.IsRHB
	teamInnings.setBatPosition(nonStriker.Id.Int64)

	// partnership - position is very crucial
	if teamInnings.startPartnershipFlag {
		teamInnings.startPartnership(delivery, delivery.BatterId.Int64, delivery.NonStrikerId.Int64)
	}

	// bowler
	bowlerValue, ok := cachedPlayers.get(playerKey{cricsheet_id: scoringInput.bowlerId})
	if !ok {
		return fmt.Errorf(`bowler %s not found in cache`, scoringInput.bowlerId)
	}
	bowler := bowlerValue.data
	delivery.BowlerId = bowler.Id
	delivery.BowlingStyle = bowler.PrimaryBowlingStyle

	// batter runs
	delivery.BatterRuns = pgtype.Int8{Int64: scoringInput.batterRuns, Valid: true}

	delivery.IsSix = pgtype.Bool{Bool: false, Valid: true}
	delivery.IsFour = pgtype.Bool{Bool: false, Valid: true}
	if scoringInput.batterRuns == 4 {
		delivery.IsFour = pgtype.Bool{Bool: true, Valid: true}
	} else if scoringInput.batterRuns == 6 {
		delivery.IsSix = pgtype.Bool{Bool: true, Valid: true}
	}

	// individual extras

	delivery.Wides = pgtype.Int8{Int64: scoringInput.wides, Valid: true}
	teamInnings.innings.Wides.Int64 += scoringInput.wides

	delivery.Noballs = pgtype.Int8{Int64: scoringInput.noballs, Valid: true}
	teamInnings.innings.Noballs.Int64 += scoringInput.noballs

	delivery.Byes = pgtype.Int8{Int64: scoringInput.byes, Valid: true}
	teamInnings.innings.Byes.Int64 += scoringInput.byes

	delivery.Legbyes = pgtype.Int8{Int64: scoringInput.legbyes, Valid: true}
	teamInnings.innings.Legbyes.Int64 += scoringInput.legbyes

	delivery.Penalty = pgtype.Int8{Int64: scoringInput.penalty, Valid: true}
	teamInnings.innings.Penalty.Int64 += scoringInput.penalty

	// bowler runs
	bowlerRuns := scoringInput.batterRuns + scoringInput.wides + scoringInput.noballs
	delivery.BowlerRuns = pgtype.Int8{Int64: bowlerRuns, Valid: true}

	// totals: runs, extras & balls
	totalRuns := scoringInput.batterRuns + scoringInput.extras
	delivery.TotalRuns = pgtype.Int8{Int64: totalRuns, Valid: true}
	teamInnings.innings.TotalRuns.Int64 += totalRuns
	delivery.TotalExtras = pgtype.Int8{Int64: scoringInput.extras, Valid: true}
	if scoringInput.wides == 0 && scoringInput.noballs == 0 {
		teamInnings.innings.TotalBalls.Int64++
	}

	if err := teamInnings.handleDismissalInput(dismissalInput); err != nil {
		return err
	}

	teamInnings.updateBowlerEntry(bowler.Id.Int64, scoringInput, dismissalInput)
	teamInnings.updateStrikerScores(striker.Id.Int64, scoringInput)

	if teamInnings.isMaidenOver(scoringInput) {
		teamInnings.addBowlerMaiden(bowler.Id.Int64)
	}

	if err := dbutils.InsertDelivery(context.Background(), tx, delivery); err != nil {
		return fmt.Errorf(`error while inserting delivery: %v`, err)
	}

	if dismissalInput.dismissedPlayer1Id != "" {
		teamInnings.addDismissalEntry(delivery, true)
	}

	if dismissalInput.dismissedPlayer2Id != "" {
		if delivery.Player2DismissedId != delivery.BatterId && delivery.Player2DismissedId != delivery.NonStrikerId {
			batter2Id := delivery.BatterId
			if delivery.Player1DismissedId == batter2Id {
				batter2Id = delivery.NonStrikerId
			}

			teamInnings.startPartnership(delivery, delivery.Player2DismissedId.Int64, batter2Id.Int64)
		}

		teamInnings.addDismissalEntry(delivery, false)
	}

	return nil
}

func (teamInnings *teamInnings) handleDismissalInput(dismissalInput *dismissalInput) error {
	delivery := teamInnings.currentDelivery

	if dismissalInput.dismissedPlayer1Id != "" {
		dismissedPlayerValue, ok := cachedPlayers.get(playerKey{cricsheet_id: dismissalInput.dismissedPlayer1Id})
		if !ok {
			return fmt.Errorf(`dismissed player 1 %s not found in cache`, dismissalInput.dismissedPlayer1Id)
		}
		dismissedPlayer := dismissedPlayerValue.data
		delivery.Player1DismissalType = pgtype.Text{String: dismissalInput.dismissalType1, Valid: true}
		delivery.Player1DismissedId = dismissedPlayer.Id
	} else {
		delivery.Player1DismissedId = pgtype.Int8{}
		delivery.Player1DismissalType = pgtype.Text{}
	}

	if dismissalInput.dismissedPlayer2Id != "" {
		dismissedPlayerValue, ok := cachedPlayers.get(playerKey{cricsheet_id: dismissalInput.dismissedPlayer2Id})
		if !ok {
			return fmt.Errorf(`dismissed player 2 %s not found in cache`, dismissalInput.dismissedPlayer2Id)
		}
		dismissedPlayer := dismissedPlayerValue.data
		delivery.Player2DismissalType = pgtype.Text{String: dismissalInput.dismissalType2, Valid: true}
		delivery.Player2DismissedId = dismissedPlayer.Id
	} else {
		delivery.Player2DismissedId = pgtype.Int8{}
		delivery.Player2DismissalType = pgtype.Text{}
	}

	return nil
}

func (teamInnings *teamInnings) isMaidenOver(scoringInput *scoringInput) bool {
	var isMaiden bool

	// reset balls if bowler changed
	if teamInnings.maidenOverData.bowlerId != scoringInput.bowlerId {
		teamInnings.maidenOverData.balls = 0
		teamInnings.maidenOverData.runs = 0
	}

	teamInnings.maidenOverData.bowlerId = scoringInput.bowlerId
	teamInnings.maidenOverData.runs += scoringInput.batterRuns + scoringInput.wides + scoringInput.noballs

	if scoringInput.wides == 0 && scoringInput.noballs == 0 {
		teamInnings.maidenOverData.balls++
	}

	if teamInnings.maidenOverData.balls == teamInnings.match.BallsPerOver.Int64 && teamInnings.maidenOverData.runs == 0 {
		isMaiden = true
	}

	if teamInnings.maidenOverData.balls == teamInnings.match.BallsPerOver.Int64 {
		teamInnings.maidenOverData.balls = 0
		teamInnings.maidenOverData.runs = 0
	}

	return isMaiden
}
