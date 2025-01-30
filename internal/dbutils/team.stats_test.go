package dbutils

import (
	"context"
	"log"
	"math"
	"net/url"
	"os"
	"slices"
	"testing"
	"time"

	"github.com/mainlycricket/CricKendra/internal/responses"
)

func TestRead_Overall_Team_Teams_Stats(t *testing.T) {
	ctx, DB_URL := context.Background(), os.Getenv("TEST_DB_URL")
	DB_POOL, err := Connect(ctx, DB_URL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	type args struct {
		ctx      context.Context
		db       DB_Exec
		queryMap url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    responses.StatsResponse[responses.Overall_Team_Teams_Group]
		wantErr bool
	}{
		{
			name: "Overall_Team_Teams_Group batting_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":  []string{"ODI"},
					"is_male":         []string{"true"},
					"min_start_date":  []string{"2008-01-01"},
					"max_start_date":  []string{"2023-12-31"},
					"min__total_runs": []string{"2000"},
					"team_total_for":  []string{"batting"},
					"sort_by":         []string{"total_runs"},
					"sort_order":      []string{"reverse"},
				},
			},
			wantErr: false,
		},

		{
			name: "Overall_Team_Teams_Group bowling_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":  []string{"ODI"},
					"is_male":         []string{"true"},
					"min_start_date":  []string{"2008-01-01"},
					"max_start_date":  []string{"2023-12-31"},
					"min__total_runs": []string{"2000"},
					"team_total_for":  []string{"bowling"},
					"sort_by":         []string{"total_runs"},
					"sort_order":      []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Team_Teams_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Team_Teams_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64
			for _, item := range got.Stats {
				if item.TotalRuns.Int64 < 2000 {
					t.Error("total_runs qualification failed")
					return
				}
				if item.TotalRuns.Int64 < prev {
					t.Error("total_runs sorting failed")
					return
				}
				prev = item.TotalRuns.Int64
			}
		})
	}
}

func TestRead_Overall_Team_Players_Stats(t *testing.T) {
	ctx, DB_URL := context.Background(), os.Getenv("TEST_DB_URL")
	DB_POOL, err := Connect(ctx, DB_URL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	type args struct {
		ctx      context.Context
		db       DB_Exec
		queryMap url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    responses.StatsResponse[responses.Overall_Team_Players_Group]
		wantErr bool
	}{
		{
			name: "Overall_Team_Players_Group batting_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":   []string{"ODI"},
					"is_male":          []string{"true"},
					"season":           []string{"2017/18", "2018", "2018/19"},
					"min__total_balls": []string{"500"},
					"team_total_for":   []string{"batting"},
					"sort_by":          []string{"total_balls"},
					"sort_order":       []string{"reverse"},
				},
			},
			wantErr: false,
		},

		{
			name: "Overall_Team_Players_Group bowling_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":   []string{"ODI"},
					"is_male":          []string{"true"},
					"season":           []string{"2017/18", "2018", "2018/19"},
					"min__total_balls": []string{"100"},
					"team_total_for":   []string{"bowling"},
					"sort_by":          []string{"total_balls"},
					"sort_order":       []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Team_Players_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Team_Players_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64
			for _, item := range got.Stats {
				if item.TotalBalls.Int64 < 100 {
					t.Error("total_balls qualification failed")
					return
				}
				if item.TotalBalls.Int64 < prev {
					t.Errorf(`total_balls sorting failed`)
					return
				}
				prev = item.TotalBalls.Int64
			}
		})
	}
}

func TestRead_Overall_Team_Matches_Stats(t *testing.T) {
	ctx, DB_URL := context.Background(), os.Getenv("TEST_DB_URL")
	DB_POOL, err := Connect(ctx, DB_URL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	type args struct {
		ctx      context.Context
		db       DB_Exec
		queryMap url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    responses.StatsResponse[responses.Overall_Team_Matches_Group]
		wantErr bool
	}{
		{
			name: "Overall_Batting_Match_Group batting_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":     []string{"ODI"},
					"is_male":            []string{"true"},
					"home_or_away":       []string{"home", "away"},
					"max__total_wickets": []string{"15"},
					"team_total_for":     []string{"batting"},
					"sort_by":            []string{"total_wickets"},
					"sort_order":         []string{"reverse"},
				},
			},
			wantErr: false,
		},

		{
			name: "Overall_Batting_Match_Group bowling_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":     []string{"ODI"},
					"is_male":            []string{"true"},
					"home_or_away":       []string{"home", "away"},
					"max__total_wickets": []string{"15"},
					"team_total_for":     []string{"bowling"},
					"sort_by":            []string{"total_wickets"},
					"sort_order":         []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Team_Matches_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Team_Matches_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64
			for _, item := range got.Stats {
				if item.TotalWickets.Int64 > 15 {
					t.Error("total_wickets qualifcation failed")
					return
				}
				if item.TotalWickets.Int64 < prev {
					t.Error(`total_wickets sorting failed`)
					return
				}
				prev = item.TotalWickets.Int64
			}
		})
	}
}

