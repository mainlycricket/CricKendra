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
		return -1, err
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	query := `INSERT INTO series (name, is_male, playing_level, playing_format, season, tournament_id, tour_flag, start_date, end_date, winner_team_id, final_status) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`

	err = tx.QueryRow(ctx, query, series.Name, series.IsMale, series.PlayingLevel, series.PlayingFormat, series.Season, series.TournamentId, series.TourFlag, series.StartDate, series.EndDate, series.WinnerTeamId, series.FinalStatus).Scan(&seriesId)
	if err != nil {
		return -1, err
	}

	if len(series.TeamsId) > 0 {
		seriesId := pgtype.Int8{Int64: seriesId, Valid: true}
		if err = UpsertSeriesTeamEntries(ctx, tx, seriesId, series.TeamsId); err != nil {
			return -1, err
		}
	}

	return seriesId, err
}

func UpsertSeriesTeamEntries(ctx context.Context, db DB_Exec, seriesId pgtype.Int8, teamsId []pgtype.Int8) error {
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
		
		s.start_date, s.end_date,s.winner_team_id, s.final_status, s.tour_flag
		
		FROM series s 
		
		LEFT JOIN series_team_entries ste ON s.id = ste.series_id 
		LEFT JOIN teams t ON ste.team_id = t.id 
		GROUP BY s.id
	) 
		
	SELECT 
	
	id, name, is_male, playing_level, playing_format, season, teams, start_date, end_date, winner_team_id, final_status, tour_flag 
	
	FROM series 
	
	%s %s %s`, queryInfoOutput.WhereClause, queryInfoOutput.OrderByClause, queryInfoOutput.PaginationClause)

	rows, err := db.Query(ctx, query, queryInfoOutput.Args...)
	if err != nil {
		return response, err
	}

	seriesList, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.AllSeries, error) {
		var series responses.AllSeries

		err := rows.Scan(&series.Id, &series.Name, &series.IsMale, &series.PlayingLevel, &series.PlayingFormat, &series.Season, &series.Teams, &series.StartDate, &series.EndDate, &series.WinnerTeamId, &series.FinalStatus, &series.TourFlag)

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

var seriesHeaderQuery = struct {
	withClause    string
	selectFields  string
	joins         string
	groupByFields string
}{
	withClause: `WITH
	top_batters AS (
		SELECT
			-- order is crucial for struct unpacking
			bs.batter_id, players.name, players.image_url,
			COUNT(DISTINCT bs.innings_id), SUM(bs.runs_scored),
			(
				CASE
					WHEN COUNT(
						CASE
							WHEN bs.dismissal_type IS NOT NULL
							AND bs.dismissal_type NOT IN ('retired hurt', 'retired not out') THEN 1
						END
					) > 0 THEN SUM(bs.runs_scored) * 1.0 / COUNT(
						CASE
							WHEN bs.dismissal_type IS NOT NULL
							AND bs.dismissal_type NOT IN ('retired hurt', 'retired not out') THEN 1
						END
					)
					ELSE NULL
				END
			)
		FROM batting_scorecards bs
			LEFT JOIN innings ON bs.innings_id = innings.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE
			LEFT JOIN match_series_entries mse ON mse.match_id = innings.match_id
			LEFT JOIN players ON bs.batter_id = players.id
		WHERE mse.series_id = $1
		GROUP BY bs.batter_id, players.name, players.image_url
		ORDER BY
			SUM(bs.runs_scored) DESC,
			(
				CASE
					WHEN COUNT(
						CASE
							WHEN bs.dismissal_type IS NOT NULL
							AND bs.dismissal_type NOT IN ('retired hurt', 'retired not out') THEN 1
						END
					) > 0 THEN SUM(bs.runs_scored) * 1.0 / COUNT(
						CASE
							WHEN bs.dismissal_type IS NOT NULL
							AND bs.dismissal_type NOT IN ('retired hurt', 'retired not out') THEN 1
						END
					)
					ELSE NULL
				END
			) DESC,
			COUNT(DISTINCT bs.innings_id) ASC
		FETCH FIRST 3 rows ONLY
	),
	top_bowlers AS (
		SELECT
			-- order is crucial for struct unpacking
			bs.bowler_id, players.name, players.image_url,
			COUNT(DISTINCT bs.innings_id), SUM(bs.wickets_taken),
			(
				CASE
					WHEN SUM(bs.wickets_taken) > 0 THEN SUM(bs.runs_conceded) * 1.0 / SUM(bs.wickets_taken)
					ELSE NULL
				END
			)
		FROM bowling_scorecards bs
			LEFT JOIN innings ON bs.innings_id = innings.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE
			LEFT JOIN match_series_entries mse ON mse.match_id = innings.match_id
			LEFT JOIN players ON bs.bowler_id = players.id
		WHERE mse.series_id = $1
		GROUP BY bs.bowler_id, players.name, players.image_url
		ORDER BY
			SUM(bs.wickets_taken) DESC,
			(
				CASE
					WHEN SUM(bs.wickets_taken) > 0 THEN SUM(bs.runs_conceded) * 1.0 / SUM(bs.wickets_taken)
					ELSE NULL
				END
			) ASC,
			COUNT(DISTINCT bs.innings_id) ASC
		FETCH FIRST 3 rows ONLY
	)`,
	selectFields: `series.id, series.name, series.season,
				   series.tournament_id, tournaments.name,
				   ( SELECT ARRAY_AGG(top_batters.*) FROM top_batters ) AS top_batters,
				   ( SELECT ARRAY_AGG(top_bowlers.*) FROM top_bowlers ) AS top_bowlers`,
	joins:         `LEFT JOIN tournaments ON series.tournament_id = tournaments.id`,
	groupByFields: `series.id, tournaments.name`,
}

func ReadSeriesOverviewById(ctx context.Context, db DB_Exec, seriesId int64) (responses.SingleSeriesOverview, error) {
	var seriesOverview responses.SingleSeriesOverview
	header := &seriesOverview.SeriesHeader

	query := fmt.Sprintf(`
		%s,
		fixture_matches AS (
			SELECT %s
			FROM matches
			LEFT JOIN match_series_entries mse ON mse.match_id = matches.id
			%s
			WHERE mse.series_id = $1 AND matches.final_result IS NULL
			GROUP BY %s
			ORDER BY matches.start_date ASC, matches.start_time ASC
			FETCH FIRST 10 ROWS ONLY
		),
		result_matches AS (
			SELECT %s
			FROM matches
			LEFT JOIN match_series_entries mse ON mse.match_id = matches.id
			%s
			WHERE mse.series_id = $1 AND matches.final_result IS NOT NULL
			GROUP BY %s
			ORDER BY matches.start_date DESC, matches.start_time DESC
			FETCH FIRST 10 ROWS ONLY
		)
		SELECT 
			%s,
			series.winner_team_id, winner_team.name, series.final_status,
			( SELECT ARRAY_AGG(fixture_matches.*) FROM fixture_matches ) AS fixture_matches,
			( SELECT ARRAY_AGG(result_matches.*) FROM result_matches ) AS result_matches
		FROM series
		   	%s
		    LEFT JOIN teams winner_team ON series.winner_team_id = winner_team.id
		WHERE
		    series.id = $1
		GROUP BY
		    %s,
			winner_team.name
	`,
		seriesHeaderQuery.withClause,
		matchInfoQuery.selectFields, matchInfoQuery.joins, matchHeaderQuery.groupByFields,
		matchInfoQuery.selectFields, matchInfoQuery.joins, matchHeaderQuery.groupByFields,
		seriesHeaderQuery.selectFields, seriesHeaderQuery.joins, seriesHeaderQuery.groupByFields)

	err := db.QueryRow(ctx, query, seriesId).Scan(
		&header.SeriesId, &header.SeriesName, &header.Season, &header.TournamentId, &header.TournamentName,
		&header.TopBatters, &header.TopBowlers,
		&seriesOverview.WinnerTeamId, &seriesOverview.WinnerTeamName, &seriesOverview.FinalStatus,
		&seriesOverview.FixtureMatches, &seriesOverview.ResultMatches,
	)

	if err != nil {
		return seriesOverview, err
	}

	return seriesOverview, err
}

func ReadSeriesMatchesById(ctx context.Context, db DB_Exec, seriesId int64) (responses.SingleSeriesMatches, error) {
	var seriesWithMatches responses.SingleSeriesMatches
	header := &seriesWithMatches.SeriesHeader

	query := fmt.Sprintf(`
		%s,
		series_matches AS (
			SELECT %s
			FROM matches
			LEFT JOIN match_series_entries mse ON mse.match_id = matches.id
			%s
			WHERE mse.series_id = $1
			GROUP BY %s
			ORDER BY matches.start_date ASC, matches.start_time ASC
		)
		SELECT 
			%s,
			( SELECT ARRAY_AGG(series_matches.*) FROM series_matches ) AS series_matches
		FROM series
		   	%s
		    LEFT JOIN teams winner_team ON series.winner_team_id = winner_team.id
		WHERE
		    series.id = $1
		GROUP BY
		    %s
	`,
		seriesHeaderQuery.withClause,
		matchInfoQuery.selectFields, matchInfoQuery.joins, matchHeaderQuery.groupByFields,
		seriesHeaderQuery.selectFields, seriesHeaderQuery.joins, seriesHeaderQuery.groupByFields)

	err := db.QueryRow(ctx, query, seriesId).Scan(
		&header.SeriesId, &header.SeriesName, &header.Season, &header.TournamentId, &header.TournamentName,
		&header.TopBatters, &header.TopBowlers,
		&seriesWithMatches.Matches,
	)

	if err != nil {
		return seriesWithMatches, err
	}

	return seriesWithMatches, err
}

func ReadSeriesTeamsById(ctx context.Context, db DB_Exec, seriesId int64) (responses.SingleSeriesTeams, error) {
	var seriesWithTeams responses.SingleSeriesTeams
	header := &seriesWithTeams.SeriesHeader

	query := fmt.Sprintf(`
		%s
		SELECT 
			%s,
			-- order is crucial for struct unpacking
			ARRAY_AGG ( ROW ( teams.id, teams.name, teams.image_url ) )
		FROM series
		   	%s
		    LEFT JOIN series_team_entries ON series_team_entries.series_id = series.id
			LEFT JOIN teams ON series_team_entries.team_id = teams.id
		WHERE
		    series.id = $1
		GROUP BY
		    %s
	`,
		seriesHeaderQuery.withClause,
		seriesHeaderQuery.selectFields, seriesHeaderQuery.joins, seriesHeaderQuery.groupByFields)

	err := db.QueryRow(ctx, query, seriesId).Scan(
		&header.SeriesId, &header.SeriesName, &header.Season, &header.TournamentId, &header.TournamentName,
		&header.TopBatters, &header.TopBowlers,
		&seriesWithTeams.Teams,
	)

	if err != nil {
		return seriesWithTeams, err
	}

	return seriesWithTeams, err
}

func ReadSeriesSquadsListById(ctx context.Context, db DB_Exec, seriesId int64) (responses.SingleSeriesSquadsList, error) {
	var seriesWithSquadsList responses.SingleSeriesSquadsList
	header := &seriesWithSquadsList.SeriesHeader

	query := fmt.Sprintf(`
		%s
		SELECT 
			%s,
			-- order is crucial for struct unpacking
			ARRAY_AGG ( ROW ( series_squads.id, series_squads.squad_label, teams.image_url ) )
		FROM series
		   	%s
		    LEFT JOIN series_squads ON series_squads.series_id = series.id
			LEFT JOIN teams ON series_squads.team_id = teams.id
		WHERE
		    series.id = $1
		GROUP BY
		    %s
	`,
		seriesHeaderQuery.withClause,
		seriesHeaderQuery.selectFields, seriesHeaderQuery.joins, seriesHeaderQuery.groupByFields)

	err := db.QueryRow(ctx, query, seriesId).Scan(
		&header.SeriesId, &header.SeriesName, &header.Season, &header.TournamentId, &header.TournamentName,
		&header.TopBatters, &header.TopBowlers,
		&seriesWithSquadsList.SquadsList,
	)

	if err != nil {
		return seriesWithSquadsList, err
	}

	return seriesWithSquadsList, err
}

func ReadSeriesSingleSquadById(ctx context.Context, db DB_Exec, seriesId, squadId int64) (responses.SingleSeriesSingleSquad, error) {
	var seriesWithSquad responses.SingleSeriesSingleSquad
	header := &seriesWithSquad.SeriesHeader

	query := fmt.Sprintf(`
		%s,

		squads_list AS (
			SELECT series_squads.id, series_squads.squad_label
			FROM series_squads
			WHERE series_squads.series_id = $1
		) 

		SELECT 
			%s,
			-- order is crucial for struct unpacking
			( SELECT ARRAY_AGG ( squads_list.* ) FROM squads_list ) AS squads_list,

			ARRAY_AGG ( ROW ( 
				series_squad_entries.player_id, players.name, players.playing_role,
				players.date_of_birth, players.is_rhb, players.primary_bowling_style,
				series_squad_entries.is_captain, series_squad_entries.is_vice_captain, series_squad_entries.is_wk
			) )
		FROM series
		   	%s
		    LEFT JOIN series_squads ON series_squads.series_id = series.id 
			LEFT JOIN series_squad_entries ON series_squad_entries.squad_id = series_squads.id 
			LEFT JOIN players ON series_squad_entries.player_id = players.id
		WHERE
		    series.id = $1 AND series_squad_entries.squad_id = $2
		GROUP BY
		    %s
	`,
		seriesHeaderQuery.withClause,
		seriesHeaderQuery.selectFields, seriesHeaderQuery.joins, seriesHeaderQuery.groupByFields)

	err := db.QueryRow(ctx, query, seriesId, squadId).Scan(
		&header.SeriesId, &header.SeriesName, &header.Season, &header.TournamentId, &header.TournamentName,
		&header.TopBatters, &header.TopBowlers,
		&seriesWithSquad.SquadsList, &seriesWithSquad.Players,
	)

	if err != nil {
		return seriesWithSquad, err
	}

	return seriesWithSquad, err
}
