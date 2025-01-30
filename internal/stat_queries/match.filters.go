package statqueries

import (
	"fmt"
	"net/url"
	"slices"
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

func (mq *matchQuery) applyMatchFilters(params *url.Values) {
	if params == nil || len(*params) == 0 {
		return
	}

	mq.playingFormat(params.Get("playing_format"))
	mq.isMale(params.Get("is_male"))

	mq.minStartDate(params.Get("min_start_date"))
	mq.maxStartDate(params.Get("max_start_date"))
	mq.season((*params)["season"])

	mq.primaryTeam((*params)["primary_team"])
	mq.oppositionTeam((*params)["opposition_team"])
	mq.setMatchResult((*params)["match_result"])
	mq.isNeutralVenue((*params)["home_or_away"])

	mq.continent((*params)["continent"])
	mq.hostNation((*params)["host_nation"])
	mq.ground((*params)["ground"])

	mq.series((*params)["series"])
	mq.tournament((*params)["tournament"])
}

func (mq *matchQuery) playingFormat(playing_format string) {
	if playing_format != "" {
		var placeholders []string
		var otherFormat string

		mq.args = append(mq.args, playing_format)
		mq.placeholdersCount++
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, mq.placeholdersCount))

		switch playing_format {
		case "first_class":
			otherFormat = "Test"
		case "list_a":
			otherFormat = "ODI"
		case "T20":
			otherFormat = "T20I"
		}

		if otherFormat != "" {
			mq.args = append(mq.args, otherFormat)
			mq.placeholdersCount++
			placeholders = append(placeholders, fmt.Sprintf(`$%d`, mq.placeholdersCount))
		}

		condition := fmt.Sprintf(`matches.playing_format IN (%s)`, strings.Join(placeholders, ","))
		mq.conditions = append(mq.conditions, condition)
	}
}

func (mq *matchQuery) isMale(is_male string) {
	if is_male != "" {
		mq.args = append(mq.args, is_male)
		mq.placeholdersCount++
		condition := fmt.Sprintf(`matches.is_male::bool = $%d`, mq.placeholdersCount)
		mq.conditions = append(mq.conditions, condition)
	}
}

func (mq *matchQuery) minStartDate(min_start_date string) {
	if min_start_date != "" {
		mq.args = append(mq.args, min_start_date)
		mq.placeholdersCount++
		condition := fmt.Sprintf(`matches.start_date::date >= $%d`, mq.placeholdersCount)
		mq.conditions = append(mq.conditions, condition)
	}
}

func (mq *matchQuery) maxStartDate(max_start_date string) {
	if max_start_date != "" {
		mq.args = append(mq.args, max_start_date)
		mq.placeholdersCount++
		condition := fmt.Sprintf(`matches.start_date::date <= $%d`, mq.placeholdersCount)
		mq.conditions = append(mq.conditions, condition)
	}
}

func (mq *matchQuery) season(seasons []string) {
	var placeholders []string

	for _, season := range seasons {
		mq.args = append(mq.args, season)
		mq.placeholdersCount++
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, mq.placeholdersCount))
	}

	if len(placeholders) > 0 {
		condition := fmt.Sprintf(`matches.season IN (%s)`, strings.Join(placeholders, ", "))
		mq.conditions = append(mq.conditions, condition)
	}
}

func (mq *matchQuery) primaryTeam(teams_id []string) {
	var placeholders []string

	for _, team_id := range teams_id {
		mq.args = append(mq.args, team_id)
		mq.placeholdersCount++
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, mq.placeholdersCount))
	}

	if len(placeholders) > 0 {
		placeholderString := strings.Join(placeholders, ", ")
		mq.conditions = append(mq.conditions, fmt.Sprintf(`(matches.team1_id IN (%s) OR matches.team2_id IN (%s))`, placeholderString, placeholderString))
	}
}

