package statqueries

import (
	"fmt"
	"net/url"
	"slices"
	"strings"
)

type inningsFilters struct {
	conditions        []string
	args              []any
	placeholdersCount int
}

func newInningsFilters() *inningsFilters {
	inningsFilters := &inningsFilters{}
	return inningsFilters
}

func (filters *inningsFilters) getClause() string {
	return prefixJoin(filters.conditions, "WHERE", " AND ")
}

func (filters *inningsFilters) applyInningsFilters(params *url.Values, stats_type int) {
	isBattingStats := stats_type == batting_stats
	if stats_type == team_stats {
		team_total_for, _ := teamTotalForAgainst(params.Get("team_total_for"))
		isBattingStats = team_total_for == "batting_team_id"
	}

	if isBattingStats {
		filters.setBattingTeams((*params)["primary_team"])
		filters.setBowlingTeams((*params)["opposition_team"])
	} else {
		filters.setBowlingTeams((*params)["primary_team"])
		filters.setBattingTeams((*params)["opposition_team"])
	}

	filters.setHomeAway((*params)["home_or_away"], isBattingStats)
	filters.setMatchResult((*params)["match_result"], isBattingStats)
	filters.setTossResult(params.Get("toss_result"), isBattingStats)
	filters.setBatFieldFirst(params.Get("bat_field_first"), isBattingStats)
	filters.setInningsNumber((*params)["innings_number"])

	switch stats_type {
	case batting_stats:
		filters.conditions = append(filters.conditions, "batting_scorecards.batting_position IS NOT NULL")
		filters.setBatterRunsRange(params.Get("min__innings_runs_scored"), params.Get("max__innings_runs_scored"))
		filters.setBatterPositionRange(params.Get("min__innings_batting_position"), params.Get("min__innings_batting_position"))
		filters.setBatterIsDismissed(params.Get("innings_is_batter_dismissed"))
		filters.setBatterDismissalTypes((*params)["innings_batter_dismissal_type"])
	case bowling_stats:
		filters.conditions = append(filters.conditions, "bowling_scorecards.bowling_position IS NOT NULL")
		filters.setBowlerBallsRange(params.Get("min__innings_balls_bowled"), params.Get("max__innings_balls_bowled"))
		filters.setBowlerRunsRange(params.Get("min__innings_runs_conceded"), params.Get("max__innings_runs_conceded"))
		filters.setBowlerWicketsRange(params.Get("min__innings_wickets_taken"), params.Get("max__innings_wickets_taken"))
		filters.setBowlerPositionRange(params.Get("min__innings_bowling_position"), params.Get("min__innings_bowling_position"))
	case team_stats:
		filters.setTeamInningsRunsRange(params.Get("min__team_innings_runs"), params.Get("max__team_innings_runs"))
		filters.setTeamInningsWktsRange(params.Get("min__team_innings_wickets"), params.Get("max__team_innings_wickets"))
		filters.setTeamInningsBallsRange(params.Get("min__team_innings_balls"), params.Get("max__team_innings_balls"))
	}
}

func (filters *inningsFilters) setBattingTeams(teams_id []string) {
	var placeholders []string

	for _, team_id := range teams_id {
		filters.args = append(filters.args, team_id)
		filters.placeholdersCount++
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, filters.placeholdersCount))
	}

	if len(placeholders) > 0 {
		placeholderString := strings.Join(placeholders, ", ")
		condition := fmt.Sprintf(`innings.batting_team_id IN (%s)`, placeholderString)
		filters.conditions = append(filters.conditions, condition)
	}
}

func (filters *inningsFilters) setBowlingTeams(teams_id []string) {
	var placeholders []string

	for _, team_id := range teams_id {
		filters.args = append(filters.args, team_id)
		filters.placeholdersCount++
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, filters.placeholdersCount))
	}

	if len(placeholders) > 0 {
		placeholderString := strings.Join(placeholders, ", ")
		condition := fmt.Sprintf(`innings.bowling_team_id IN (%s)`, placeholderString)
		filters.conditions = append(filters.conditions, condition)
	}
}

func (filters *inningsFilters) setHomeAway(values []string, isBattingTeam bool) {
	inningsField := "innings.bowling_team_id"
	if isBattingTeam {
		inningsField = "innings.batting_team_id"
	}

	if slices.Contains(values, "home") {
		condition := fmt.Sprintf(`matches.home_team_id = %s`, inningsField)
		filters.conditions = append(filters.conditions, condition)
	}

	if slices.Contains(values, "away") {
		condition := fmt.Sprintf(`matches.away_team_id = %s`, inningsField)
		filters.conditions = append(filters.conditions, condition)
	}
}

