package statqueries

import (
	"fmt"
	"net/url"
	"strings"
)

const match_grounds_join string = `JOIN grounds ON matches.ground_id = grounds.id`
const match_cities_join string = `JOIN cities ON grounds.city_id = cities.id`
const match_hostNations_join string = `JOIN host_nations ON cities.host_nation_id = host_nations.id`
const match_continents_join string = `JOIN continents ON host_nations.continent_id = continents.id`
const match_seriesEntires_join string = `JOIN match_series_entries ON match_series_entries.match_id = matches.id`
const match_series_join string = `JOIN series ON match_series_entries.series_id = series.id`
const match_tournaments_join string = `JOIN tournaments ON series.tournament_id = tournaments.id`

type matchQuery struct {
	tableQuery
}

func newMatchQuery() *matchQuery {
	matchQuery := &matchQuery{}

	matchQuery.tableName = "matches"

	matchQuery.fields = []string{
		"matches.id", "matches.start_date", "matches.season",
		"matches.team1_id", "matches.team2_id",
		"matches.home_team_id", "matches.away_team_id", "matches.is_neutral_venue",
		"matches.toss_winner_team_id", "matches.toss_loser_team_id", "matches.is_toss_decision_bat",
		"matches.match_winner_team_id", "matches.match_loser_team_id",
		"matches.final_result",
	}

	matchQuery.joins = map[string]int{}

	return matchQuery
}

/*
Playing Format
Is Male
Min Start Date
Max Start Date
Season
Primary Team
Opposition Team
Home Team, Away Team, Neutral
Continent
Host Nation
Ground
Series
Tournament
*/

func (mq *matchQuery) playingFormat(params *url.Values) {
	if playing_format := params.Get("playing_format"); playing_format != "" {
		mq.args = append(mq.args, playing_format)
		condition := fmt.Sprintf(`matches.playing_format = $%d`, len(mq.args))
		mq.conditions = append(mq.conditions, condition)
	}
}

func (mq *matchQuery) isMale(params *url.Values) {
	if is_male := params.Get("is_male"); is_male != "" {
		mq.args = append(mq.args, is_male)
		condition := fmt.Sprintf(`matches.is_male::bool = $%d`, len(mq.args))
		mq.conditions = append(mq.conditions, condition)
	}
}

func (mq *matchQuery) minStartDate(params *url.Values) {
	if min_start_date := params.Get("min_start_date"); min_start_date != "" {
		mq.args = append(mq.args, min_start_date)
		condition := fmt.Sprintf(`matches.start_date::date >= $%d`, len(mq.args))
		mq.conditions = append(mq.conditions, condition)
	}
}

func (mq *matchQuery) maxStartDate(params *url.Values) {
	if max_start_date := params.Get("max_start_date"); max_start_date != "" {
		mq.args = append(mq.args, max_start_date)
		condition := fmt.Sprintf(`matches.start_date::date <= $%d`, len(mq.args))
		mq.conditions = append(mq.conditions, condition)
	}
}

func (mq *matchQuery) season(params *url.Values) {
	var placeholders []string

	seasons := (*params)["season"]
	for _, season := range seasons {
		mq.args = append(mq.args, season)
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, len(mq.args)))
	}

	if len(placeholders) > 0 {
		condition := fmt.Sprintf(`matches.season IN (%s)`, strings.Join(placeholders, ", "))
		mq.conditions = append(mq.conditions, condition)
	}
}

func (mq *matchQuery) primaryTeam(params *url.Values, statsType int, inningsFilters *inningsFilters) {
	var placeholders []string

	teams_id := (*params)["primary_team"]
	for _, team_id := range teams_id {
		mq.args = append(mq.args, team_id)
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, len(mq.args)))
	}

	if len(placeholders) > 0 {
		placeholderString := strings.Join(placeholders, ", ")

		mq.conditions = append(mq.conditions, fmt.Sprintf(`(matches.team1_id IN (%s) OR matches.team2_id IN (%s))`, placeholderString, placeholderString))

		switch statsType {
		case batting_stats:
			inningsFilters.setBattingTeams(placeholderString)
		case bowling_stats:
			inningsFilters.setBowlingTeams(placeholderString)
		case team_stats:
			team_total_for, _ := teamTotalForAgainst(params.Get("team_total_for"))
			if team_total_for == "batting_team_id" {
				inningsFilters.setBattingTeams(placeholderString)
			} else {
				inningsFilters.setBowlingTeams(placeholderString)
			}
		}
	}
}

func (mq *matchQuery) oppositionTeam(params *url.Values, statsType int, inningsFilters *inningsFilters) {
	var placeholders []string

	teams_id := (*params)["opposition_team"]
	for _, team_id := range teams_id {
		mq.args = append(mq.args, team_id)
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, len(mq.args)))
	}

	if len(placeholders) > 0 {
		placeholderString := strings.Join(placeholders, ", ")

		mq.conditions = append(mq.conditions, fmt.Sprintf(`(matches.team1_id IN (%s) OR matches.team2_id IN (%s))`, placeholderString, placeholderString))

		switch statsType {
		case batting_stats:
			inningsFilters.setBowlingTeams(placeholderString)
		case bowling_stats:
			inningsFilters.setBattingTeams(placeholderString)
		case team_stats:
			team_total_for, _ := teamTotalForAgainst(params.Get("team_total_for"))
			if team_total_for == "batting_team_id" {
				inningsFilters.setBowlingTeams(placeholderString)
			} else {
				inningsFilters.setBattingTeams(placeholderString)
			}
		}
	}
}

