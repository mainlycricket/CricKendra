package statqueries

import (
	"slices"
)

type sortingKey struct {
	name        string
	defaultSort string
	reverseSort string
}

func getSortingClause(sortBy, sortOrder string, sortingKeys []sortingKey) string {
	idx := slices.IndexFunc(sortingKeys, func(item sortingKey) bool {
		return item.name == sortBy
	})

	newSortingKeys := make([]sortingKey, 0, len(sortingKeys))

	if idx != -1 {
		newSortingKeys = append(newSortingKeys, sortingKeys[idx])
		newSortingKeys = append(newSortingKeys, sortingKeys[:idx]...)
		newSortingKeys = append(newSortingKeys, sortingKeys[idx+1:]...)
	} else {
		newSortingKeys = sortingKeys[:]
	}

	var parameters []string

	for _, sortingKey := range newSortingKeys {
		if sortingKey.name == sortBy && sortOrder == "reverse" {
			parameters = append(parameters, sortingKey.reverseSort)
		} else {
			parameters = append(parameters, sortingKey.defaultSort)
		}
	}

	return prefixJoin(parameters, "ORDER BY", ",")
}

func getSortingKeys(stats_type, group_type int) []sortingKey {
	switch stats_type {
	case batting_stats:
		return getBattingSortingKeys(group_type)
	case bowling_stats:
		return getBowlingSortingKeys(group_type)
	case team_stats:
		return getTeamSortingKeys(group_type)
	default:
		return nil
	}

}

func getBattingSortingKeys(group_type int) []sortingKey {
	sortingKeys := []sortingKey{
		{
			name:        "runs_scored",
			defaultSort: batting_runs_scored + " DESC",
			reverseSort: batting_runs_scored + " ASC",
		},
		{
			name:        "average",
			defaultSort: `(CASE WHEN COUNT(CASE WHEN batting_scorecards.dismissal_type IS NOT NULL AND batting_scorecards.dismissal_type NOT IN ('retired hurt', 'retired not out') THEN 1 END) > 0 THEN SUM(batting_scorecards.runs_scored) * 1.0 / COUNT(CASE WHEN batting_scorecards.dismissal_type IS NOT NULL AND batting_scorecards.dismissal_type NOT IN ('retired hurt', 'retired not out') THEN 1 END) ELSE '+infinity' END) DESC`,
			reverseSort: `(CASE WHEN COUNT(CASE WHEN batting_scorecards.dismissal_type IS NOT NULL AND batting_scorecards.dismissal_type NOT IN ('retired hurt', 'retired not out') THEN 1 END) > 0 THEN SUM(batting_scorecards.runs_scored) * 1.0 / COUNT(CASE WHEN batting_scorecards.dismissal_type IS NOT NULL AND batting_scorecards.dismissal_type NOT IN ('retired hurt', 'retired not out') THEN 1 END) ELSE '+infinity' END) ASC`,
		},
		{
			name:        "strike_rate",
			defaultSort: batting_strike_rate + " DESC",
			reverseSort: batting_strike_rate + " ASC",
		},
		{
			name:        "centuries",
			defaultSort: batting_centuries + " DESC",
			reverseSort: batting_centuries + " ASC",
		},
		{
			name:        "fifty_plus_scores",
			defaultSort: batting_fifty_plus + " DESC",
			reverseSort: batting_fifty_plus + " ASC",
		},
		{
			name:        "half_centuries",
			defaultSort: batting_half_centuries + " DESC",
			reverseSort: batting_half_centuries + " ASC",
		},
		{
			name:        "fours_scored",
			defaultSort: batting_fours_scored + " DESC",
			reverseSort: batting_fours_scored + " ASC",
		},
		{
			name:        "sixes_scored",
			defaultSort: batting_fours_scored + " DESC",
			reverseSort: batting_fours_scored + " ASC",
		},
		{
			name:        "balls_faced",
			defaultSort: batting_balls_faced + " DESC",
			reverseSort: batting_balls_faced + " ASC",
		},
		{
			name:        "innings_batted",
			defaultSort: batting_innings_count + " DESC",
			reverseSort: batting_innings_count + " ASC",
		},
		{
			name:        "not_outs",
			defaultSort: batting_not_outs + " DESC",
			reverseSort: batting_not_outs + " ASC",
		},
		{
			name:        "matches_played",
			defaultSort: matches_count_query + " DESC",
			reverseSort: matches_count_query + " ASC",
		},
		{
			name:        "ducks",
			defaultSort: batting_ducks + " DESC",
			reverseSort: batting_ducks + " ASC",
		},
	}

	if group_type == inidividual_innings_group {
		sortingKeys = append(sortingKeys,
			sortingKey{
				name:        "innings_number",
				defaultSort: "innings.innings_number ASC",
				reverseSort: "innings.innings_number DESC",
			},
			sortingKey{
				name:        "batting_position",
				defaultSort: "batting_scorecards.batting_position ASC",
				reverseSort: "batting_scorecards.batting_position DESC",
			},
			sortingKey{
				name:        "dismissal_type",
				defaultSort: "batting_scorecards.dismissal_type ASC",
				reverseSort: "batting_scorecards.dismissal_type DESC",
			},
		)
	}

	if group_type == overall_players_group || group_type >= inidividual_innings_group {
		sortingKeys = append(sortingKeys, sortingKey{
			name:        "player_name",
			defaultSort: "batter_name ASC",
			reverseSort: "batter_name DESC",
		})
	}

	if group_type > overall_players_group && group_type < inidividual_innings_group {
		sortingKeys = append(sortingKeys, sortingKey{
			name:        "players_count",
			defaultSort: "COUNT(DISTINCT mse.player_id) DESC",
			reverseSort: "COUNT(DISTINCT mse.player_id) ASC",
		})
	}

	sortingKeys = append(sortingKeys, sortingKey{
		name:        "start_date",
		defaultSort: "MIN(matches.start_date) ASC, MAX(matches.start_date) ASC",
		reverseSort: "MAX(matches.start_date) DESC, MIN(matches.start_date) DESC",
	})

	return sortingKeys
}

