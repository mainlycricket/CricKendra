package statqueries

import (
	"fmt"
	"net/url"
)

func Query_Stat_Filter_Options(params *url.Values) (string, []any, error) {
	sqlWhere := newSqlWhere(-1, -1)

	if params != nil && len(*params) != 0 {
		sqlWhere.matchQuery.isMale(params.Get("is_male"))
		sqlWhere.matchQuery.playingFormat(params.Get("playing_format"))
	}

	sqlWhere.matchQuery.fields = append(sqlWhere.matchQuery.fields, "matches.ground_id")

	matchQuery := sqlWhere.matchQuery.prepareQuery()

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		), unique_teams AS (
			SELECT DISTINCT combined_teams.team_id
			FROM (
				SELECT matches.team1_id AS team_id FROM matches
				UNION
				SELECT matches.team2_id AS team_id FROM matches
			) combined_teams
		)

		SELECT ARRAY_AGG(DISTINCT ROW(unique_teams.team_id, teams.name)) AS teams,

			ARRAY_AGG(DISTINCT ROW(hn.id, hn.name))
			FILTER (WHERE hn.id IS NOT NULL AND hn.name IS NOT NULL) AS host_nations,

			ARRAY_AGG(DISTINCT ROW(continents.id, continents.name))
			FILTER (WHERE continents.id IS NOT NULL AND continents.name IS NOT NULL) AS continents,
			
			ARRAY_AGG(DISTINCT ROW(grounds.id, grounds.name, cities.name, hn.name))
			FILTER (WHERE grounds.id IS NOT NULL AND grounds.name IS NOT NULL) AS grounds,

			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			ARRAY_AGG(DISTINCT matches.season) AS seasons,

			ARRAY_AGG(DISTINCT ROW(series.id, series.name, series.season))
			FILTER (WHERE series.id IS NOT NULL AND series.name IS NOT NULL) AS series,

			ARRAY_AGG(DISTINCT ROW(tournaments.id, tournaments.name))
			FILTER (WHERE tournaments.id IS NOT NULL AND tournaments.name IS NOT NULL) AS tournaments
		
		FROM matches

		JOIN unique_teams ON TRUE
		LEFT JOIN teams ON unique_teams.team_id = teams.id

		LEFT JOIN grounds ON matches.ground_id = grounds.id
		LEFT JOIN cities ON cities.id = grounds.city_id
		LEFT JOIN host_nations hn ON hn.id = cities.host_nation_id
		LEFT JOIN continents ON continents.id = hn.continent_id

		LEFT JOIN match_series_entries mse ON mse.match_id = matches.id
		LEFT JOIN series ON mse.series_id = series.id AND (series.tour_flag IS NULL OR series.tour_flag != 'tour_series')
		LEFT JOIN tournaments ON series.tournament_id = tournaments.id
	`, matchQuery)

	return query, sqlWhere.matchQuery.args, nil
}