func TestRead_Overall_Team_Series_Stats(t *testing.T) {
	ctx, DB_URL := context.Background(), os.Getenv("TEST_DB_URL")
	DB_POOL, err := Connect(ctx, DB_URL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	type args struct {
		ctx      context.Context
		db       DB_Exec
		queryMap url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    responses.StatsResponse[responses.Overall_Team_Series_Group]
		wantErr bool
	}{
		{
			name: "Overall_Team_Series_Group batting_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format": []string{"ODI"},
					"is_male":        []string{"true"},
					"continent":      []string{"1", "2", "3"},
					"min__average":   []string{"30"},
					"team_total_for": []string{"batting"},
					"sort_by":        []string{"average"},
					"sort_order":     []string{"reverse"},
				},
			},
			wantErr: false,
		},

		{
			name: "Overall_Team_Series_Group bowling_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format": []string{"ODI"},
					"is_male":        []string{"true"},
					"continent":      []string{"1", "2", "3"},
					"min__average":   []string{"30"},
					"team_total_for": []string{"bowling"},
					"sort_by":        []string{"average"},
					"sort_order":     []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Team_Series_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Team_Series_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev float64
			for _, item := range got.Stats {
				if item.Average.Valid && item.Average.Float64 < 30 {
					t.Error("average qualification failed")
					return
				}

				if item.Average.Valid && item.Average.Float64 < prev {
					t.Error("average sorting failed")
					return
				}

				if item.Average.Valid {
					prev = item.Average.Float64
				}
			}
		})
	}
}

func TestRead_Overall_Team_Tournaments_Stats(t *testing.T) {
	ctx, DB_URL := context.Background(), os.Getenv("TEST_DB_URL")
	DB_POOL, err := Connect(ctx, DB_URL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	type args struct {
		ctx      context.Context
		db       DB_Exec
		queryMap url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    responses.StatsResponse[responses.Overall_Team_Tournament_Group]
		wantErr bool
	}{
		{
			name: "Overall_Team_Tournament_Group batting_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":    []string{"ODI"},
					"is_male":           []string{"true"},
					"host_nation":       []string{"1", "2", "3", "7", "9", "11"},
					"min__scoring_rate": []string{"4"},
					"team_total_for":    []string{"batting"},
					"sort_by":           []string{"scoring_rate"},
					"sort_order":        []string{"reverse"},
				},
			},
			wantErr: false,
		},

		{
			name: "Overall_Team_Tournament_Group bowling_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":    []string{"ODI"},
					"is_male":           []string{"true"},
					"host_nation":       []string{"1", "2", "3", "7", "9", "11"},
					"min__scoring_rate": []string{"4"},
					"team_total_for":    []string{"bowling"},
					"sort_by":           []string{"scoring_rate"},
					"sort_order":        []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Team_Tournaments_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Team_Tournaments_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev float64
			for _, item := range got.Stats {
				if item.ScoringRate.Valid && item.ScoringRate.Float64 < 4 {
					t.Error("scoring_rate qualification failed")
					return
				}

				if item.ScoringRate.Valid && item.ScoringRate.Float64 < prev {
					t.Error("scoring_rate sorting failed")
					return
				}

				if item.ScoringRate.Valid {
					prev = item.ScoringRate.Float64
				}
			}
		})
	}
}

