package main

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mainlycricket/CricKendra/internal/models"
)

type BattingScorecardEntries map[int64]models.BattingScorecard
type BowlingScorecardEntries map[int64]models.BowlingScorecard

func (entires *BattingScorecardEntries) EnsureEntry(inningsId, batterId int64) {
	if *entires == nil {
		*entires = make(BattingScorecardEntries, 11)
	}

	_, exists := (*entires)[batterId]

	if !exists {
		batPosition := int64(len(*entires)) + 1

		(*entires)[batterId] = models.BattingScorecard{
			InningsId:       pgtype.Int8{Int64: inningsId, Valid: true},
			BatterId:        pgtype.Int8{Int64: batterId, Valid: true},
			BattingPosition: pgtype.Int8{Int64: batPosition, Valid: true},
			RunsScored:      pgtype.Int8{Int64: 0, Valid: true},
			BallsFaced:      pgtype.Int8{Int64: 0, Valid: true},
			FoursScored:     pgtype.Int8{Int64: 0, Valid: true},
			SixesScored:     pgtype.Int8{Int64: 0, Valid: true},
		}
	}
}

func (entires *BattingScorecardEntries) UpdateStrikerEntry(batterId, batterRuns, wides int64) {
	updatedEntry := (*entires)[batterId]

	updatedEntry.RunsScored.Int64 += batterRuns

	if batterRuns == 4 {
		updatedEntry.FoursScored.Int64++
	} else if batterRuns == 6 {
		updatedEntry.SixesScored.Int64++
	}

	if wides == 0 {
		updatedEntry.BallsFaced.Int64++
	}

	(*entires)[batterId] = updatedEntry
}

func (entires *BattingScorecardEntries) AddDismissalEntry(batterId, bowlerId, deliveryId int64, dismissalType string) {
	updatedEntry := (*entires)[batterId]

	updatedEntry.DismissalType = pgtype.Text{String: dismissalType, Valid: true}
	updatedEntry.DismissalBallId = pgtype.Int8{Int64: deliveryId, Valid: true}

	if models.IsBowlerDismissal(dismissalType) {
		updatedEntry.DismissedById = pgtype.Int8{Int64: bowlerId, Valid: true}
	}

	(*entires)[batterId] = updatedEntry
}

func (entires *BowlingScorecardEntries) EnsureEntry(inningsId, bowlerId int64) {
	if *entires == nil {
		*entires = make(BowlingScorecardEntries, 5)
	}

	_, exists := (*entires)[bowlerId]

	if !exists {
		bowlPosition := int64(len(*entires)) + 1

		(*entires)[bowlerId] = models.BowlingScorecard{
			InningsId:       pgtype.Int8{Int64: inningsId, Valid: true},
			BowlerId:        pgtype.Int8{Int64: bowlerId, Valid: true},
			BowlingPosition: pgtype.Int8{Int64: bowlPosition, Valid: true},
			WicketsTaken:    pgtype.Int8{Int64: 0, Valid: true},
			RunsConceded:    pgtype.Int8{Int64: 0, Valid: true},
			BallsBowled:     pgtype.Int8{Int64: 0, Valid: true},
			FoursConceded:   pgtype.Int8{Int64: 0, Valid: true},
			SixesConceded:   pgtype.Int8{Int64: 0, Valid: true},
			WidesConceded:   pgtype.Int8{Int64: 0, Valid: true},
			NoballsConceded: pgtype.Int8{Int64: 0, Valid: true},
		}
	}
}

func (entires *BowlingScorecardEntries) UpdateBowlerEntry(bowlerId, bowlerRuns, batterRuns, wides, noballs int64, isBowlerWicket bool) {
	updatedEntry := (*entires)[bowlerId]

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

	(*entires)[bowlerId] = updatedEntry
}
