package models

import "github.com/jackc/pgx/v5/pgtype"

type LoginInput struct {
	Email    pgtype.Text `json:"email"`
	Password pgtype.Text `json:"password"`
}

type TossDecisionInput struct {
	MatchId           pgtype.Int8
	TossWinnerId      pgtype.Int8 `json:"toss_winner_team_id"`
	TossLoserId       pgtype.Int8 `json:"toss_loser_team_id"`
	IsTossDecisionBat pgtype.Bool `json:"is_toss_decision_bat"`
}

type BatterPositionInput struct {
	InningsId       pgtype.Int8
	BatterId        pgtype.Int8
	HasBatted       pgtype.Bool `json:"has_batted"`
	BattingPosition pgtype.Int8 `json:"batting_position"`
}

type DeliveryScoringInput struct {
	MatchId               pgtype.Int8
	InningsId             pgtype.Int8
	InningsDeliveryNumber pgtype.Int8 `json:"innings_delivery_number"`

	BallNumber   pgtype.Float8 `json:"ball_number"`
	OverNumber   pgtype.Int8   `json:"over_number"`
	BatterId     pgtype.Int8   `json:"batter_id"`
	BowlerId     pgtype.Int8   `json:"bowler_id"`
	NonStrikerId pgtype.Int8   `json:"non_striker_id"`
	BatterRuns   pgtype.Int8   `json:"batter_runs"`
	Wides        pgtype.Int8   `json:"wides"`
	Noballs      pgtype.Int8   `json:"noballs"`
	Legbyes      pgtype.Int8   `json:"legbyes"`
	Byes         pgtype.Int8   `json:"byes"`
	Penalty      pgtype.Int8   `json:"penalty"`
	TotalExtras  pgtype.Int8   `json:"total_extras"`
	TotalRuns    pgtype.Int8   `json:"total_runs"`
	BowlerRuns   pgtype.Int8   `json:"bowler_runs"`
	IsFour       pgtype.Bool   `json:"is_four"`
	IsSix        pgtype.Bool   `json:"is_six"`

	Player1DismissedId   pgtype.Int8 `json:"player1_dismissed_id"`
	Player1DismissalType pgtype.Text `json:"player1_dismissal_type"`
	Fielder1Id           pgtype.Int8 `json:"fielder1_id"`
	Fielder2Id           pgtype.Int8 `json:"fielder2_id"`

	NewStrikerId    pgtype.Int8 `json:"new_striker_id"`
	NewNonStrikerId pgtype.Int8 `json:"new_non_striker_id"`
	NewBowler1Id    pgtype.Int8 `json:"new_bowler1_id"`
	NewBowler2Id    pgtype.Int8 `json:"new_bowler2_id"`
}

type DeliveryPlayer2DismissedInput struct {
	InningsId             pgtype.Int8
	InningsDeliveryNumber pgtype.Int8

	Player2DismissedId   pgtype.Int8 `json:"player2_dismissed_id"`
	Player2DismissalType pgtype.Text `json:"player2_dismissal_type"`
}

type DeliveryCommentaryInput struct {
	InningsId             pgtype.Int8
	InningsDeliveryNumber pgtype.Int8
	Commentary            pgtype.Text `json:"commentary"`
}

type DeliveryAdvanceInfoInput struct {
	InningsId             pgtype.Int8
	InningsDeliveryNumber pgtype.Int8

	IsPace          pgtype.Bool   `json:"is_pace"`            // true if pacer, false if spin
	BowlingStyle    pgtype.Text   `json:"bowling_style"`      // RAFM, LAFM, LAF etc
	IsBatterRHB     pgtype.Bool   `json:"is_batter_rhb"`      // true if batter is RHB, false if LHB
	IsNonStrikerRHB pgtype.Bool   `json:"is_non_striker_rhb"` // true if non-striker is RHB, false if LHB
	Line            pgtype.Text   `json:"line"`
	Length          pgtype.Text   `json:"length"`
	BallType        pgtype.Text   `json:"ball_type"`  // inswinger, googly
	BallSpeed       pgtype.Float8 `json:"ball_speed"` // kph
	Misc            pgtype.Text   `json:"misc"`       // edged, missed
	WwRegion        pgtype.Text   `json:"ww_region"`  // cover, mid-wkt
	FootType        pgtype.Text   `json:"foot_type"`  // front foot, back foot, step out
	ShotType        pgtype.Text   `json:"shot_type"`  // straight drive, pull shot
}

type InningsEndInput struct {
	InningsId  pgtype.Int8
	InningsEnd pgtype.Text `json:"innings_end"`
}

type InningsCurrentBattersInput struct {
	InningsId    pgtype.Int8
	StrikerId    pgtype.Int8 `json:"striker_id"`
	NonStrikerId pgtype.Int8 `json:"non_striker_id"`
}

type InningsCurrentBowlersInput struct {
	InningsId pgtype.Int8
	Bowler1Id pgtype.Int8 `json:"bowler1_id"`
	Bowler2Id pgtype.Int8 `json:"bowler2_id"`
}

type MatchResultInput struct {
	MatchId pgtype.Int8

	FinalResult          pgtype.Text `json:"final_result"` // winner decided, abandoned, no result
	MatchWinnerId        pgtype.Int8 `json:"match_winner_team_id"`
	MatchLoserId         pgtype.Int8 `json:"match_loser_team_id"`
	BowlOutWinnerId      pgtype.Int8 `json:"bowl_out_winner_id"`
	SuperOverWinnerId    pgtype.Int8 `json:"super_over_winner_id"`
	IsWonByInnings       pgtype.Bool `json:"is_won_by_innings"`
	IsWonByRuns          pgtype.Bool `json:"is_won_by_runs"`
	WinMargin            pgtype.Int8 `json:"win_margin"`                // runs or wickets
	BallsMargin          pgtype.Int8 `json:"balls_remaining_after_win"` // successful chases
	OutcomeSpecialMethod pgtype.Text `json:"outcome_special_method"`
}

type MatchStateInput struct {
	MatchId pgtype.Int8

	State            pgtype.Text `json:"state"`
	StateDescription pgtype.Text `json:"state_description"`
}

type SeriesFinalResult struct {
	SeriesId pgtype.Int8

	WinnerTeamId pgtype.Int8 `json:"winner_team_id"`
	FinalStatus  pgtype.Text `json:"final_status"`
}