func getBowlingSortingKeys(group_type int) []sortingKey {
	sortingKeys := []sortingKey{
		{
			name:        "wickets_taken",
			defaultSort: wicketsTaken_query + " DESC",
			reverseSort: wicketsTaken_query + " ASC",
		},
		{
			name:        "average",
			defaultSort: `(CASE WHEN SUM(bowling_scorecards.wickets_taken) > 0 THEN SUM(bowling_scorecards.runs_conceded) * 1.0 / SUM(bowling_scorecards.wickets_taken) ELSE '+infinity' END) ASC`,
			reverseSort: `(CASE WHEN SUM(bowling_scorecards.wickets_taken) > 0 THEN SUM(bowling_scorecards.runs_conceded) * 1.0 / SUM(bowling_scorecards.wickets_taken) ELSE '+infinity' END) DESC`,
		},
		{
			name:        "strike_rate",
			defaultSort: `(CASE WHEN SUM(bowling_scorecards.wickets_taken) > 0 THEN SUM(bowling_scorecards.balls_bowled) * 1.0 / SUM(bowling_scorecards.wickets_taken) ELSE '+infinity' END) ASC`,
			reverseSort: `(CASE WHEN SUM(bowling_scorecards.wickets_taken) > 0 THEN SUM(bowling_scorecards.balls_bowled) * 1.0 / SUM(bowling_scorecards.wickets_taken) ELSE '+infinity' END) DESC`,
		},
		{
			name:        "economy",
			defaultSort: `(CASE WHEN SUM(bowling_scorecards.balls_bowled) > 0 THEN SUM(bowling_scorecards.runs_conceded) * 6.0 / SUM(bowling_scorecards.balls_bowled) ELSE '+infinity' END) ASC`,
			reverseSort: `(CASE WHEN SUM(bowling_scorecards.balls_bowled) > 0 THEN SUM(bowling_scorecards.runs_conceded) * 6.0 / SUM(bowling_scorecards.balls_bowled) ELSE '+infinity' END) DESC`,
		},
		{
			name:        "overs_bowled",
			defaultSort: oversBowled_query + " DESC",
			reverseSort: oversBowled_query + " ASC",
		},
		{
			name:        "runs_conceded",
			defaultSort: runsConceded_query + " DESC",
			reverseSort: runsConceded_query + " ASC",
		},
		{
			name:        "maiden_overs",
			defaultSort: maidenOvers_query + " DESC",
			reverseSort: maidenOvers_query + " ASC",
		},
		{
			name:        "innings_bowled",
			defaultSort: inningsBowled_query + " DESC",
			reverseSort: inningsBowled_query + " ASC",
		},
		{
			name:        "matches_played",
			defaultSort: matches_count_query + " DESC",
			reverseSort: matches_count_query + " ASC",
		},
	}

	if group_type < inidividual_innings_group || group_type > inidividual_matchTotals_group {
		sortingKeys = append(sortingKeys,
			sortingKey{
				name:        "four_wkt_hauls",
				defaultSort: fourWktHauls_query + " DESC",
				reverseSort: fourWktHauls_query + " ASC",
			},
			sortingKey{
				name:        "five_wkt_hauls",
				defaultSort: fiveWktHauls_query + " DESC",
				reverseSort: fiveWktHauls_query + " ASC",
			},
			sortingKey{
				name:        "ten_wkt_hauls",
				defaultSort: tenWktHauls_query + " DESC",
				reverseSort: tenWktHauls_query + " ASC",
			},
			sortingKey{
				name:        "best_bowling_match",
				defaultSort: bestMatchWkts_query + " DESC," + bestMatchRuns_query + " ASC",
				reverseSort: bestMatchWkts_query + " ASC," + bestMatchRuns_query + " DESC",
			},
			sortingKey{
				name:        "best_bowling_innings",
				defaultSort: bestInningsWkts_query + " DESC," + bestInningsRuns_query + " ASC",
				reverseSort: bestInningsWkts_query + " ASC," + bestInningsRuns_query + " DESC",
			})
	}

	if group_type == inidividual_innings_group {
		sortingKeys = append(sortingKeys,
			sortingKey{
				name:        "innings_number",
				defaultSort: "innings.innings_number ASC",
				reverseSort: "innings.innings_number DESC",
			},
			sortingKey{
				name:        "bowling_position",
				defaultSort: "bowling_scorecards.bowling_position ASC",
				reverseSort: "bowling_scorecards.batting_position DESC",
			},
		)
	}

	sortingKeys = append(sortingKeys,
		sortingKey{
			name:        "fours_conceded",
			defaultSort: foursConceded_query + " DESC",
			reverseSort: foursConceded_query + " ASC",
		},
		sortingKey{
			name:        "sixes_conceded",
			defaultSort: sixesConceded_query + " DESC",
			reverseSort: sixesConceded_query + " ASC",
		},
	)

	if group_type == overall_players_group || group_type >= inidividual_innings_group {
		sortingKeys = append(sortingKeys, sortingKey{
			name:        "player_name",
			defaultSort: "bowler_name ASC",
			reverseSort: "bowler_name DESC",
		})
	}

	if group_type > overall_players_group && group_type < inidividual_innings_group {
		sortingKeys = append(sortingKeys, sortingKey{
			name:        "players_count",
			defaultSort: "COUNT(DISTINCT mse.player_id) DESC",
			reverseSort: "COUNT(DISTINCT mse.player_id) ASC",
		})
	}

	sortingKeys = append(sortingKeys, sortingKey{
		name:        "start_date",
		defaultSort: "MIN(matches.start_date) ASC, MAX(matches.start_date) ASC",
		reverseSort: "MAX(matches.start_date) DESC, MIN(matches.start_date) DESC",
	})

	return sortingKeys
}

