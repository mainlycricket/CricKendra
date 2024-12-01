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

func InsertSeries(ctx context.Context, db *pgxpool.Pool, series *models.Series) (int64, error) {
	var seriesId int64

	tx, err := db.Begin(ctx)
	if err != nil {
		return seriesId, err
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	query := `INSERT INTO series (name, is_male, playing_level, playing_format, season, tournament_id, parent_series_id, start_date, end_date, winner_team_id, final_status) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`

	err = tx.QueryRow(ctx, query, series.Name, series.IsMale, series.PlayingLevel, series.PlayingFormat, series.Season, series.TournamentId, series.ParentSeriesId, series.StartDate, series.EndDate, series.WinnerTeamId, series.FinalStatus).Scan(&seriesId)
	if err != nil {
		return seriesId, err
	}

	if len(series.TeamsId) > 0 {
		if err = UpsertSeriesTeamEntries(ctx, tx, seriesId, series.TeamsId); err != nil {
			return seriesId, err
		}
	}

	return seriesId, err
}

func UpsertSeriesTeamEntries(ctx context.Context, db DB_Exec, seriesId int64, teamsId []pgtype.Int8) error {
	query := `INSERT INTO series_team_entries (series_id, team_id) VALUES ($1, $2) ON CONFLICT (series_id, team_id) DO NOTHING`

	batch := &pgx.Batch{}
	for _, teamId := range teamsId {
		batch.Queue(query, seriesId, teamId)
	}

	batchResults := db.SendBatch(ctx, batch)
	return batchResults.Close()
}

func ReadSeries(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.AllSeriesResponse, error) {
	var response responses.AllSeriesResponse

	queryInfoInput := pgxutils.QueryInfoInput{
		UrlQuery:     queryMap,
		TableName:    "series",
		DefaultLimit: 20,
		DefaultSort:  []string{"-season"},
	}

	queryInfoOutput, err := pgxutils.ParseQuery[models.Series](queryInfoInput)
	if err != nil {
		return response, err
	}

	query := fmt.Sprintf(`
	
	WITH series AS (
		SELECT s.id, s.name, s.is_male, s.playing_level, s.playing_format, s.season, 
	
		ARRAY_AGG(ste.team_id) as teams_id, 
	
		ARRAY_AGG(ROW(ste.team_id, t.name)) AS teams, 
		
		s.start_date, s.end_date,s.winner_team_id, s.final_status
		
		FROM series s 
		
		LEFT JOIN series_team_entries ste ON s.id = ste.series_id 
		LEFT JOIN teams t ON ste.team_id = t.id 
		GROUP BY s.id
	) 
		
	SELECT 
	
	id, name, is_male, playing_level, playing_format, season, teams, start_date, end_date, winner_team_id, final_status 
	
	FROM series 
	
	%s %s %s`, queryInfoOutput.WhereClause, queryInfoOutput.OrderByClause, queryInfoOutput.PaginationClause)

	rows, err := db.Query(ctx, query, queryInfoOutput.Args...)
	if err != nil {
		return response, err
	}

	seriesList, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.AllSeries, error) {
		var series responses.AllSeries

		err := rows.Scan(&series.Id, &series.Name, &series.IsMale, &series.PlayingLevel, &series.PlayingFormat, &series.Season, &series.Teams, &series.StartDate, &series.EndDate, &series.WinnerTeamId, &series.FinalStatus)

		return series, err
	})

	if len(seriesList) > queryInfoOutput.RecordsCount {
		response.Series = seriesList[:queryInfoOutput.RecordsCount]
		response.Next = true
	} else {
		response.Series = seriesList
		response.Next = false
	}

	return response, err
}

func ReadSeriesById(ctx context.Context, db DB_Exec, id int64) (responses.SingleSeries, error) {
	var series responses.SingleSeries

	query := `SELECT 

	s.id, s.name, s.is_male, s.playing_level, s.playing_format, s.season, s.start_date, s.end_date, s.winner_team_id, s.final_status,

	s.tournament_id, tournaments.name,

	ARRAY_AGG(ROW(ste.team_id, t.name)) AS teams,

	s.parent_series_id, ps.name AS parent_series_name,

	(SELECT 
		ARRAY_AGG(ROW(awards.player_id, players.name))	
		FROM player_awards awards LEFT JOIN players ON awards.player_id = players.id
		WHERE awards.award_type = 'player_of_the_series' AND awards.series_id = s.id
	) AS players_of_the_series,

	get_series_matches(s.id) as matches

	FROM series s 

	LEFT JOIN tournaments ON s.tournament_id = tournaments.id
	LEFT JOIN series_team_entries ste ON s.id = ste.series_id 
	LEFT JOIN teams t ON ste.team_id = t.id 
	LEFT JOIN series ps ON s.parent_series_id = ps.id 

	WHERE s.id = $1 
	GROUP BY s.id, tournaments.name, ps.name;`

	err := db.QueryRow(ctx, query, id).Scan(&series.Id, &series.Name, &series.IsMale, &series.PlayingLevel, &series.PlayingFormat, &series.Season, &series.StartDate, &series.EndDate, &series.WinnerTeamId, &series.FinalStatus, &series.TournamentId, &series.TournamentName, &series.Teams, &series.ParentSeriesId, &series.ParentSeriesName, &series.PlayersOfTheSeries, &series.Matches)

	return series, err
}
