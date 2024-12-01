package dbutils

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
	"github.com/mainlycricket/CricKendra/pkg/pgxutils"
)

func InsertPlayer(ctx context.Context, db *pgxpool.Pool, player *models.Player) (int64, error) {
	var playerId int64
	var err error

	tx, err := db.Begin(ctx)
	if err != nil {
		return playerId, err
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	query := `INSERT INTO players (name, full_name, playing_role, nationality, is_male, date_of_birth, image_url, biography, is_rhb, bowling_styles, primary_bowling_style, unavailable_test_stats, unavailable_odi_stats, unavailable_t20i_stats, unavailable_fc_stats, unavailable_lista_stats, unavailable_t20_stats, cricsheet_id, cricinfo_id, cricbuzz_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20) RETURNING id`

	err = tx.QueryRow(ctx, query, player.Name, player.FullName, player.PlayingRole, player.Nationality, player.IsMale, player.DateOfBirth, player.ImageURL, player.Biography, player.IsRHB, player.BowlingStyles, player.PrimaryBowlingStyle, player.TestStats, player.OdiStats, player.T20iStats, player.FcStats, player.ListAStats, player.T20Stats, player.CricsheetId, player.CricinfoId, player.CricbuzzId).Scan(&playerId)
	if err != nil {
		return playerId, err
	}

	if len(player.TeamsRepresentedId) > 0 {
		if err = UpsertPlayerTeamEntries(ctx, tx, playerId, player.TeamsRepresentedId); err != nil {
			return playerId, err
		}
	}

	return playerId, nil
}

func UpsertPlayerTeamEntries(ctx context.Context, db DB_Exec, playerId int64, teamsId []pgtype.Int8) error {
	query := `INSERT INTO player_team_entries (player_id, team_id) VALUES ($1, $2) ON CONFLICT (player_id, team_id) DO NOTHING`

	batch := &pgx.Batch{}
	for _, teamId := range teamsId {
		batch.Queue(query, playerId, teamId)
	}

	batchResults := db.SendBatch(ctx, batch)
	return batchResults.Close()
}

func UpsertPlayerTeamEntry(ctx context.Context, db DB_Exec, entry *models.PlayerTeamEntry) error {
	query := `INSERT INTO player_team_entries (player_id, team_id) VALUES($1, $2) ON CONFLICT (player_id, team_id) DO NOTHING`

	_, err := db.Exec(ctx, query, entry.PlayerId, entry.TeamId)

	return err
}

func ReadPlayers(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.AllPlayersResponse, error) {
	var response responses.AllPlayersResponse

	queryInfoInput := pgxutils.QueryInfoInput{
		UrlQuery:     queryMap,
		TableName:    "players",
		DefaultLimit: 20,
		DefaultSort:  []string{"id"},
	}

	queryInfoOutput, err := pgxutils.ParseQuery[models.Player](queryInfoInput)
	if err != nil {
		return response, err
	}

	query := fmt.Sprintf(`SELECT id, name, playing_role, nationality, is_male, date_of_birth, is_rhb, primary_bowling_style FROM players %s %s %s`, queryInfoOutput.WhereClause, queryInfoOutput.OrderByClause, queryInfoOutput.PaginationClause)

	rows, err := db.Query(ctx, query, queryInfoOutput.Args...)
	if err != nil {
		return response, err
	}

	players, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.AllPlayers, error) {
		var player responses.AllPlayers

		err := rows.Scan(&player.Id, &player.Name, &player.PlayingRole, &player.Nationality, &player.IsMale, &player.DateOfBirth, &player.IsRHB, &player.PrimaryBowlingStyle)

		return player, err
	})

	if len(players) > queryInfoOutput.RecordsCount {
		response.Players = players[:queryInfoOutput.RecordsCount]
		response.Next = true
	} else {
		response.Players = players
		response.Next = false
	}

	return response, err
}

func ReadPlayerById(ctx context.Context, db *pgxpool.Pool, id int) (responses.SinglePlayer, error) {
	query := `SELECT 
	
	id, name, full_name, playing_role, nationality, is_male, date_of_birth, image_url, biography, is_rhb, bowling_styles, primary_bowling_style, 
	
	(SELECT 
		ARRAY_AGG(ROW(pte.team_id, teams.name)) 
		FROM player_team_entries pte
		LEFT JOIN teams ON pte.team_id = teams.id
		WHERE pte.player_id = players.id
	) AS teams_represented, 
	 
	db_test_stats, db_odi_stats, db_t20i_stats, db_fc_stats, db_lista_stats, db_t20_stats, cricsheet_id, cricinfo_id, cricbuzz_id 
	
	FROM players 
	WHERE id = $1`

	row := db.QueryRow(ctx, query, id)

	var player responses.SinglePlayer

	err := row.Scan(&player.Id, &player.Name, &player.FullName, &player.PlayingRole, &player.Nationality, &player.IsMale, &player.DateOfBirth, &player.ImageURL, &player.Biography, &player.IsRHB, &player.BowlingStyles, &player.PrimaryBowlingStyle, &player.TeamsRepresented, &player.TestStats, &player.OdiStats, &player.T20iStats, &player.FcStats, &player.ListAStats, &player.T20Stats, &player.CricsheetId, &player.CricinfoId, &player.CricbuzzId)

	return player, err
}

func ReadBowlingStyleOptions(ctx context.Context, db *pgxpool.Pool) ([]string, error) {
	query := `SELECT unnest(enum_range(null::bowling_style))`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	bowlingStyles, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (string, error) {
		var bowlingStyle string
		err := row.Scan(&bowlingStyle)
		return bowlingStyle, err
	})

	return bowlingStyles, err
}

func ReadDismissalTypeOptions(ctx context.Context, db *pgxpool.Pool) ([]string, error) {
	query := `SELECT unnest(enum_range(null::dismissal_type))`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	dismissalTypes, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (string, error) {
		var dismissalType string
		err := row.Scan(&dismissalType)
		return dismissalType, err
	})

	return dismissalTypes, err
}

func ReadPlayingStatusOptions(ctx context.Context, db *pgxpool.Pool) ([]string, error) {
	query := `SELECT unnest(enum_range(null::playing_status))`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	playingStatusOptions, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (string, error) {
		var playingStatus string
		err := row.Scan(&playingStatus)
		return playingStatus, err
	})

	return playingStatusOptions, err
}