func (mq *matchQuery) oppositionTeam(teams_id []string) {
	var placeholders []string

	for _, team_id := range teams_id {
		mq.args = append(mq.args, team_id)
		mq.placeholdersCount++
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, mq.placeholdersCount))
	}

	if len(placeholders) > 0 {
		placeholderString := strings.Join(placeholders, ", ")
		mq.conditions = append(mq.conditions, fmt.Sprintf(`(matches.team1_id IN (%s) OR matches.team2_id IN (%s))`, placeholderString, placeholderString))
	}
}

func (mq *matchQuery) isNeutralVenue(values []string) {
	if slices.Contains(values, "neutral") {
		mq.conditions = append(mq.conditions, "matches.is_neutral_venue = TRUE")
	}
}

func (mq *matchQuery) setMatchResult(values []string) {
	if slices.Contains(values, "won") || slices.Contains(values, "lost") {
		mq.conditions = append(mq.conditions, `matches.final_result = 'winner decided'`)
	}

	if slices.Contains(values, "tied") {
		mq.conditions = append(mq.conditions, `matches.final_result = 'tie'`)
	}

	if slices.Contains(values, "drawn") {
		mq.conditions = append(mq.conditions, `matches.final_result = 'draw'`)
	}

	if slices.Contains(values, "no_result") {
		mq.conditions = append(mq.conditions, `matches.final_result = 'no result'`)
	}
}

func (mq *matchQuery) continent(continentsId []string) {
	var placeholders []string

	for _, continentId := range continentsId {
		mq.args = append(mq.args, continentId)
		mq.placeholdersCount++
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, mq.placeholdersCount))
	}

	if len(placeholders) > 0 {
		mq.joins[match_grounds_join] = 0
		mq.joins[match_cities_join] = 1
		mq.joins[match_hostNations_join] = 2
		mq.joins[match_continents_join] = 3

		mq.conditions = append(mq.conditions, fmt.Sprintf(`continents.id IN (%s)`, strings.Join(placeholders, ", ")))
	}
}

func (mq *matchQuery) hostNation(hostNationsId []string) {
	var placeholders []string

	for _, hostNationId := range hostNationsId {
		mq.args = append(mq.args, hostNationId)
		mq.placeholdersCount++
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, mq.placeholdersCount))
	}

	if len(placeholders) > 0 {
		mq.joins[match_grounds_join] = 0
		mq.joins[match_cities_join] = 1
		mq.joins[match_hostNations_join] = 2

		mq.conditions = append(mq.conditions, fmt.Sprintf(`host_nations.id IN (%s)`, strings.Join(placeholders, ", ")))
	}
}

func (mq *matchQuery) ground(grounds_id []string) {
	var placeholders []string

	for _, ground_id := range grounds_id {
		mq.args = append(mq.args, ground_id)
		mq.placeholdersCount++
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, mq.placeholdersCount))
	}

	if len(placeholders) > 0 {
		mq.joins[match_grounds_join] = 0

		mq.conditions = append(mq.conditions, fmt.Sprintf(`grounds.id IN (%s)`, strings.Join(placeholders, ", ")))
	}
}

func (mq *matchQuery) series(series_list_id []string) {
	var placeholders []string

	for _, series_id := range series_list_id {
		mq.args = append(mq.args, series_id)
		mq.placeholdersCount++
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, mq.placeholdersCount))
	}

	if len(placeholders) > 0 {
		mq.joins[match_seriesEntires_join] = len(mq.joins)
		mq.joins[match_series_join] = len(mq.joins)
		mq.conditions = append(mq.conditions, fmt.Sprintf(`match_series_entries.series_id IN (%s)`, strings.Join(placeholders, ", ")))
		mq.conditions = append(mq.conditions, `(series.tour_flag IS NULL OR series.tour_flag != 'tour_series')`)
	}
}

func (mq *matchQuery) tournament(tournaments []string) {
	var placeholders []string

	for _, tournament := range tournaments {
		mq.args = append(mq.args, tournament)
		mq.placeholdersCount++
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, mq.placeholdersCount))
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