func TestRead_Overall_Team_Grounds_Stats(t *testing.T) {
	ctx, DB_URL := context.Background(), os.Getenv("TEST_DB_URL")
	DB_POOL, err := Connect(ctx, DB_URL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	type args struct {
		ctx      context.Context
		db       DB_Exec
		queryMap url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    responses.StatsResponse[responses.Overall_Team_Grounds_Group]
		wantErr bool
	}{
		{
			name: "Overall_Team_Grounds_Group batting_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":   []string{"ODI"},
					"is_male":          []string{"true"},
					"ground":           []string{"1", "2", "3", "7", "9", "11"},
					"min__matches_won": []string{"2"},
					"team_total_for":   []string{"batting"},
					"sort_by":          []string{"matches_won"},
					"sort_order":       []string{"reverse"},
				},
			},
			wantErr: false,
		},
		{
			name: "Overall_Team_Grounds_Group bowling_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":   []string{"ODI"},
					"is_male":          []string{"true"},
					"ground":           []string{"1", "2", "3", "7", "9", "11"},
					"min__matches_won": []string{"2"},
					"team_total_for":   []string{"bowling"},
					"sort_by":          []string{"matches_won"},
					"sort_order":       []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Team_Grounds_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Team_Grounds_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			validGrounds := []int64{1, 2, 3, 5, 7, 9, 11}

			var prev int64
			for _, item := range got.Stats {
				if !slices.Contains(validGrounds, item.GroundId.Int64) {
					t.Error("ground filter found")
					return
				}

				if item.MatchesWon.Int64 < 2 {
					t.Error("matches_won qualification failed")
					return
				}

				if item.MatchesWon.Int64 < prev {
					t.Error("matches_won sorting failed")
					return
				}

				prev = item.MatchesWon.Int64
			}
		})
	}
}

func TestRead_Overall_Team_HostNations_Stats(t *testing.T) {
	ctx, DB_URL := context.Background(), os.Getenv("TEST_DB_URL")
	DB_POOL, err := Connect(ctx, DB_URL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	type args struct {
		ctx      context.Context
		db       DB_Exec
		queryMap url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    responses.StatsResponse[responses.Overall_Team_HostNations_Group]
		wantErr bool
	}{
		{
			name: "Overall_Team_HostNations_Group batting_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":    []string{"ODI"},
					"is_male":           []string{"true"},
					"series":            []string{"1", "2", "3", "7", "9", "11"},
					"min__matches_lost": []string{"1"},
					"team_total_for":    []string{"batting"},
					"sort_by":           []string{"matches_lost"},
					"sort_order":        []string{"reverse"},
				},
			},
			wantErr: false,
		},

		{
			name: "Overall_Team_HostNations_Group bowling_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":    []string{"ODI"},
					"is_male":           []string{"true"},
					"series":            []string{"1", "2", "3", "7", "9", "11"},
					"min__matches_lost": []string{"1"},
					"team_total_for":    []string{"bowling"},
					"sort_by":           []string{"matches_lost"},
					"sort_order":        []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Team_HostNations_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Team_HostNations_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64
			for _, item := range got.Stats {
				if item.MatchesLost.Int64 < 1 {
					t.Error("matches_lost qualification failed")
					return
				}

				if item.MatchesLost.Int64 < prev {
					t.Error("matches_lost sorting failed")
					return
				}

				prev = item.MatchesLost.Int64
			}
		})
	}
}

func TestRead_Overall_Team_Continents_Stats(t *testing.T) {
	ctx, DB_URL := context.Background(), os.Getenv("TEST_DB_URL")
	DB_POOL, err := Connect(ctx, DB_URL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	type args struct {
		ctx      context.Context
		db       DB_Exec
		queryMap url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    responses.StatsResponse[responses.Overall_Team_Continents_Group]
		wantErr bool
	}{
		{
			name: "Overall_Team_Continents_Group batting_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":              []string{"ODI"},
					"is_male":                     []string{"true"},
					"tournament":                  []string{"1", "2"},
					"max__matches_with_no_result": []string{"5"},
					"team_total_for":              []string{"batting"},
					"sort_by":                     []string{"matches_with_no_result"},
					"sort_order":                  []string{"reverse"},
				},
			},
			wantErr: false,
		},
		{
			name: "Overall_Team_Continents_Group bowling_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":              []string{"ODI"},
					"is_male":                     []string{"true"},
					"tournament":                  []string{"1", "2"},
					"max__matches_with_no_result": []string{"10"},
					"team_total_for":              []string{"bowling"},
					"sort_by":                     []string{"matches_with_no_result"},
					"sort_order":                  []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Team_Continents_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Team_Continents_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64
			for _, item := range got.Stats {
				if item.MatchesNoResult.Int64 > 10 {
					t.Error("matches_with_no_result qualification failed")
					return
				}

				if item.MatchesNoResult.Int64 < prev {
					t.Error("matches_with_no_result sorting failed")
					return
				}

				prev = item.MatchesNoResult.Int64
			}
		})
	}
}