func (filters *inningsFilters) setMatchResult(values []string, isBattingTeam bool) {
	team_field := `innings.bowling_team_id`
	if isBattingTeam {
		team_field = `innings.batting_team_id`
	}

	if slices.Contains(values, "won") {
		filters.conditions = append(filters.conditions, fmt.Sprintf(`matches.match_winner_team_id = %s`, team_field))
	}

	if slices.Contains(values, "lost") {
		filters.conditions = append(filters.conditions, fmt.Sprintf(`matches.match_loser_team_id = %s`, team_field))
	}
}

func (filters *inningsFilters) setTossResult(value string, isBattingTeam bool) {
	if value != "won" && value != "lost" {
		return
	}

	matchField, inningsField := "matches.toss_loser_team_id", "innings.bowling_team_id"

	isTossWon := value == "won"
	if isTossWon {
		matchField = "matches.toss_winner_team_id"
	}

	if isBattingTeam {
		inningsField = "innings.batting_team_id"
	}

	condition := fmt.Sprintf(`%s = %s`, matchField, inningsField)
	filters.conditions = append(filters.conditions, condition)
}

func (filters *inningsFilters) setBatFieldFirst(value string, isBattingTeam bool) {
	if value != "bat" && value != "field" {
		return
	}

	isBatFirst := value == "bat"

	inningsField := "innings.bowling_team_id"
	if isBattingTeam {
		inningsField = "innings.batting_team_id"
	}

	condition := fmt.Sprintf(`(
			(matches.toss_winner_team_id = %s AND matches.is_toss_decision_bat = %v)
			OR
			(matches.toss_loser_team_id = %s AND matches.is_toss_decision_bat = %v)
	)`, inningsField, isBatFirst, inningsField, !isBatFirst)

	filters.conditions = append(filters.conditions, condition)
}

func (filters *inningsFilters) setInningsNumber(inningsNumbers []string) {
	var placeholders []string

	for _, inningsNumber := range inningsNumbers {
		filters.args = append(filters.args, inningsNumber)
		filters.placeholdersCount++
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, filters.placeholdersCount))
	}

	if len(placeholders) > 0 {
		condition := fmt.Sprintf(`innings.innings_number IN (%s)`, strings.Join(placeholders, ", "))
		filters.conditions = append(filters.conditions, condition)
	}
}

/* Batting Scorecard Filters */

func (filters *inningsFilters) setBatterRunsRange(minRuns, maxRuns string) {
	if minRuns != "" {
		filters.args = append(filters.args, minRuns)
		filters.placeholdersCount++
		filters.conditions = append(filters.conditions, fmt.Sprintf(`batting_scorecards.runs_scored >= $%d`, filters.placeholdersCount))
	}

	if maxRuns != "" {
		filters.args = append(filters.args, maxRuns)
		filters.placeholdersCount++
		filters.conditions = append(filters.conditions, fmt.Sprintf(`batting_scorecards.runs_scored <= $%d`, filters.placeholdersCount))
	}
}

func (filters *inningsFilters) setBatterPositionRange(minPosition, maxPosition string) {
	if minPosition != "" {
		filters.args = append(filters.args, minPosition)
		filters.placeholdersCount++
		filters.conditions = append(filters.conditions, fmt.Sprintf(`batting_scorecards.batting_position >= $%d`, filters.placeholdersCount))
	}

	if maxPosition != "" {
		filters.args = append(filters.args, maxPosition)
		filters.placeholdersCount++
		filters.conditions = append(filters.conditions, fmt.Sprintf(`batting_scorecards.batting_position <= $%d`, filters.placeholdersCount))
	}
}

func (filters *inningsFilters) setBatterIsDismissed(isDismissed string) {
	if isDismissed == "dismissed" {
		filters.conditions = append(filters.conditions, `batting_scorecards.dismissal_type IS NOT NULL AND batting_scorecards.dismissal_type NOT IN ('retired hurt', 'retired not out')`)
		return
	}

	if isDismissed == "not_out" {
		filters.conditions = append(filters.conditions, `batting_scorecards.dismissal_type IS NULL OR batting_scorecards.dismissal_type NOT IN ('retired hurt', 'retired not out')`)
		return
	}
}

func (filters *inningsFilters) setBatterDismissalTypes(dismissalTypes []string) {
	var placeholders []string

	for _, dismissalType := range dismissalTypes {
		filters.args = append(filters.args, dismissalType)
		filters.placeholdersCount++
		placeholders = append(placeholders, fmt.Sprintf("$%d", filters.placeholdersCount))
	}

	if len(placeholders) > 0 {
		filters.conditions = append(filters.conditions, fmt.Sprintf(`batting_scorecards.dismissal_type IN (%s)`, strings.Join(placeholders, ",")))
	}
}

