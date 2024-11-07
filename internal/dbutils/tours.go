package dbutils

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
	"github.com/mainlycricket/CricKendra/pkg/pgxutils"
)

func InsertTour(ctx context.Context, db *pgxpool.Pool, tour *models.Tour) error {
	query := `INSERT INTO tours (touring_team_id, host_nations_id, season) VALUES($1, $2, $3)`

	cmd, err := db.Exec(ctx, query, tour.TouringTeamId, tour.HostNationsId, tour.Season)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert tour")
	}

	return nil
}

func ReadTours(ctx context.Context, db *pgxpool.Pool, queryMap url.Values) (responses.AllToursResponse, error) {
	var response responses.AllToursResponse

	queryInfoInput := pgxutils.QueryInfoInput{
		UrlQuery:     queryMap,
		TableName:    "tours",
		DefaultLimit: 20,
		DefaultSort:  []string{"-season"},
	}

	queryInfoOutput, err := pgxutils.ParseQuery[models.Tour](queryInfoInput)
	if err != nil {
		return response, err
	}

	query := fmt.Sprintf(`SELECT 
		tours.id, tours.touring_team_id, teams.name AS touring_team_name, 
		COALESCE(
			jsonb_agg(
				jsonb_build_object('id', host_team.id, 'name', host_team.name)
			) FILTER (WHERE host_team.id IS NOT NULL),
			 '[]'
			) AS host_nations,
		tours.season 
		FROM tours 
		LEFT JOIN teams ON tours.touring_team_id = teams.id 
		LEFT JOIN LATERAL unnest(tours.host_nations_id) AS host_id ON true 
		LEFT JOIN teams AS host_team ON host_team.id = host_id
		%s 
		GROUP BY tours.id, teams.name 
		%s 
		%s`, queryInfoOutput.WhereClause, queryInfoOutput.OrderByClause, queryInfoOutput.PaginationClause)

	rows, err := db.Query(ctx, query, queryInfoOutput.Args...)
	if err != nil {
		return response, err
	}

	tours, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.AllTours, error) {
		var tour responses.AllTours

		err := rows.Scan(&tour.Id, &tour.TouringTeamId, &tour.TouringTeamName, &tour.HostNations, &tour.Season)

		return tour, err
	})

	if len(tours) > queryInfoOutput.RecordsCount {
		response.Tours = tours[:queryInfoOutput.RecordsCount]
		response.Next = true
	} else {
		response.Tours = tours
		response.Next = false
	}

	return response, err
}