func TestRead_Overall_Team_Years_Stats(t *testing.T) {
	ctx, DB_URL := context.Background(), os.Getenv("TEST_DB_URL")
	DB_POOL, err := Connect(ctx, DB_URL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	type args struct {
		ctx      context.Context
		db       DB_Exec
		queryMap url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    responses.StatsResponse[responses.Overall_Team_Years_Group]
		wantErr bool
	}{
		{
			name: "Overall_Team_Years_Group batting_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":    []string{"ODI"},
					"is_male":           []string{"true"},
					"match_result":      []string{"won", "tied", "drawn"},
					"min__matches_tied": []string{"2"},
					"team_total_for":    []string{"batting"},
					"sort_by":           []string{"matches_tied"},
					"sort_order":        []string{"reverse"},
				},
			},
			wantErr: false,
		},

		{
			name: "Overall_Team_Years_Group bowling_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":    []string{"ODI"},
					"is_male":           []string{"true"},
					"match_result":      []string{"won", "tied", "drawn"},
					"min__matches_tied": []string{"2"},
					"team_total_for":    []string{"bowling"},
					"sort_by":           []string{"matches_tied"},
					"sort_order":        []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Team_Years_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Team_Years_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64
			for _, item := range got.Stats {
				if item.MatchesTied.Int64 < 2 {
					t.Error("matches_tied qualification failed")
					return
				}

				if item.MatchesTied.Int64 < prev {
					t.Error("matches_tied sorting failed")
					return
				}

				prev = item.MatchesTied.Int64
			}
		})
	}
}

func TestRead_Overall_Team_Seasons_Stats(t *testing.T) {
	ctx, DB_URL := context.Background(), os.Getenv("TEST_DB_URL")
	DB_POOL, err := Connect(ctx, DB_URL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	type args struct {
		ctx      context.Context
		db       DB_Exec
		queryMap url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    responses.StatsResponse[responses.Overall_Team_Seasons_Group]
		wantErr bool
	}{
		{
			name: "Overall_Team_Seasons_Group batting_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":     []string{"ODI"},
					"is_male":            []string{"true"},
					"toss_result":        []string{"won"},
					"min__innings_count": []string{"25"},
					"team_total_for":     []string{"batting"},
					"sort_by":            []string{"innings_count"},
					"sort_order":         []string{"reverse"},
				},
			},
			wantErr: false,
		},
		{
			name: "Overall_Team_Seasons_Group bowling_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":     []string{"ODI"},
					"is_male":            []string{"true"},
					"toss_result":        []string{"won"},
					"min__innings_count": []string{"25"},
					"team_total_for":     []string{"bowling"},
					"sort_by":            []string{"innings_count"},
					"sort_order":         []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Team_Seasons_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Team_Seasons_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64
			for _, item := range got.Stats {
				if item.InningsCount.Int64 < 25 {
					t.Error("innings_count qualification failed")
					return
				}

				if item.InningsCount.Int64 < prev {
					t.Error("innings_count sorting failed")
					return
				}

				prev = item.InningsCount.Int64
			}
		})
	}
}

func TestRead_Overall_Team_Decades_Stats(t *testing.T) {
	ctx, DB_URL := context.Background(), os.Getenv("TEST_DB_URL")
	DB_POOL, err := Connect(ctx, DB_URL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	type args struct {
		ctx      context.Context
		db       DB_Exec
		queryMap url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    responses.StatsResponse[responses.Overall_Team_Decades_Group]
		wantErr bool
	}{
		{
			name: "Overall_Team_Decades_Group batting_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":      []string{"ODI"},
					"is_male":             []string{"true"},
					"bat_field_first":     []string{"first"},
					"min__matches_played": []string{"50"},
					"team_total_for":      []string{"batting"},
					"sort_by":             []string{"matches_played"},
					"sort_order":          []string{"reverse"},
				},
			},
			wantErr: false,
		},

		{
			name: "Overall_Team_Decades_Group bowling_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":      []string{"ODI"},
					"is_male":             []string{"true"},
					"bat_field_first":     []string{"first"},
					"min__matches_played": []string{"50"},
					"team_total_for":      []string{"bowling"},
					"sort_by":             []string{"matches_played"},
					"sort_order":          []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Team_Decades_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Team_Decades_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64
			for _, item := range got.Stats {
				if item.MatchesPlayed.Int64 < 50 {
					t.Error("matches_played qualification failed")
					return
				}

				if item.MatchesPlayed.Int64 < prev {
					t.Error("matches_played sorting failed")
					return
				}

				prev = item.MatchesPlayed.Int64
			}
		})
	}
}

