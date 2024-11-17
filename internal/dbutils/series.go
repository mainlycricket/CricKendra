package dbutils

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
	"github.com/mainlycricket/CricKendra/pkg/pgxutils"
)

func InsertSeries(ctx context.Context, db DB_Exec, series *models.Series) error {
	query := `INSERT INTO series (name, is_male, playing_level, playing_format, season, teams_id, host_nations_id, tournament_id, parent_series_id, tour_id, players_of_the_series_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	cmd, err := db.Exec(ctx, query, series.Name, series.IsMale, series.PlayingLevel, series.PlayingFormat, series.Season, series.TeamsId, series.HostNationsId, series.TournamentId, series.ParentSeriesId, series.TourId, series.PoTSsId)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert series")
	}

	return nil
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

	query := fmt.Sprintf(`SELECT id, name, is_male, playing_level, playing_format, season FROM series %s %s %s`, queryInfoOutput.WhereClause, queryInfoOutput.OrderByClause, queryInfoOutput.PaginationClause)

	rows, err := db.Query(ctx, query, queryInfoOutput.Args...)
	if err != nil {
		return response, err
	}

	seriesList, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.AllSeries, error) {
		var series responses.AllSeries

		err := rows.Scan(&series.Id, &series.Name, &series.IsMale, &series.PlayingLevel, &series.PlayingFormat, &series.Season)

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