func (mq *matchQuery) homeAwayTeam(params *url.Values, statsType int, inningsFilters *inningsFilters) {
	var isHome, isAway, isNeutral bool
	for _, value := range (*params)["home_or_away"] {
		switch value {
		case "home":
			isHome = true
		case "away":
			isAway = true
		case "neutral":
			isNeutral = true
		}
	}

	var team_total_for string
	if statsType == team_stats {
		team_total_for, _ = teamTotalForAgainst(params.Get("team_total_for"))
	}

	if isHome {
		switch statsType {
		case batting_stats:
			inningsFilters.setHomeAway(true, true)
		case bowling_stats:
			inningsFilters.setHomeAway(true, false)
		case team_stats:
			if team_total_for == "batting_team_id" {
				inningsFilters.setHomeAway(true, true)
			} else {
				inningsFilters.setHomeAway(true, false)
			}
		}
	}

	if isAway {
		switch statsType {
		case batting_stats:
			inningsFilters.setHomeAway(false, true)
		case bowling_stats:
			inningsFilters.setHomeAway(false, false)
		case team_stats:
			if team_total_for == "batting_team_id" {
				inningsFilters.setHomeAway(false, true)
			} else {
				inningsFilters.setHomeAway(false, false)
			}
		}
	}

	if isNeutral {
		mq.conditions = append(mq.conditions, "matches.is_neutral_venue = TRUE")
	}
}

func (mq *matchQuery) tossResult(params *url.Values, statsType int, inningsFilters *inningsFilters) {
	switch params.Get("toss_result") {
	case "won":
		switch statsType {
		case batting_stats:
			inningsFilters.setTossResult(true, true)
		case bowling_stats:
			inningsFilters.setTossResult(true, false)
		case team_stats:
			team_total_for, _ := teamTotalForAgainst(params.Get("team_total_for"))
			if team_total_for == "batting_team_id" {
				inningsFilters.setTossResult(true, true)
			} else {
				inningsFilters.setTossResult(true, false)
			}
		}
	case "lost":
		switch statsType {
		case batting_stats:
			inningsFilters.setTossResult(false, true)
		case bowling_stats:
			inningsFilters.setTossResult(false, false)
		case team_stats:
			team_total_for, _ := teamTotalForAgainst(params.Get("team_total_for"))
			if team_total_for == "batting_team_id" {
				inningsFilters.setTossResult(false, true)
			} else {
				inningsFilters.setTossResult(false, false)
			}
		}
	}
}

func (mq *matchQuery) batFieldFirst(params *url.Values, statsType int, inningsFilters *inningsFilters) {
	switch params.Get("bat_field_first") {
	case "bat":
		switch statsType {
		case batting_stats:
			inningsFilters.setBatFieldFirst(true, true)
		case bowling_stats:
			inningsFilters.setBatFieldFirst(true, false)
		case team_stats:
			team_total_for, _ := teamTotalForAgainst(params.Get("team_total_for"))
			if team_total_for == "batting_team_id" {
				inningsFilters.setBatFieldFirst(true, true)
			} else {
				inningsFilters.setBatFieldFirst(true, false)
			}
		}
	case "field":
		switch statsType {
		case batting_stats:
			inningsFilters.setBatFieldFirst(false, true)
		case bowling_stats:
			inningsFilters.setBatFieldFirst(false, false)
		case team_stats:
			team_total_for, _ := teamTotalForAgainst(params.Get("team_total_for"))
			if team_total_for == "batting_team_id" {
				inningsFilters.setBatFieldFirst(false, true)
			} else {
				inningsFilters.setBatFieldFirst(false, false)
			}
		}
	}
}

func (mq *matchQuery) continent(params *url.Values) {
	var placeholders []string

	continentsId := (*params)["continent"]
	for _, continentId := range continentsId {
		mq.args = append(mq.args, continentId)
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, len(mq.args)))
	}

	if len(placeholders) > 0 {
		mq.joins[match_grounds_join] = 0
		mq.joins[match_cities_join] = 1
		mq.joins[match_hostNations_join] = 2
		mq.joins[match_continents_join] = 3

		mq.conditions = append(mq.conditions, fmt.Sprintf(`continents.id IN (%s)`, strings.Join(placeholders, ", ")))
	}
}

func (mq *matchQuery) hostNation(params *url.Values) {
	var placeholders []string

	hostNationsId := (*params)["host_nation"]
	for _, hostNationId := range hostNationsId {
		mq.args = append(mq.args, hostNationId)
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, len(mq.args)))
	}

	if len(placeholders) > 0 {
		mq.joins[match_grounds_join] = 0
		mq.joins[match_cities_join] = 1
		mq.joins[match_hostNations_join] = 2

		mq.conditions = append(mq.conditions, fmt.Sprintf(`host_nations.id IN (%s)`, strings.Join(placeholders, ", ")))
	}
}

