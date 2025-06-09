package dbutils

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5"
	"github.com/mainlycricket/CricKendra/backend/internal/models"
	"github.com/mainlycricket/CricKendra/backend/internal/responses"
	"github.com/mainlycricket/CricKendra/backend/pkg/pgxutils"
)

func InsertHostNation(ctx context.Context, db DB_Exec, host_nation *models.HostNation) (int64, error) {
	var id int64

	query := `INSERT INTO host_nations (name, continent_id) VALUES($1, $2) RETURNING id`

	err := db.QueryRow(ctx, query, host_nation.Name, host_nation.ContinentId).Scan(&id)

	return id, err
}

func ReadHostNations(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.AllHostNationsResponse, error) {
	var response responses.AllHostNationsResponse

	queryInfoInput := pgxutils.QueryInfoInput{
		UrlQuery:     queryMap,
		TableName:    "host_nations",
		DefaultLimit: 50,
		DefaultSort:  []string{"name"},
	}

	queryInfoOutput, err := pgxutils.ParseQuery[models.HostNation](queryInfoInput)
	if err != nil {
		return response, err
	}

	query := fmt.Sprintf(`SELECT host_nations.id, host_nations.name, host_nations.continent_id, continents.name FROM host_nations LEFT JOIN continents ON host_nations.continent_id = continents.id %s %s %s`, queryInfoOutput.WhereClause, queryInfoOutput.OrderByClause, queryInfoOutput.PaginationClause)

	rows, err := db.Query(ctx, query, queryInfoOutput.Args...)
	if err != nil {
		return response, err
	}

	hostNations, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.AllHostNations, error) {
		var hostNation responses.AllHostNations
		err := rows.Scan(&hostNation.Id, &hostNation.Name, &hostNation.ContinetId, &hostNation.ContinentName)
		return hostNation, err
	})

	if len(hostNations) > queryInfoOutput.RecordsCount {
		response.HostNations = hostNations[:queryInfoOutput.RecordsCount]
		response.Next = true
	} else {
		response.HostNations = hostNations
		response.Next = false
	}

	return response, err
}