func TestRead_Overall_Team_Aggregate_Stats(t *testing.T) {
	ctx, DB_URL := context.Background(), os.Getenv("TEST_DB_URL")
	DB_POOL, err := Connect(ctx, DB_URL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	type args struct {
		ctx      context.Context
		db       DB_Exec
		queryMap url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    responses.StatsResponse[responses.Overall_Team_Aggregate_Group]
		wantErr bool
	}{
		{
			name: "Overall_Team_Aggregate_Group batting_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":  []string{"ODI"},
					"is_male":         []string{"true"},
					"toss_result":     []string{"lost"},
					"bat_field_first": []string{"field"},
					"team_total_for":  []string{"batting"},
				},
			},
			wantErr: false,
		},

		{
			name: "Overall_Team_Aggregate_Group bowling_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":  []string{"ODI"},
					"is_male":         []string{"true"},
					"toss_result":     []string{"lost"},
					"bat_field_first": []string{"field"},
					"team_total_for":  []string{"bowling"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Team_Aggregate_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Team_Aggregate_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(got.Stats) > 1 {
				t.Error("got more than more 1 item")
				return
			}
		})
	}
}

func TestRead_Individual_Team_Innings_Stats(t *testing.T) {
	ctx, DB_URL := context.Background(), os.Getenv("TEST_DB_URL")
	DB_POOL, err := Connect(ctx, DB_URL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	type args struct {
		ctx      context.Context
		db       DB_Exec
		queryMap url.Values
	}

	tests := []struct {
		name    string
		args    args
		want    responses.StatsResponse[responses.Individual_Team_Innings_Group]
		wantErr bool
	}{
		{
			name: "Individual_Team_Innings_Group batting_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":            []string{"ODI"},
					"is_male":                   []string{"true"},
					"min__team_innings_runs":    []string{"250"},
					"max__team_innings_balls":   []string{"270"},
					"max__team_innings_wickets": []string{"5"},
					"team_total_for":            []string{"batting"},
					"sort_by":                   []string{"total_runs"},
					"sort_order":                []string{"reverse"},
				},
			},
			wantErr: false,
		},
		{
			name: "Individual_Team_Innings_Group bowling_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":            []string{"ODI"},
					"is_male":                   []string{"true"},
					"min__team_innings_runs":    []string{"250"},
					"max__team_innings_balls":   []string{"270"},
					"max__team_innings_wickets": []string{"5"},
					"team_total_for":            []string{"bowling"},
					"sort_by":                   []string{"total_runs"},
					"sort_order":                []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Individual_Team_Innings_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Team_Innings_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64
			for _, item := range got.Stats {
				if item.TotalRuns.Int64 < 250 {
					t.Error("min__innings_total_runs qualification failed")
					return
				}

				if item.TotalOvers.Float64 > 45 {
					t.Error("min__innings_total_balls qualification failed")
					return
				}

				if item.TotalWickets.Int64 > 5 {
					t.Error("min__innings_total_wickets qualification failed")
					return
				}

				if item.TotalRuns.Int64 < prev {
					t.Error("total_runs sorting failed")
					return
				}

				prev = item.TotalRuns.Int64
			}
		})
	}
}

