package dbutils

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
)

func InsertPlayer(ctx context.Context, db *pgxpool.Pool, player *models.Player) error {
	query := `INSERT INTO players (name, full_name, playing_role, nationality, is_male, date_of_birth, image_url, biography, is_rhb, bowling_styles, primary_bowling_style, teams_represented_id, test_stats, odi_stats, t20i_stats, fc_stats, lista_stats, t20_stats, cricsheet_id, cricinfo_id, cricbuzz_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)`

	cmd, err := db.Exec(ctx, query, player.Name, player.FullName, player.PlayingRole, player.Nationality, player.IsMale, player.DateOfBirth, player.ImageURL, player.Biography, player.IsRHB, player.BowlingStyles, player.PrimaryBowlingStyle, player.TeamsRepresentedId, player.TestStats, player.OdiStats, player.T20iStats, player.FcStats, player.ListAStats, player.T20Stats, player.CricsheetId, player.CricinfoId, player.CricbuzzId)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert player")
	}

	return nil
}

func ReadPlayers(ctx context.Context, db *pgxpool.Pool) ([]responses.AllPlayers, error) {
	query := `SELECT id, name, playing_role, nationality, is_male, date_of_birth, is_rhb, primary_bowling_style FROM players`

	rows, err := db.Query(ctx, query)

	if err != nil {
		return nil, err
	}

	data, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.AllPlayers, error) {
		var player responses.AllPlayers

		err := rows.Scan(&player.Id, &player.Name, &player.PlayingRole, &player.Nationality, &player.IsMale, &player.DateOfBirth, &player.IsRHB, &player.PrimaryBowlingStyle)

		return player, err
	})

	return data, err
}

func ReadPlayerById(ctx context.Context, db *pgxpool.Pool, id int) (responses.SinglePlayer, error) {
	query := `SELECT id, name, full_name, playing_role, nationality, is_male, date_of_birth, image_url, biography, is_rhb, bowling_styles, primary_bowling_style, teams_represented, test_stats, odi_stats, t20i_stats, fc_stats, lista_stats, t20_stats, cricsheet_id, cricinfo_id, cricbuzz_id FROM get_player_profile_by_id($1)`

	row := db.QueryRow(ctx, query, id)

	var player responses.SinglePlayer

	err := row.Scan(&player.Id, &player.Name, &player.FullName, &player.PlayingRole, &player.Nationality, &player.IsMale, &player.DateOfBirth, &player.ImageURL, &player.Biography, &player.IsRHB, &player.BowlingStyles, &player.PrimaryBowlingStyle, &player.TeamsRepresented, &player.TestStats, &player.OdiStats, &player.T20iStats, &player.FcStats, &player.ListAStats, &player.T20Stats, &player.CricsheetId, &player.CricinfoId, &player.CricbuzzId)

	return player, err
}
