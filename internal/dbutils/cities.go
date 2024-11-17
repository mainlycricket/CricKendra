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

func InsertCity(ctx context.Context, db DB_Exec, city *models.City) error {
	query := `INSERT INTO cities (name, host_nation_id) VALUES($1, $2)`

	cmd, err := db.Exec(ctx, query, city.Name, city.HostNationId)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert city")
	}

	return nil
}

func ReadCities(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.AllCitiesResponse, error) {
	var response responses.AllCitiesResponse

	queryInfoInput := pgxutils.QueryInfoInput{
		UrlQuery:     queryMap,
		TableName:    "cities",
		DefaultLimit: 50,
		DefaultSort:  []string{"name"},
	}

	queryInfoOutput, err := pgxutils.ParseQuery[models.City](queryInfoInput)
	if err != nil {
		return response, err
	}

	query := fmt.Sprintf(`SELECT cities.id, cities.name, cities.host_nation_id, host_nations.name, host_nations.continent_id, continents.name FROM cities LEFT JOIN host_nations ON cities.host_nation_id = host_nations.id LEFT JOIN continents ON host_nations.continent_id = continents.id %s %s %s`, queryInfoOutput.WhereClause, queryInfoOutput.OrderByClause, queryInfoOutput.PaginationClause)

	rows, err := db.Query(ctx, query, queryInfoOutput.Args...)
	if err != nil {
		return response, err
	}

	cities, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.AllCities, error) {
		var city responses.AllCities
		err := rows.Scan(&city.Id, &city.Name, &city.HostNationId, &city.HostNationName, &city.ContinetId, &city.ContinentName)
		return city, err
	})

	if len(cities) > queryInfoOutput.RecordsCount {
		response.Cities = cities[:queryInfoOutput.RecordsCount]
		response.Next = true
	} else {
		response.Cities = cities
		response.Next = false
	}

	return response, err
}
