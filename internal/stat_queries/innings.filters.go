package statqueries

import "fmt"

type inningsFilters struct {
	conditions []string
}

func newInningsFilters() *inningsFilters {
	inningsFilters := &inningsFilters{}
	return inningsFilters
}

func (filters *inningsFilters) setBattingTeams(placeholderString string) {
	if len(placeholderString) > 0 {
		condition := fmt.Sprintf(`innings.batting_team_id IN (%s)`, placeholderString)
		filters.conditions = append(filters.conditions, condition)
	}
}

func (filters *inningsFilters) setBowlingTeams(placeholderString string) {
	if len(placeholderString) > 0 {
		condition := fmt.Sprintf(`innings.bowling_team_id IN (%s)`, placeholderString)
		filters.conditions = append(filters.conditions, condition)
	}
}

func (filters *inningsFilters) setHomeAway(isHomeTeam, isBattingTeam bool) {
	matchField, inningsField := "matches.away_team_id", "innings.bowling_team_id"

	if isHomeTeam {
		matchField = "matches.home_team_id"
	}

	if isBattingTeam {
		inningsField = "innings.batting_team_id"
	}

	condition := fmt.Sprintf(`%s = %s`, matchField, inningsField)
	filters.conditions = append(filters.conditions, condition)
}

func (filters *inningsFilters) setTossResult(isTossWon, isBattingTeam bool) {
	matchField, inningsField := "matches.toss_loser_team_id", "innings.bowling_team_id"

	if isTossWon {
		matchField = "matches.toss_winner_team_id"
	}

	if isBattingTeam {
		inningsField = "innings.batting_team_id"
	}

	condition := fmt.Sprintf(`%s = %s`, matchField, inningsField)
	filters.conditions = append(filters.conditions, condition)
}

func (filters *inningsFilters) setBatFieldFirst(isBatFirst, isBattingTeam bool) {
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