func TestRead_Individual_Team_MatchTotals_Stats(t *testing.T) {
	ctx, DB_URL := context.Background(), os.Getenv("TEST_DB_URL")
	DB_POOL, err := Connect(ctx, DB_URL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	type args struct {
		ctx      context.Context
		db       DB_Exec
		queryMap url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    responses.StatsResponse[responses.Individual_Team_MatchTotals_Group]
		wantErr bool
	}{
		{
			name: "Individual_Team_MatchTotals_Group batting_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":            []string{"ODI"},
					"is_male":                   []string{"true"},
					"min__team_innings_runs":    []string{"250"},
					"max__team_innings_runs":    []string{"300"},
					"max__team_innings_wickets": []string{"5"},
					"team_total_for":            []string{"batting"},
					"sort_by":                   []string{"start_date"},
					"sort_order":                []string{"reverse"},
				},
			},
			wantErr: false,
		},
		{
			name: "Individual_Team_MatchTotals_Group bowling_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":            []string{"ODI"},
					"is_male":                   []string{"true"},
					"min__team_innings_runs":    []string{"250"},
					"max__team_innings_runs":    []string{"300"},
					"max__team_innings_wickets": []string{"5"},
					"team_total_for":            []string{"bowling"},
					"sort_by":                   []string{"start_date"},
					"sort_order":                []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Individual_Team_MatchTotals_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Team_MatchTotals_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev time.Time
			for _, item := range got.Stats {
				if item.TotalRuns.Int64 < 250 {
					t.Error("min__innings_total_runs qualification failed")
					return
				}

				if item.TotalRuns.Int64 > 300 {
					t.Error("max__innings_total_runs qualification failed")
					return
				}

				if item.TotalWickets.Int64 > 5 {
					t.Error("min__innings_total_wickets qualification failed")
					return
				}

				if item.StartDate.Time.Nanosecond() < prev.Nanosecond() {
					t.Error("start_date sorting failed")
					return
				}

				prev = item.StartDate.Time
			}
		})
	}
}

func TestRead_Individual_Team_MatchResults_Stats(t *testing.T) {
	ctx, DB_URL := context.Background(), os.Getenv("TEST_DB_URL")
	DB_POOL, err := Connect(ctx, DB_URL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	type args struct {
		ctx      context.Context
		db       DB_Exec
		queryMap url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    responses.StatsResponse[responses.Individual_Team_MatchResults_Group]
		wantErr bool
	}{
		{
			name: "Individual_Team_MatchResults_Group batting_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":  []string{"ODI"},
					"is_male":         []string{"true"},
					"primary_team":    []string{"1", "2", "3"},
					"opposition_team": []string{"1", "2", "3"},
					"team_total_for":  []string{"batting"},
					"sort_by":         []string{"start_date"},
					"sort_order":      []string{"reverse"},
				},
			},
			wantErr: false,
		},

		{
			name: "Individual_Team_MatchResults_Group bowling_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":  []string{"ODI"},
					"is_male":         []string{"true"},
					"primary_team":    []string{"1", "2", "3"},
					"opposition_team": []string{"1", "2", "3"},
					"team_total_for":  []string{"bowling"},
					"sort_by":         []string{"start_date"},
					"sort_order":      []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Individual_Team_MatchResults_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Team_MatchResults_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev time.Time
			validTeams := []int64{1, 2, 3}

			for _, item := range got.Stats {
				if !slices.Contains(validTeams, item.TeamId.Int64) {
					t.Error("invalid team id")
					return
				}

				if !slices.Contains(validTeams, item.OppositionId.Int64) {
					t.Error("invalid opposition id")
					return
				}

				if item.StartDate.Time.Nanosecond() < prev.Nanosecond() {
					t.Error("start_date sorting failed")
					return
				}

				prev = item.StartDate.Time
			}
		})
	}
}

func TestRead_Individual_Team_Series_Stats(t *testing.T) {
	ctx, DB_URL := context.Background(), os.Getenv("TEST_DB_URL")
	DB_POOL, err := Connect(ctx, DB_URL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	type args struct {
		ctx      context.Context
		db       DB_Exec
		queryMap url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    responses.StatsResponse[responses.Individual_Team_Series_Group]
		wantErr bool
	}{
		{
			name: "Individual_Team_Series_Group batting_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":    []string{"ODI"},
					"is_male":           []string{"true"},
					"tournament":        []string{"1"},
					"min__scoring_rate": []string{"5"},
					"team_total_for":    []string{"batting"},
					"sort_by":           []string{"team_name"},
					"sort_order":        []string{"default"},
				},
			},
			wantErr: false,
		},

		{
			name: "Individual_Team_Series_Group bowling_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":    []string{"ODI"},
					"is_male":           []string{"true"},
					"tournament":        []string{"1"},
					"min__scoring_rate": []string{"5"},
					"team_total_for":    []string{"bowling"},
					"sort_by":           []string{"team_name"},
					"sort_order":        []string{"default"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Individual_Team_Series_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Team_Series_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for _, item := range got.Stats {
				if item.ScoringRate.Valid && item.ScoringRate.Float64 < 5 {
					t.Error("min__scoring_rate qualification failed")
					return
				}
			}
		})
	}
}

