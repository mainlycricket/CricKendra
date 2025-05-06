package dbutils

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/mainlycricket/CricKendra/internal/responses"
	statqueries "github.com/mainlycricket/CricKendra/internal/stat_queries"
)

func Read_Stat_Filter_Options(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsFilters, error) {
	var statFilters responses.StatsFilters

	query, args, err := statqueries.Query_Stat_Filter_Options(&queryMap)
	if err != nil {
		return statFilters, fmt.Errorf(`failed to parse query: %v`, err)
	}

	err = db.QueryRow(ctx, query, args...).Scan(&statFilters.PrimaryTeams, &statFilters.OppositionTeams, &statFilters.HostNations, &statFilters.Continents, &statFilters.Grounds, &statFilters.MinDate, &statFilters.MaxDate, &statFilters.Seasons, &statFilters.Series, &statFilters.Tournaments)

	if err != nil {
		log.Println(query)
		return statFilters, err
	}

	return statFilters, nil
}
