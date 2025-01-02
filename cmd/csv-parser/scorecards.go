package main

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mainlycricket/CricKendra/internal/models"
)

type BattingScorecardEntries map[int64]models.BattingScorecard
type BowlingScorecardEntries map[int64]models.BowlingScorecard

func (entries *BattingScorecardEntries) EnsurePlayers(inningsId int64, battersId []int64) {
	if *entries == nil {
		*entries = make(BattingScorecardEntries, 11)
	}

	for _, batterId := range battersId {
		(*entries)[batterId] = models.BattingScorecard{
			InningsId:   pgtype.Int8{Int64: inningsId, Valid: true},
			BatterId:    pgtype.Int8{Int64: batterId, Valid: true},
			RunsScored:  pgtype.Int8{Int64: 0, Valid: true},
			BallsFaced:  pgtype.Int8{Int64: 0, Valid: true},
			FoursScored: pgtype.Int8{Int64: 0, Valid: true},
			SixesScored: pgtype.Int8{Int64: 0, Valid: true},
		}
	}
}

func (entries *BowlingScorecardEntries) EnsurePlayers(inningsId int64, bowlersId []int64) {
	if *entries == nil {
		*entries = make(BowlingScorecardEntries, 11)
	}

	for _, bowlerId := range bowlersId {
		(*entries)[bowlerId] = models.BowlingScorecard{
			InningsId:       pgtype.Int8{Int64: inningsId, Valid: true},
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
}

func (entries *BattingScorecardEntries) SetBatPosition(batterId int64) {
	batterEntry := (*entries)[batterId]
	if batterEntry.BattingPosition.Valid {
		return
	}

	var batPosition int64 = 1

	for _, entry := range *entries {
		if entry.BattingPosition.Valid {
			batPosition++
		}
	}

	batterEntry.BattingPosition = pgtype.Int8{Int64: batPosition, Valid: true}
	(*entries)[batterId] = batterEntry
}

func (entries *BattingScorecardEntries) UpdateStrikerEntry(batterId, batterRuns, wides int64) {
	entries.SetBatPosition(batterId)
	updatedEntry := (*entries)[batterId]

	updatedEntry.RunsScored.Int64 += batterRuns

	if batterRuns == 4 {
		updatedEntry.FoursScored.Int64++
	} else if batterRuns == 6 {
		updatedEntry.SixesScored.Int64++
	}

	if wides == 0 {
		updatedEntry.BallsFaced.Int64++
	}

	(*entries)[batterId] = updatedEntry
}

func (entries *BattingScorecardEntries) AddDismissalEntry(batterId, bowlerId, deliveryId int64, dismissalType string) {
	entries.SetBatPosition(batterId)
	updatedEntry := (*entries)[batterId]

	updatedEntry.DismissalType = pgtype.Text{String: dismissalType, Valid: true}
	updatedEntry.DismissalBallId = pgtype.Int8{Int64: deliveryId, Valid: true}

	if models.IsBowlerDismissal(dismissalType) {
		updatedEntry.DismissedById = pgtype.Int8{Int64: bowlerId, Valid: true}
	}

	(*entries)[batterId] = updatedEntry
}

func (entries *BowlingScorecardEntries) SetBowlPosition(bowlerId int64) {
	bowlerEntry := (*entries)[bowlerId]
	if bowlerEntry.BowlingPosition.Valid {
		return
	}

	var bowlPosition int64 = 1

	for _, entry := range *entries {
		if entry.BowlingPosition.Valid {
			bowlPosition++
		}
	}

	bowlerEntry.BowlingPosition = pgtype.Int8{Int64: bowlPosition, Valid: true}
	(*entries)[bowlerId] = bowlerEntry
}

func (entries *BowlingScorecardEntries) UpdateBowlerEntry(bowlerId, bowlerRuns, batterRuns, wides, noballs int64, isBowlerWicket bool) {
	entries.SetBowlPosition(bowlerId)
	updatedEntry := (*entries)[bowlerId]

	updatedEntry.RunsConceded.Int64 += bowlerRuns
	updatedEntry.WidesConceded.Int64 += wides
	updatedEntry.NoballsConceded.Int64 += noballs

	if wides == 0 && noballs == 0 {
		updatedEntry.BallsBowled.Int64++
	}

	if batterRuns == 4 {
		updatedEntry.FoursConceded.Int64++
	} else if batterRuns == 6 {
		updatedEntry.SixesConceded.Int64++
	}

	if isBowlerWicket {
		updatedEntry.WicketsTaken.Int64++
	}

	(*entries)[bowlerId] = updatedEntry
}

func (entries *BowlingScorecardEntries) AddBowlerMaiden(bowlerId int64) {
	entries.SetBowlPosition(bowlerId)
	updatedEntry := (*entries)[bowlerId]
	updatedEntry.MaidenOvers.Int64++
	(*entries)[bowlerId] = updatedEntry
}