func getTeamSortingKeys(group_type int) []sortingKey {
	sortingKeys := []sortingKey{
		{
			name:        "matches_played",
			defaultSort: matches_count_query + " DESC",
			reverseSort: matches_count_query + " ASC",
		},
		{
			name:        "matches_tied",
			defaultSort: matchesTied_query + " DESC",
			reverseSort: matchesTied_query + " ASC",
		},
		{
			name:        "matches_drawn",
			defaultSort: matchesDrawn_query + " DESC",
			reverseSort: matchesDrawn_query + " ASC",
		},
		{
			name:        "matches_with_no_result",
			defaultSort: matchesNoResult_query + " DESC",
			reverseSort: matchesNoResult_query + " ASC",
		},
		{
			name:        "total_runs",
			defaultSort: teamTotalRuns_query + " DESC",
			reverseSort: teamTotalRuns_query + " ASC",
		},
		{
			name:        "total_wickets",
			defaultSort: teamTotalWkts_query + " DESC",
			reverseSort: teamTotalWkts_query + " ASC",
		},
		{
			name:        "total_balls",
			defaultSort: teamTotalBalls_query + " DESC",
			reverseSort: teamTotalBalls_query + " ASC",
		},
		{
			name:        "average",
			defaultSort: `(CASE WHEN SUM(innings.total_wickets) > 0 THEN SUM (innings.total_runs) * 1.0 / SUM(innings.total_wickets) ELSE '+infinity' END) DESC`,
			reverseSort: `(CASE WHEN SUM(innings.total_wickets) > 0 THEN SUM (innings.total_runs) * 1.0 / SUM(innings.total_wickets) ELSE '+infinity' END) ASC`,
		},
		{
			name:        "scoring_rate",
			defaultSort: `(CASE WHEN SUM(innings.total_balls) > 0 THEN SUM (innings.total_runs) * 6.0 / SUM(innings.total_balls) ELSE '+infinity' END) DESC`,
			reverseSort: `(CASE WHEN SUM(innings.total_balls) > 0 THEN SUM (innings.total_runs) * 6.0 / SUM(innings.total_balls) ELSE '+infinity' END) ASC`,
		},
		{
			name:        "innings_count",
			defaultSort: teamInningsCount_query + " DESC",
			reverseSort: teamInningsCount_query + " ASC",
		},
		{
			name:        "highest_score",
			defaultSort: teamHighestScore_query + " DESC",
			reverseSort: teamHighestScore_query + " ASC",
		},
		{
			name:        "lowest_score",
			defaultSort: team_LowestScore_query + " DESC",
			reverseSort: team_LowestScore_query + " ASC",
		},
		{
			name:        "start_date",
			defaultSort: "MIN(matches.start_date) ASC, MAX(matches.start_date) ASC",
			reverseSort: "MAX(matches.start_date) DESC, MIN(matches.start_date) DESC",
		},
	}

	if group_type == inidividual_matchTotals_group || group_type == inidividual_matchResults_group {
		sortingKeys = append(sortingKeys, sortingKey{
			name:        "match_id",
			defaultSort: "matches.id DESC",
			reverseSort: "matches.id ASC",
		})
	}

	if group_type == inidividual_innings_group {
		sortingKeys = append(sortingKeys, sortingKey{
			name:        "innings_number",
			defaultSort: "innings.innings_number DESC",
			reverseSort: "innings.innings_number ASC",
		})
	}

	if group_type == overall_players_group {
		sortingKeys = append(sortingKeys, sortingKey{
			name:        "player_name",
			defaultSort: "player_name ASC",
			reverseSort: "player_name DESC",
		})
	}

	if group_type == overall_teams_group || group_type >= inidividual_innings_group {
		sortingKeys = append(sortingKeys, sortingKey{
			name:        "team_name",
			defaultSort: "team_name ASC",
			reverseSort: "team_name DESC",
		})
	}

	return sortingKeys
}

func getTeamMatchSortingKeys(teamField string) []sortingKey {
	matchesWonQuery := getMatchesWonQuery(teamField)
	matchesLostQuery := getMatchesLostQuery(teamField)
	winLossRatioQuery := getWinLossRatioInfiniteQuery(teamField)

	return []sortingKey{
		{
			name:        "matches_won",
			defaultSort: matchesWonQuery + " DESC",
			reverseSort: matchesWonQuery + " ASC",
		},
		{
			name:        "matches_lost",
			defaultSort: matchesLostQuery + " DESC",
			reverseSort: matchesLostQuery + " ASC",
		},
		{
			name:        "win_loss_ratio",
			defaultSort: winLossRatioQuery + " DESC",
			reverseSort: winLossRatioQuery + " ASC",
		},
	}
}