func TestRead_Individual_Team_Tournaments_Stats(t *testing.T) {
	ctx, DB_URL := context.Background(), os.Getenv("TEST_DB_URL")
	DB_POOL, err := Connect(ctx, DB_URL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	type args struct {
		ctx      context.Context
		db       DB_Exec
		queryMap url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    responses.StatsResponse[responses.Individual_Team_Tournaments_Group]
		wantErr bool
	}{
		{
			name: "Individual_Team_Tournaments_Group batting_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":    []string{"ODI"},
					"is_male":           []string{"true"},
					"home":              []string{"home"},
					"min__average":      []string{"25"},
					"min__scoring_rate": []string{"4.5"},
					"team_total_for":    []string{"batting"},
					"sort_by":           []string{"average"},
					"sort_order":        []string{"default"},
				},
			},
			wantErr: false,
		},
		{
			name: "Individual_Team_Tournaments_Group bowling_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":    []string{"ODI"},
					"is_male":           []string{"true"},
					"home":              []string{"home"},
					"min__average":      []string{"25"},
					"min__scoring_rate": []string{"4.5"},
					"team_total_for":    []string{"bowling"},
					"sort_by":           []string{"average"},
					"sort_order":        []string{"default"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Individual_Team_Tournaments_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Team_Tournaments_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev float64 = math.MaxFloat64
			for _, item := range got.Stats {
				if item.Average.Valid && item.Average.Float64 < 25 {
					t.Error("average qualification failed")
					return
				}

				if item.ScoringRate.Valid && item.ScoringRate.Float64 < 4.5 {
					t.Error("scoring_rate qualification failed")
					return
				}

				if item.Average.Valid && item.Average.Float64 > prev {
					t.Error("average sorting failed")
					return
				}

				if item.Average.Valid {
					prev = item.Average.Float64
				}
			}
		})
	}
}

func TestRead_Individual_Team_Grounds_Stats(t *testing.T) {
	ctx, DB_URL := context.Background(), os.Getenv("TEST_DB_URL")
	DB_POOL, err := Connect(ctx, DB_URL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	type args struct {
		ctx      context.Context
		db       DB_Exec
		queryMap url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    responses.StatsResponse[responses.Individual_Team_Grounds_Group]
		wantErr bool
	}{
		{
			name: "Individual_Team_Grounds_Group batting_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":      []string{"ODI"},
					"is_male":             []string{"true"},
					"host_nation":         []string{"1", "2"},
					"min__win_loss_ratio": []string{"0.7"},
					"max__win_loss_ratio": []string{"2"},
					"team_total_for":      []string{"batting"},
					"sort_by":             []string{"win_loss_ratio"},
					"sort_order":          []string{"reverse"},
				},
			},
			wantErr: false,
		},
		{
			name: "Individual_Team_Grounds_Group bowling_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":      []string{"ODI"},
					"is_male":             []string{"true"},
					"host_nation":         []string{"1", "2"},
					"min__win_loss_ratio": []string{"0.7"},
					"max__win_loss_ratio": []string{"2"},
					"team_total_for":      []string{"bowling"},
					"sort_by":             []string{"win_loss_ratio"},
					"sort_order":          []string{"reverse"},
					"__limit":             []string{"200"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Individual_Team_Grounds_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Team_Grounds_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev float64
			for _, item := range got.Stats {
				if item.WinLossRatio.Valid && item.WinLossRatio.Float64 < 0.7 {
					t.Error("min__win_loss_ratio qualification failed")
					return
				}

				if item.WinLossRatio.Valid && item.WinLossRatio.Float64 > 2 {
					t.Error("min__win_loss_ratio qualification failed")
					return
				}

				if item.WinLossRatio.Valid && item.WinLossRatio.Float64 < prev {
					t.Error("win_loss_ratio sorting failed")
					return
				}

				if item.WinLossRatio.Valid {
					prev = item.WinLossRatio.Float64
				}
			}
		})
	}
}

