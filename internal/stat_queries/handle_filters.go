package statqueries

import (
	"fmt"
	"net/url"
	"strings"
)

const (
	batting_stats = iota
	bowling_stats
	team_stats
)

type sqlWhere struct {
	conditions []string
	args       []any
}

func (sqlWhere *sqlWhere) applyFilters(params *url.Values, stats_type int) {
	sqlWhere.playingFormat(params)
	sqlWhere.isMale(params)
	sqlWhere.minStartDate(params)
	sqlWhere.maxStartDate(params)
	sqlWhere.season(params)
	sqlWhere.primaryTeam(params, stats_type)
	sqlWhere.oppositionTeam(params, stats_type)
	sqlWhere.ground(params)
}

func (sqlWhere *sqlWhere) getConditionString(prefix string) string {
	var condition string

	if len(sqlWhere.conditions) > 0 {
		condition = prefix + strings.Join(sqlWhere.conditions, " AND ")
	}

	return condition
}

func (sqlWhere *sqlWhere) playingFormat(params *url.Values) {
	playing_format := params.Get("playing_format")
	if playing_format != "" {
		sqlWhere.args = append(sqlWhere.args, playing_format)
		sqlWhere.conditions = append(sqlWhere.conditions, fmt.Sprintf(`matches.playing_format = $%d`, len(sqlWhere.args)))
	}
}

func (sqlWhere *sqlWhere) isMale(params *url.Values) {
	is_male := params.Get("is_male")
	if is_male != "" {
		sqlWhere.args = append(sqlWhere.args, is_male)
		sqlWhere.conditions = append(sqlWhere.conditions, fmt.Sprintf(`matches.is_male::bool = $%d`, len(sqlWhere.args)))
	}
}

func (sqlWhere *sqlWhere) minStartDate(params *url.Values) {
	min_start_date := params.Get("min_start_date")
	if min_start_date != "" {
		sqlWhere.args = append(sqlWhere.args, min_start_date)
		sqlWhere.conditions = append(sqlWhere.conditions, fmt.Sprintf(`matches.start_date::date >= $%d`, len(sqlWhere.args)))
	}
}

func (sqlWhere *sqlWhere) maxStartDate(params *url.Values) {
	max_start_date := params.Get("max_start_date")
	if max_start_date != "" {
		sqlWhere.args = append(sqlWhere.args, max_start_date)
		sqlWhere.conditions = append(sqlWhere.conditions, fmt.Sprintf(`matches.start_date::date <= $%d`, len(sqlWhere.args)))
	}
}

func (sqlWhere *sqlWhere) season(params *url.Values) {
	var placeholders []string

	seasons := (*params)["season"]
	for _, season := range seasons {
		sqlWhere.args = append(sqlWhere.args, season)
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, len(sqlWhere.args)))
	}

	if len(placeholders) > 0 {
		sqlWhere.conditions = append(sqlWhere.conditions, fmt.Sprintf(`matches.season IN (%s)`, strings.Join(placeholders, ", ")))
	}
}

func (sqlWhere *sqlWhere) primaryTeam(params *url.Values, statsType int) {
	var placeholders []string

	teams_id := (*params)["primary_team"]
	for _, team_id := range teams_id {
		sqlWhere.args = append(sqlWhere.args, team_id)
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, len(sqlWhere.args)))
	}

	if len(placeholders) > 0 {
		switch statsType {
		case batting_stats:
			sqlWhere.conditions = append(sqlWhere.conditions, fmt.Sprintf(`innings.batting_team_id IN (%s)`, strings.Join(placeholders, ", ")))
		case bowling_stats:
			sqlWhere.conditions = append(sqlWhere.conditions, fmt.Sprintf(`innings.bowling_team_id IN (%s)`, strings.Join(placeholders, ", ")))
		case team_stats:
			sqlWhere.conditions = append(sqlWhere.conditions, fmt.Sprintf(`matches.team1_id IN (%s)`, strings.Join(placeholders, ", ")))
			sqlWhere.conditions = append(sqlWhere.conditions, fmt.Sprintf(`matches.team2_id IN (%s)`, strings.Join(placeholders, ", ")))
		}
	}
}

func (sqlWhere *sqlWhere) oppositionTeam(params *url.Values, statsType int) {
	var placeholders []string

	teams_id := (*params)["opposition_team"]
	for _, team_id := range teams_id {
		sqlWhere.args = append(sqlWhere.args, team_id)
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, len(sqlWhere.args)))
	}

	if len(placeholders) > 0 {
		switch statsType {
		case batting_stats:
			sqlWhere.conditions = append(sqlWhere.conditions, fmt.Sprintf(`innings.bowling_team_Id IN (%s)`, strings.Join(placeholders, ", ")))
		case bowling_stats:
			sqlWhere.conditions = append(sqlWhere.conditions, fmt.Sprintf(`innings.batting_team_id IN (%s)`, strings.Join(placeholders, ", ")))
		case team_stats:
			sqlWhere.conditions = append(sqlWhere.conditions, fmt.Sprintf(`matches.team1_id IN (%s)`, strings.Join(placeholders, ", ")))
			sqlWhere.conditions = append(sqlWhere.conditions, fmt.Sprintf(`matches.team2_id IN (%s)`, strings.Join(placeholders, ", ")))
		}
	}
}

func (sqlWhere *sqlWhere) ground(params *url.Values) {
	var placeholders []string

	grounds_id := (*params)["ground"]
	for _, ground_id := range grounds_id {
		sqlWhere.args = append(sqlWhere.args, ground_id)
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, len(sqlWhere.args)))
	}

	if len(placeholders) > 0 {
		sqlWhere.conditions = append(sqlWhere.conditions, fmt.Sprintf(`matches.ground_id IN (%s)`, strings.Join(placeholders, ", ")))
	}
}