/* Bowling Scorecard Filters */

func (filters *inningsFilters) setBowlerBallsRange(minBalls, maxBalls string) {
	if minBalls != "" {
		filters.args = append(filters.args, minBalls)
		filters.placeholdersCount++
		filters.conditions = append(filters.conditions, fmt.Sprintf(`bowling_scorecards.balls_bowled >= $%d`, filters.placeholdersCount))
	}

	if maxBalls != "" {
		filters.args = append(filters.args, maxBalls)
		filters.placeholdersCount++
		filters.conditions = append(filters.conditions, fmt.Sprintf(`bowling_scorecards.balls_bowled <= $%d`, filters.placeholdersCount))
	}
}

func (filters *inningsFilters) setBowlerRunsRange(minRuns, maxRuns string) {
	if minRuns != "" {
		filters.args = append(filters.args, minRuns)
		filters.placeholdersCount++
		filters.conditions = append(filters.conditions, fmt.Sprintf(`bowling_scorecards.runs_conceded >= $%d`, filters.placeholdersCount))
	}

	if maxRuns != "" {
		filters.args = append(filters.args, maxRuns)
		filters.placeholdersCount++
		filters.conditions = append(filters.conditions, fmt.Sprintf(`bowling_scorecards.runs_conceded <= $%d`, filters.placeholdersCount))
	}
}

func (filters *inningsFilters) setBowlerWicketsRange(minWickets, maxWickets string) {
	if minWickets != "" {
		filters.args = append(filters.args, minWickets)
		filters.placeholdersCount++
		filters.conditions = append(filters.conditions, fmt.Sprintf(`bowling_scorecards.wickets_taken >= $%d`, filters.placeholdersCount))
	}

	if maxWickets != "" {
		filters.args = append(filters.args, maxWickets)
		filters.placeholdersCount++
		filters.conditions = append(filters.conditions, fmt.Sprintf(`bowling_scorecards.wickets_taken <= $%d`, filters.placeholdersCount))
	}
}

func (filters *inningsFilters) setBowlerPositionRange(minPosition, maxPosition string) {
	if minPosition != "" {
		filters.args = append(filters.args, minPosition)
		filters.placeholdersCount++
		filters.conditions = append(filters.conditions, fmt.Sprintf(`bowling_scorecards.bowling_position >= $%d`, filters.placeholdersCount))
	}

	if maxPosition != "" {
		filters.args = append(filters.args, maxPosition)
		filters.placeholdersCount++
		filters.conditions = append(filters.conditions, fmt.Sprintf(`bowling_scorecards.bowling_position <= $%d`, filters.placeholdersCount))
	}
}

/* Team Innings Score Filters */

func (filters *inningsFilters) setTeamInningsRunsRange(minRuns, maxRuns string) {
	if minRuns != "" {
		filters.args = append(filters.args, minRuns)
		filters.placeholdersCount++
		filters.conditions = append(filters.conditions, fmt.Sprintf(`innings.total_runs >= $%d`, filters.placeholdersCount))
	}

	if maxRuns != "" {
		filters.args = append(filters.args, maxRuns)
		filters.placeholdersCount++
		filters.conditions = append(filters.conditions, fmt.Sprintf(`innings.total_runs <= $%d`, filters.placeholdersCount))
	}
}

func (filters *inningsFilters) setTeamInningsWktsRange(minWickets, maxWickets string) {
	if minWickets != "" {
		filters.args = append(filters.args, minWickets)
		filters.placeholdersCount++
		filters.conditions = append(filters.conditions, fmt.Sprintf(`innings.total_wickets >= $%d`, filters.placeholdersCount))
	}

	if maxWickets != "" {
		filters.args = append(filters.args, maxWickets)
		filters.placeholdersCount++
		filters.conditions = append(filters.conditions, fmt.Sprintf(`innings.total_wickets <= $%d`, filters.placeholdersCount))
	}
}

func (filters *inningsFilters) setTeamInningsBallsRange(minBalls, maxBalls string) {
	if minBalls != "" {
		filters.args = append(filters.args, minBalls)
		filters.placeholdersCount++
		filters.conditions = append(filters.conditions, fmt.Sprintf(`innings.total_balls >= $%d`, filters.placeholdersCount))
	}

	if maxBalls != "" {
		filters.args = append(filters.args, maxBalls)
		filters.placeholdersCount++
		filters.conditions = append(filters.conditions, fmt.Sprintf(`innings.total_balls <= $%d`, filters.placeholdersCount))
	}
}
