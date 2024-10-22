package dbutils

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mainlycricket/CricKendra/internal/models"
)

func InsertPlayer(ctx context.Context, db *pgxpool.Pool, player *models.Player) error {
	query := `INSERT INTO players (name, playing_role, nationality, is_male, date_of_birth, image_url, biography, batting_styles, primary_batting_style, bowling_styles, primary_bowling_style, teams_represented_id, test_stats, odi_stats, t20i_stats, fc_stats, lista_stats, t20_stats, cricsheet_id, cricinfo_id, cricbuzz_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)`

	cmd, err := db.Exec(ctx, query, player.Name, player.PlayingRole, player.Nationality, player.IsMale, player.DateOfBirth, player.ImageURL, player.Biography, player.BattingStyles, player.PrimaryBattingStyle, player.BowlingStyles, player.PrimaryBowlingStyle, player.TeamsRepresentedId, player.TestStats, player.OdiStats, player.T20iStats, player.FcStats, player.ListAStats, player.T20Stats, player.CricsheetId, player.CricinfoId, player.CricbuzzId)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert player")
	}

	return nil
}

func ReadPlayers(ctx context.Context, db *pgxpool.Pool) ([]models.AllPlayers, error) {
	query := `SELECT id, name, playing_role, nationality, is_male, date_of_birth, primary_batting_style, primary_bowling_style FROM players`

	rows, err := db.Query(ctx, query)

	if err != nil {
		return nil, err
	}

	data, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (models.AllPlayers, error) {
		var player models.AllPlayers

		err := rows.Scan(&player.Id, &player.Name, &player.PlayingRole, &player.Nationality, &player.IsMale, &player.DateOfBirth, &player.PrimaryBattingStyle, &player.PrimaryBowlingStyle)

		return player, err
	})

	return data, err
}

func ReadPlayerById(ctx context.Context, db *pgxpool.Pool, id int) (models.SinglePlayer, error) {
	query := `SELECT id, name, playing_role, nationality, is_male, date_of_birth, image_url, biography, batting_styles, primary_batting_style, bowling_styles, primary_bowling_style, teams_represented, test_stats, odi_stats, t20i_stats, fc_stats, lista_stats, t20_stats, cricsheet_id, cricinfo_id, cricbuzz_id FROM get_player_profile_by_id($1)`

	row := db.QueryRow(ctx, query, id)

	var player models.SinglePlayer

	err := row.Scan(&player.Id, &player.Name, &player.PlayingRole, &player.Nationality, &player.IsMale, &player.DateOfBirth, &player.ImageURL, &player.Biography, &player.BattingStyles, &player.PrimaryBattingStyle, &player.BowlingStyles, &player.PrimaryBowlingStyle, &player.TeamsRepresented, &player.TestStats, &player.OdiStats, &player.T20iStats, &player.FcStats, &player.ListAStats, &player.T20Stats, &player.CricsheetId, &player.CricinfoId, &player.CricbuzzId)

	return player, err
}
