package statqueries

import (
	"fmt"
	"net/url"
)

type qualificationField struct {
	name   string
	clause string
}

type qualificationFilters struct {
	fields            []qualificationField
	conditions        []string
	args              []any
	placeholdersCount int
}

func newQualificationFilters(stats_type, group_type int) *qualificationFilters {
	qf := &qualificationFilters{}
	qf.fields = getQualificationFields(stats_type, group_type)
	return qf
}

func (qf *qualificationFilters) getClause() string {
	return prefixJoin(qf.conditions, "HAVING", " AND ")
}

func (qf *qualificationFilters) applyQualifications(params *url.Values) {
	for _, field := range qf.fields {
		if min_value := params.Get("min__" + field.name); len(min_value) > 0 {
			qf.args = append(qf.args, min_value)
			qf.placeholdersCount++
			condition := fmt.Sprintf(`%s >= $%d`, field.clause, qf.placeholdersCount)
			qf.conditions = append(qf.conditions, condition)
		}

		if max_value := params.Get("max__" + field.name); len(max_value) > 0 {
			qf.args = append(qf.args, max_value)
			qf.placeholdersCount++
			condition := fmt.Sprintf(`%s <= $%d`, field.clause, qf.placeholdersCount)
			qf.conditions = append(qf.conditions, condition)
		}
	}
}

func getQualificationFields(stats_type, group_type int) []qualificationField {
	switch stats_type {
	case batting_stats:
		return getBattingQualificationFields(group_type)
	case bowling_stats:
		return getBowlingQualificationFields(group_type)
	case team_stats:
		return getTeamQualificationFields(group_type)
	default:
		return nil
	}
}

func getBattingQualificationFields(_ int) []qualificationField {
	return []qualificationField{
		{
			name:   "matches_played",
			clause: matches_count_query,
		},
		{
			name:   "innings_batted",
			clause: batting_innings_count,
		},
		{
			name:   "not_outs",
			clause: batting_not_outs,
		},
		{
			name:   "runs_scored",
			clause: batting_runs_scored,
		},
		{
			name:   "balls_faced",
			clause: batting_balls_faced,
		},
		{
			name:   "average",
			clause: `(CASE WHEN COUNT(CASE WHEN batting_scorecards.dismissal_type IS NOT NULL AND batting_scorecards.dismissal_type NOT IN ('retired hurt', 'retired not out') THEN 1 END) > 0 THEN SUM(batting_scorecards.runs_scored) * 1.0 / COUNT(CASE WHEN batting_scorecards.dismissal_type IS NOT NULL AND batting_scorecards.dismissal_type NOT IN ('retired hurt', 'retired not out') THEN 1 END) ELSE '+infinity' END)`,
		},
		{
			name:   "strike_rate",
			clause: batting_strike_rate,
		},
		{
			name:   "centuries",
			clause: batting_centuries,
		},
		{
			name:   "half_centuries",
			clause: batting_half_centuries,
		},
		{
			name:   "fifty_plus_scores",
			clause: batting_fifty_plus,
		},
		{
			name:   "ducks",
			clause: batting_ducks,
		},
		{
			name:   "fours_scored",
			clause: batting_fours_scored,
		},
		{
			name:   "sixes_scored",
			clause: batting_sixes_scored,
		},
	}
}

func getBowlingQualificationFields(group_type int) []qualificationField {
	fields := []qualificationField{
		{
			name:   "matches_played",
			clause: matches_count_query,
		},
		{
			name:   "innings_bowled",
			clause: inningsBowled_query,
		},
		{
			name:   "overs_bowled",
			clause: oversBowled_query,
		},
		{
			name:   "maiden_overs",
			clause: maidenOvers_query,
		},
		{
			name:   "runs_conceded",
			clause: runsConceded_query,
		},
		{
			name:   "wickets_taken",
			clause: wicketsTaken_query,
		},
		{
			name:   "average",
			clause: bowlingAverage_query,
		},
		{
			name:   "strike_rate",
			clause: bowlingStrikeRate_query,
		},
		{
			name:   "economy",
			clause: bowlingEconomy_query,
		},
		{
			name:   "fours_conceded",
			clause: foursConceded_query,
		},
		{
			name:   "sixes_conceded",
			clause: sixesConceded_query,
		},
	}

	if group_type < inidividual_innings_group || group_type > inidividual_matchTotals_group {
		fields = append(fields,
			qualificationField{
				name:   "four_wkt_hauls",
				clause: fourWktHauls_query,
			},
			qualificationField{
				name:   "five_wkt_hauls",
				clause: fiveWktHauls_query,
			},
			qualificationField{
				name:   "ten_wkt_hauls",
				clause: tenWktHauls_query,
			})
	}

	return fields
}

func getTeamQualificationFields(_ int) []qualificationField {
	return []qualificationField{
		{
			name:   "matches_played",
			clause: matches_count_query,
		},
		{
			name:   "matches_tied",
			clause: matchesNoResult_query,
		},
		{
			name:   "matches_drawn",
			clause: matchesDrawn_query,
		},
		{
			name:   "matches_with_no_result",
			clause: matchesNoResult_query,
		},
		{
			name:   "innings_count",
			clause: teamInningsCount_query,
		},
		{
			name:   "total_runs",
			clause: teamTotalRuns_query,
		},
		{
			name:   "total_balls",
			clause: teamTotalBalls_query,
		},
		{
			name:   "total_wickets",
			clause: teamTotalWkts_query,
		},
		{
			name:   "average",
			clause: `(CASE WHEN SUM(innings.total_wickets) > 0 THEN SUM (innings.total_runs) * 1.0 / SUM(innings.total_wickets) ELSE '+infinity' END)`,
		},
		{
			name:   "scoring_rate",
			clause: `(CASE WHEN SUM(innings.total_balls) > 0 THEN SUM (innings.total_runs) * 6.0 / SUM(innings.total_balls) ELSE '+infinity' END)`,
		},
	}
}

func getTeamMatchQualifications(teamField string) []qualificationField {
	return []qualificationField{
		{
			name:   "matches_won",
			clause: getMatchesWonQuery(teamField),
		},
		{
			name:   "matches_lost",
			clause: getMatchesLostQuery(teamField),
		},
		{
			name:   "win_loss_ratio",
			clause: getWinLossRatioInfiniteQuery(teamField),
		},
	}
}