func TestRead_Individual_Team_HostNations_Stats(t *testing.T) {
	ctx, DB_URL := context.Background(), os.Getenv("TEST_DB_URL")
	DB_POOL, err := Connect(ctx, DB_URL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	type args struct {
		ctx      context.Context
		db       DB_Exec
		queryMap url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    responses.StatsResponse[responses.Individual_Team_HostNations_Group]
		wantErr bool
	}{
		{
			name: "Individual_Team_HostNations_Group batting_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":      []string{"ODI"},
					"is_male":             []string{"true"},
					"continent":           []string{"3"},
					"team_total_for":      []string{"batting"},
					"min__matches_played": []string{"5"},
					"sort_by":             []string{"lowest_score"},
					"sort_order":          []string{"default"},
					"__limit":             []string{"200"},
				},
			},
			wantErr: false,
		},
		{
			name: "Individual_Team_HostNations_Group bowling_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":      []string{"ODI"},
					"is_male":             []string{"true"},
					"continent":           []string{"3"},
					"team_total_for":      []string{"bowling"},
					"min__matches_played": []string{"5"},
					"sort_by":             []string{"lowest_score"},
					"sort_order":          []string{"default"},
					"__limit":             []string{"200"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Individual_Team_HostNations_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Team_HostNations_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64 = math.MaxInt64
			for _, item := range got.Stats {
				if item.MatchesPlayed.Int64 < 5 {
					t.Error("matches_played qualification failed")
					return
				}

				if item.LowestScore.Valid && item.LowestScore.Int64 > prev {
					t.Error("lowest_score sorting failed")
				}

				if item.LowestScore.Valid {
					prev = item.LowestScore.Int64
				}
			}
		})
	}
}

func TestRead_Individual_Team_Years_Stats(t *testing.T) {
	ctx, DB_URL := context.Background(), os.Getenv("TEST_DB_URL")
	DB_POOL, err := Connect(ctx, DB_URL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	type args struct {
		ctx      context.Context
		db       DB_Exec
		queryMap url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    responses.StatsResponse[responses.Individual_Team_Years_Group]
		wantErr bool
	}{
		{
			name: "Individual_Team_Years_Group batting_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format": []string{"T20I"},
					"is_male":        []string{"true"},
					"min_start_date": []string{"2019-01-01"},
					"max_start_date": []string{"2023-12-31"},
					"sort_by":        []string{"start_date"},
					"sort_order":     []string{"reverse"},
					"team_total_for": []string{"batting"},
					"__limit":        []string{"200"},
				},
			},
			wantErr: false,
		},
		{
			name: "Individual_Team_Years_Group bowling_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format": []string{"T20I"},
					"is_male":        []string{"true"},
					"min_start_date": []string{"2019-01-01"},
					"max_start_date": []string{"2023-12-31"},
					"sort_by":        []string{"start_date"},
					"sort_order":     []string{"reverse"},
					"team_total_for": []string{"bowling"},
					"__limit":        []string{"200"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Individual_Team_Years_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Team_Years_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64 = math.MaxInt64
			for _, item := range got.Stats {
				if item.Year.Int64 < 2019 || item.Year.Int64 > 2023 {
					t.Error("year exceeded range")
					return
				}

				if item.Year.Int64 > prev {
					t.Error("year sorting failed")
					return
				}

				prev = item.Year.Int64
			}
		})
	}
}

func TestRead_Individual_Team_Seasons_Stats(t *testing.T) {
	ctx, DB_URL := context.Background(), os.Getenv("TEST_DB_URL")
	DB_POOL, err := Connect(ctx, DB_URL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	type args struct {
		ctx      context.Context
		db       DB_Exec
		queryMap url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    responses.StatsResponse[responses.Individual_Team_Seasons_Group]
		wantErr bool
	}{
		{
			name: "Individual_Team_Seasons_Group batting_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format": []string{"T20I"},
					"is_male":        []string{"true"},
					"min_start_date": []string{"2019-01-01"},
					"max_start_date": []string{"2023-12-31"},
					"team_total_for": []string{"batting"},
					"sort_by":        []string{"highest_score"},
					"__limit":        []string{"200"},
				},
			},
			wantErr: false,
		},
		{
			name: "Individual_Team_Seasons_Group bowling_team",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format": []string{"T20I"},
					"is_male":        []string{"true"},
					"min_start_date": []string{"2019-01-01"},
					"max_start_date": []string{"2023-12-31"},
					"team_total_for": []string{"bowling"},
					"sort_by":        []string{"highest_score"},
					"__limit":        []string{"200"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Individual_Team_Seasons_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Team_Seasons_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64 = math.MaxInt64
			for _, item := range got.Stats {
				if item.HighestScore.Int64 > prev {
					t.Error("highest_score sorting failed")
					return
				}

				prev = item.HighestScore.Int64
			}
		})
	}
}
