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