func (mq *matchQuery) ground(params *url.Values) {
	var placeholders []string

	grounds_id := (*params)["ground"]
	for _, ground_id := range grounds_id {
		mq.args = append(mq.args, ground_id)
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, len(mq.args)))
	}

	if len(placeholders) > 0 {
		mq.joins[match_grounds_join] = 0

		mq.conditions = append(mq.conditions, fmt.Sprintf(`grounds.id IN (%s)`, strings.Join(placeholders, ", ")))
	}
}

func (mq *matchQuery) series(params *url.Values) {
	var placeholders []string

	series_list_id := (*params)["series"]
	for _, series_id := range series_list_id {
		mq.args = append(mq.args, series_id)
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, len(mq.args)))
	}

	if len(placeholders) > 0 {
		mq.joins[match_seriesEntires_join] = len(mq.joins)
		mq.conditions = append(mq.conditions, fmt.Sprintf(`match_series_entries.series_id IN (%s)`, strings.Join(placeholders, ", ")))
		mq.conditions = append(mq.conditions, `(series.tour_flag IS NULL OR series.tour_flag != 'tour_series')`)
	}
}

func (mq *matchQuery) tournament(params *url.Values) {
	var placeholders []string

	tournaments := (*params)["tournament"]
	for _, tournament := range tournaments {
		mq.args = append(mq.args, tournament)
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, len(mq.args)))
	}

	if len(placeholders) > 0 {
		if _, ok := mq.joins[match_seriesEntires_join]; !ok {
			mq.joins[match_seriesEntires_join] = len(mq.joins)
		}
		mq.joins[match_series_join] = len(mq.joins)
		mq.joins[match_tournaments_join] = len(mq.joins)

		mq.conditions = append(mq.conditions, fmt.Sprintf(`tournaments.id IN (%s)`, strings.Join(placeholders, ", ")))
	}
}

func (mq *matchQuery) ensureGround() {
	if _, ok := mq.joins[match_grounds_join]; !ok {
		mq.joins[match_grounds_join] = len(mq.joins)
	}
	mq.fields = append(mq.fields, "matches.ground_id", "grounds.name AS ground_name")
}

func (mq *matchQuery) ensureCity() {
	if _, ok := mq.joins[match_grounds_join]; !ok {
		mq.joins[match_grounds_join] = len(mq.joins)
	}

	if _, ok := mq.joins[match_cities_join]; !ok {
		mq.joins[match_cities_join] = len(mq.joins)
	}

	mq.fields = append(mq.fields, "grounds.city_id", "cities.name AS city_name")
}

func (mq *matchQuery) ensureHostNation() {
	if _, ok := mq.joins[match_grounds_join]; !ok {
		mq.joins[match_grounds_join] = len(mq.joins)
	}

	if _, ok := mq.joins[match_cities_join]; !ok {
		mq.joins[match_cities_join] = len(mq.joins)
	}

	if _, ok := mq.joins[match_hostNations_join]; !ok {
		mq.joins[match_hostNations_join] = len(mq.joins)
	}

	mq.fields = append(mq.fields, "cities.host_nation_id", "host_nations.name AS host_nation_name")
}

func (mq *matchQuery) ensureContinent() {
	if _, ok := mq.joins[match_grounds_join]; !ok {
		mq.joins[match_grounds_join] = len(mq.joins)
	}

	if _, ok := mq.joins[match_cities_join]; !ok {
		mq.joins[match_cities_join] = len(mq.joins)
	}

	if _, ok := mq.joins[match_hostNations_join]; !ok {
		mq.joins[match_hostNations_join] = len(mq.joins)
	}

	if _, ok := mq.joins[match_continents_join]; !ok {
		mq.joins[match_continents_join] = len(mq.joins)
	}

	mq.fields = append(mq.fields, "host_nations.continent_id", "continents.name AS continent_name")
}

func (mq *matchQuery) ensureSeries() {
	if _, ok := mq.joins[match_seriesEntires_join]; !ok {
		mq.joins[match_seriesEntires_join] = len(mq.joins)
	}

	if _, ok := mq.joins[match_series_join]; !ok {
		mq.joins[match_series_join] = len(mq.joins)
		mq.conditions = append(mq.conditions, `(series.tour_flag IS NULL OR series.tour_flag != 'tour_series')`)
	}

	mq.fields = append(mq.fields, "series.id AS series_id", "series.name AS series_name", "series.season AS series_season")
}

func (mq *matchQuery) ensureTournament() {
	if _, ok := mq.joins[match_seriesEntires_join]; !ok {
		mq.joins[match_seriesEntires_join] = len(mq.joins)
	}

	if _, ok := mq.joins[match_series_join]; !ok {
		mq.joins[match_series_join] = len(mq.joins)
	}

	if _, ok := mq.joins[match_tournaments_join]; !ok {
		mq.joins[match_tournaments_join] = len(mq.joins)
	}

	mq.fields = append(mq.fields, "series.tournament_id", "tournaments.name AS tournament_name")
}
