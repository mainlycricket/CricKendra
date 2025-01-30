package statqueries

import "testing"

func Test_getSortingClause(t *testing.T) {
	type args struct {
		sortBy      string
		sortOrder   string
		sortingKeys []sortingKey
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "common batting sorting keys, sort by balls faced reverse",
			args: args{
				sortBy:    "balls_faced",
				sortOrder: "reverse",
				sortingKeys: []sortingKey{
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
				},
			},
			want: `ORDER BY SUM(batting_scorecards.balls_faced) ASC,SUM(batting_scorecards.runs_scored) DESC,(CASE WHEN COUNT(CASE WHEN batting_scorecards.dismissal_type IS NOT NULL AND batting_scorecards.dismissal_type NOT IN ('retired hurt', 'retired not out') THEN 1 END) > 0 THEN SUM(batting_scorecards.runs_scored) * 1.0 / COUNT(CASE WHEN batting_scorecards.dismissal_type IS NOT NULL AND batting_scorecards.dismissal_type NOT IN ('retired hurt', 'retired not out') THEN 1 END) ELSE '+infinity' END) DESC,(CASE WHEN SUM(batting_scorecards.balls_faced) > 0 THEN SUM(batting_scorecards.runs_scored) * 100.0 / SUM(batting_scorecards.balls_faced) ELSE 0 END) DESC,COUNT(CASE WHEN batting_scorecards.runs_scored >= 100 THEN 1 END) DESC,COUNT(CASE WHEN batting_scorecards.runs_scored >= 50 THEN 1 END) DESC,COUNT(CASE WHEN batting_scorecards.runs_scored BETWEEN 50 AND 99 THEN 1 END) DESC,SUM(batting_scorecards.fours_scored) DESC,SUM(batting_scorecards.fours_scored) DESC,COUNT(CASE WHEN batting_scorecards.has_batted THEN innings.id END) DESC,COUNT(CASE WHEN batting_scorecards.has_batted AND (batting_scorecards.dismissal_type IS NULL OR batting_scorecards.dismissal_type IN ('retired hurt', 'retired not out')) THEN 1 END) DESC,COUNT(DISTINCT matches.id) DESC,COUNT(CASE WHEN batting_scorecards.has_batted AND batting_scorecards.runs_scored = 0 AND batting_scorecards.dismissal_type IS NOT NULL AND batting_scorecards.dismissal_type NOT IN ('retired hurt', 'retired not out') THEN 1 END) DESC`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSortingClause(tt.args.sortBy, tt.args.sortOrder, tt.args.sortingKeys); got != tt.want {
				t.Errorf("getSortingClause() = %v, want %v", got, tt.want)
			}
		})
	}
}
