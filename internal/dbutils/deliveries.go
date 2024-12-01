package dbutils

import (
	"context"

	"github.com/mainlycricket/CricKendra/internal/models"
)

func InsertDelivery(ctx context.Context, db DB_Exec, delivery *models.Delivery) (int64, error) {
	var id int64

	query := `INSERT INTO deliveries (innings_id, ball_number, over_number, batter_id, bowler_id, non_striker_id, batter_runs, wides, noballs, legbyes, byes, penalty, total_extras, total_runs, bowler_runs, is_four, is_six, player1_dismissed_id, player1_dismissal_type, player2_dismissed_id, player2_dismissal_type, is_pace, bowling_style, is_batter_rhb, is_non_striker_rhb, line, length, ball_type, ball_speed, misc, ww_region, foot_type, shot_type, fielder1_id, fielder2_id, commentary, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38) RETURNING id`

	err := db.QueryRow(ctx, query, delivery.InningsId, delivery.BallNumber, delivery.OverNumber, delivery.BatterId, delivery.BowlerId, delivery.NonStrikerId, delivery.BatterRuns, delivery.Wides, delivery.Noballs, delivery.Legbyes, delivery.Byes, delivery.Penalty, delivery.TotalExtras, delivery.TotalRuns, delivery.BowlerRuns, delivery.IsFour, delivery.IsSix, delivery.Player1DismissedId, delivery.Player1DismissalType, delivery.Player2DismissedId, delivery.Player2DismissalType, delivery.IsPace, delivery.BowlingStyle, delivery.IsBatterRHB, delivery.IsNonStrikerRHB, delivery.Line, delivery.Length, delivery.BallType, delivery.BallSpeed, delivery.Misc, delivery.WwRegion, delivery.FootType, delivery.ShotType, delivery.Fielder1Id, delivery.Fielder2Id, delivery.Commentary, delivery.CreatedAt, delivery.UpdatedAt).Scan(&id)

	return id, err
}
