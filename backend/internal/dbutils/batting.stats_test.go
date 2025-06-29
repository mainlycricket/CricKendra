package dbutils

import (
	"context"
	"log"
	"math"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/mainlycricket/CricKendra/backend/internal/responses"
)

func TestRead_Overall_Batting_Summary_Stats(t *testing.T) {
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
		wantErr bool
	}{
		{
			name: "simple batting summary read",
			args: args{
				ctx: ctx,
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format": []string{"ODI"},
					"is_male":        []string{"true"},
					"min_start_date": []string{"2008-01-01"},
					"max_start_date": []string{"2023-12-31"},
					"batter":         []string{"114"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Read_Overall_Batting_Summary_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Batting_Summary_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRead_Overall_Batting_Batters_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Batting_Batter_Group]
		wantErr bool
	}{
		{
			name: "Overall_Batting_Batter_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":   []string{"ODI"},
					"is_male":          []string{"true"},
					"min_start_date":   []string{"2008-01-01"},
					"max_start_date":   []string{"2023-12-31"},
					"min__runs_scored": []string{"1000"},
					"sort_by":          []string{"runs_scored"},
					"sort_order":       []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Batting_Batters_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Batting_Batters_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64
			for _, item := range got.Stats {
				if item.RunsScored.Int64 < 1000 {
					t.Error("runs_scored qualification failed")
					return
				}
				if item.RunsScored.Int64 < prev {
					t.Error("runs_scored sorting failed")
					return
				}
				prev = item.RunsScored.Int64
			}
		})
	}
}

func TestRead_Overall_Batting_TeamInnings_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Batting_TeamInnings_Group]
		wantErr bool
	}{
		{
			name: "Overall_Batting_TeamInnings_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":   []string{"ODI"},
					"is_male":          []string{"true"},
					"season":           []string{"2017/18", "2018", "2018/19"},
					"min__balls_faced": []string{"250"},
					"sort_by":          []string{"balls_faced"},
					"sort_order":       []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Batting_TeamInnings_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Batting_TeamInnings_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64
			for _, item := range got.Stats {
				if item.BallsFaced.Int64 < 250 {
					t.Error("balls_faced qualification failed")
					return
				}
				if item.BallsFaced.Int64 < prev {
					t.Errorf(`balls_faced sorting failed`)
					return
				}
				prev = item.BallsFaced.Int64
			}
		})
	}
}

func TestRead_Overall_Batting_Matches_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Batting_Match_Group]
		wantErr bool
	}{
		{
			name: "Overall_Batting_Match_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":         []string{"ODI"},
					"is_male":                []string{"true"},
					"home_or_away":           []string{"home", "away"},
					"max__fifty_plus_scores": []string{"3"},
					"sort_by":                []string{"fifty_plus_scores"},
					"sort_order":             []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Batting_Matches_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Batting_Matches_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64
			for _, item := range got.Stats {
				if item.FiftyPlusScores.Int64 > 3 {
					t.Error("fifty_plus_scores qualifcation failed")
					return
				}
				if item.FiftyPlusScores.Int64 < prev {
					t.Error(`fifty_plus_scores sorting failed`)
					return
				}
				prev = item.FiftyPlusScores.Int64
			}
		})
	}
}

func TestRead_Overall_Batting_Teams_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Batting_Team_Group]
		wantErr bool
	}{
		{
			name: "Overall_Batting_Team_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format": []string{"ODI"},
					"is_male":        []string{"true"},
					"continent":      []string{"1", "2", "3"},
					"min__average":   []string{"30"},
					"sort_by":        []string{"average"},
					"sort_order":     []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Batting_Teams_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Batting_Teams_Stats() error = %v, wantErr %v", err, tt.wantErr)
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

func TestRead_Overall_Batting_Oppositions_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Batting_Opposition_Group]
		wantErr bool
	}{
		{
			name: "Overall_Batting_Opposition_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":   []string{"ODI"},
					"is_male":          []string{"true"},
					"host_nation":      []string{"1", "2", "3", "7", "9", "11"},
					"min__strike_rate": []string{"75"},
					"sort_by":          []string{"strike_rate"},
					"sort_order":       []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Batting_Oppositions_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Batting_Oppositions_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev float64
			for _, item := range got.Stats {
				if item.StrikeRate.Valid && item.StrikeRate.Float64 < 75 {
					t.Error("strike_rate qualification failed")
					return
				}

				if item.StrikeRate.Valid && item.StrikeRate.Float64 < prev {
					t.Error("strike_rate sorting failed")
					return
				}

				if item.StrikeRate.Valid {
					prev = item.StrikeRate.Float64
				}
			}
		})
	}
}

func TestRead_Overall_Batting_Grounds_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Batting_Ground_Group]
		wantErr bool
	}{
		{
			name: "Overall_Batting_Ground_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format": []string{"ODI"},
					"is_male":        []string{"true"},
					"ground":         []string{"1", "2", "3", "7", "9", "11"},
					"min__centuries": []string{"1"},
					"sort_by":        []string{"centuries"},
					"sort_order":     []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Batting_Grounds_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Batting_Grounds_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64
			for _, item := range got.Stats {
				if item.Centuries.Int64 < 1 {
					t.Error("centuries qualification failed")
					return
				}

				if item.Centuries.Int64 < prev {
					t.Error("centuries sorting failed")
					return
				}

				prev = item.Centuries.Int64
			}
		})
	}
}

func TestRead_Overall_Batting_HostNations_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Batting_HostNation_Group]
		wantErr bool
	}{
		{
			name: "Overall_Batting_HostNation_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":      []string{"ODI"},
					"is_male":             []string{"true"},
					"series":              []string{"1", "2", "3", "7", "9", "11"},
					"min__half_centuries": []string{"1"},
					"sort_by":             []string{"half_centuries"},
					"sort_order":          []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Batting_HostNations_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Batting_HostNations_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64
			for _, item := range got.Stats {
				if item.HalfCenturies.Int64 < 1 {
					t.Error("half_centuries qualification failed")
					return
				}

				if item.HalfCenturies.Int64 < prev {
					t.Error("half_centuries sorting failed")
					return
				}

				prev = item.HalfCenturies.Int64
			}
		})
	}
}

func TestRead_Overall_Batting_Continents_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Batting_Continent_Group]
		wantErr bool
	}{
		{
			name: "Overall_Batting_Continent_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":    []string{"ODI"},
					"is_male":           []string{"true"},
					"tournament":        []string{"1", "2"},
					"min__fours_scored": []string{"50"},
					"sort_by":           []string{"fours_scored"},
					"sort_order":        []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Batting_Continents_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Batting_Continents_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64
			for _, item := range got.Stats {
				if item.FoursScored.Int64 < 50 {
					t.Error("fours_scored qualification failed")
					return
				}

				if item.FoursScored.Int64 < prev {
					t.Error("fours_scored sorting failed")
					return
				}

				prev = item.FoursScored.Int64
			}
		})
	}
}

func TestRead_Overall_Batting_Series_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Batting_Series_Group]
		wantErr bool
	}{
		{
			name: "Overall_Batting_Series_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":    []string{"ODI"},
					"is_male":           []string{"true"},
					"match_result":      []string{"won", "tied", "drawn"},
					"min__sixes_scored": []string{"25"},
					"sort_by":           []string{"sixes_scored"},
					"sort_order":        []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Batting_Series_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Batting_Series_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64
			for _, item := range got.Stats {
				if item.SixesScored.Int64 < 25 {
					t.Error("sixes_scored qualification failed")
					return
				}

				if item.SixesScored.Int64 < prev {
					t.Error("sixes_scored sorting failed")
					return
				}

				prev = item.SixesScored.Int64
			}
		})
	}
}

func TestRead_Overall_Batting_Tournaments_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Batting_Tournament_Group]
		wantErr bool
	}{
		{
			name: "Overall_Batting_Tournament_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":      []string{"ODI"},
					"is_male":             []string{"true"},
					"toss_result":         []string{"won"},
					"min__innings_batted": []string{"50"},
					"sort_by":             []string{"innings_batted"},
					"sort_order":          []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Batting_Tournaments_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Batting_Tournaments_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64
			for _, item := range got.Stats {
				if item.InningsBatted.Int64 < 50 {
					t.Error("innings_batted qualification failed")
					return
				}

				if item.InningsBatted.Int64 < prev {
					t.Error("innings_batted sorting failed")
					return
				}

				prev = item.InningsBatted.Int64
			}
		})
	}
}

func TestRead_Overall_Batting_Years_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Batting_Year_Group]
		wantErr bool
	}{
		{
			name: "Overall_Batting_Year_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":      []string{"ODI"},
					"is_male":             []string{"true"},
					"bat_field_first":     []string{"first"},
					"min__matches_played": []string{"50"},
					"sort_by":             []string{"matches_played"},
					"sort_order":          []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Batting_Years_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Batting_Years_Stats() error = %v, wantErr %v", err, tt.wantErr)
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

func TestRead_Overall_Batting_Seasons_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Batting_Season_Group]
		wantErr bool
	}{
		{
			name: "Overall_Batting_Season_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format": []string{"ODI"},
					"is_male":        []string{"true"},
					"innings_number": []string{"2"},
					"min__ducks":     []string{"5"},
					"sort_by":        []string{"ducks"},
					"sort_order":     []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Batting_Seasons_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Batting_Seasons_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64
			for _, item := range got.Stats {
				if item.Ducks.Int64 < 5 {
					t.Error("ducks qualification failed")
					return
				}

				if item.Ducks.Int64 < prev {
					t.Error("ducks sorting failed")
					return
				}

				prev = item.Ducks.Int64
			}
		})
	}
}

func TestRead_Overall_Batting_Decades_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Batting_Decade_Group]
		wantErr bool
	}{
		{
			name: "Overall_Batting_Decade_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":           []string{"ODI"},
					"is_male":                  []string{"true"},
					"min__innings_runs_scored": []string{"30"},
					"min__not_outs":            []string{"50"},
					"sort_by":                  []string{"not_outs"},
					"sort_order":               []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Batting_Decades_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Batting_Decades_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64
			for _, item := range got.Stats {
				if item.NotOuts.Int64 < 50 {
					t.Error("not_outs qualification failed")
					return
				}

				if item.NotOuts.Int64 < prev {
					t.Error("not_outs sorting failed")
					return
				}

				prev = item.NotOuts.Int64
			}
		})
	}
}

func TestRead_Overall_Batting_Aggregate_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Batting_Aggregate_Group]
		wantErr bool
	}{
		{
			name: "Overall_Batting_Aggregate_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":                []string{"ODI"},
					"is_male":                       []string{"true"},
					"min__innings_batting_position": []string{"1"},
					"max__innings_batting_position": []string{"3"},
					"toss_result":                   []string{"lost"},
					"bat_field_first":               []string{"field"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Batting_Aggregate_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Batting_Aggregate_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got.Stats) > 1 {
				t.Error("got more than more 1 item")
				return
			}
		})
	}
}

func TestRead_Individual_Batting_Innings_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Individual_Batting_Innings_Group]
		wantErr bool
	}{
		{
			name: "Individual_Batting_Innings_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":                []string{"ODI"},
					"is_male":                       []string{"true"},
					"max__innings_runs_scored":      []string{"50"},
					"innings_is_batter_dismissed":   []string{"dismissed"},
					"min__innings_batting_position": []string{"1"},
					"max__innings_batting_position": []string{"7"},
					"innings_batter_dismissal_type": []string{"bowled", "caught"},
					"sort_by":                       []string{"runs_scored"},
					"sort_order":                    []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Individual_Batting_Innings_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Batting_Innings_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64
			for _, item := range got.Stats {
				if item.RunsScored.Int64 > 50 {
					t.Error("max__innings_runs_scored qualification failed")
					return
				}

				if item.RunsScored.Int64 < prev {
					t.Error("runs_scored sorting failed")
					return
				}

				prev = item.RunsScored.Int64
			}
		})
	}
}

func TestRead_Individual_Batting_MatchTotals_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Individual_Batting_MatchTotals_Group]
		wantErr bool
	}{
		{
			name: "Individual_Batting_MatchTotals_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":                []string{"ODI"},
					"is_male":                       []string{"true"},
					"min__innings_runs_scored":      []string{"50"},
					"max__innings_runs_scored":      []string{"99"},
					"innings_is_batter_dismissed":   []string{"dismissed"},
					"min__innings_batting_position": []string{"1"},
					"max__innings_batting_position": []string{"7"},
					"sort_by":                       []string{"start_date"},
					"sort_order":                    []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Individual_Batting_MatchTotals_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Batting_MatchTotals_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev time.Time
			for _, item := range got.Stats {
				if item.RunsScored.Int64 < 50 {
					t.Error("min__innings_runs_scored qualification failed")
					return
				}

				if item.RunsScored.Int64 > 99 {
					t.Error("max__innings_runs_scored qualification failed")
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

func TestRead_Individual_Batting_Series_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Individual_Batting_Series_Group]
		wantErr bool
	}{
		{
			name: "Individual_Batting_Series_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":                []string{"ODI"},
					"is_male":                       []string{"true"},
					"tournament":                    []string{"1"},
					"min__innings_batting_position": []string{"1"},
					"max__innings_batting_position": []string{"7"},
					"sort_by":                       []string{"player_name"},
					"sort_order":                    []string{"default"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Read_Individual_Batting_Series_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Batting_Series_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRead_Individual_Batting_Tournaments_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Individual_Batting_Tournaments_Group]
		wantErr bool
	}{
		{
			name: "Individual_Batting_Tournaments_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":   []string{"ODI"},
					"is_male":          []string{"true"},
					"home":             []string{"home"},
					"min__strike_rate": []string{"75"},
					"sort_by":          []string{"average"},
					"sort_order":       []string{"default"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Individual_Batting_Tournaments_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Batting_Tournaments_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev float64 = math.MaxFloat64
			for _, item := range got.Stats {
				if item.StrikeRate.Valid && item.StrikeRate.Float64 < 75 {
					t.Error("strike_rate qualification failed")
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

func TestRead_Individual_Batting_Grounds_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Individual_Batting_Ground_Group]
		wantErr bool
	}{
		{
			name: "Individual_Batting_Ground_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":   []string{"ODI"},
					"is_male":          []string{"true"},
					"host_nation":      []string{"1", "2"},
					"min__average":     []string{"25"},
					"min__strike_rate": []string{"75"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Individual_Batting_Grounds_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Batting_Grounds_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64 = math.MaxInt64
			for _, item := range got.Stats {
				if item.StrikeRate.Valid && item.StrikeRate.Float64 < 75 {
					t.Error("strike_rate qualification failed")
					return
				}

				if item.Average.Valid && item.Average.Float64 < 25 {
					t.Error("average qualification failed")
					return
				}

				if item.RunsScored.Int64 > prev {
					t.Error("runs_scored sorting failed")
				}

				prev = item.RunsScored.Int64
			}
		})
	}
}

func TestRead_Individual_Batting_HostNations_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Individual_Batting_HostNation_Group]
		wantErr bool
	}{
		{
			name: "Individual_Batting_Ground_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":      []string{"ODI"},
					"is_male":             []string{"true"},
					"continent":           []string{"3"},
					"min__matches_played": []string{"5"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Individual_Batting_HostNations_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Batting_HostNations_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64 = math.MaxInt64
			for _, item := range got.Stats {
				if item.MatchesPlayed.Int64 < 5 {
					t.Error("matches_played qualification failed")
					return
				}

				if item.RunsScored.Int64 > prev {
					t.Error("runs_scored sorting failed")
				}

				prev = item.RunsScored.Int64
			}
		})
	}
}

func TestRead_Individual_Batting_Oppositions_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Individual_Batting_Opposition_Group]
		wantErr bool
	}{
		{
			name: "Individual_Batting_Opposition_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":      []string{"ODI"},
					"is_male":             []string{"true"},
					"primary_team":        []string{"1", "2", "3"},
					"opposition_team":     []string{"1", "2", "3", "4", "5"},
					"max__innings_batted": []string{"15"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Individual_Batting_Oppositions_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Batting_Oppositions_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64 = math.MaxInt64
			for _, item := range got.Stats {
				if item.InningsBatted.Int64 > 15 {
					t.Error("innings_batted qualification failed")
					return
				}

				if item.RunsScored.Int64 > prev {
					t.Error("runs_scored sorting failed")
				}

				prev = item.RunsScored.Int64
			}
		})
	}
}

func TestRead_Individual_Batting_Years_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Individual_Batting_Year_Group]
		wantErr bool
	}{
		{
			name: "Individual_Batting_Year_Group",
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
					"__page":         []string{"2"},
					"__limit":        []string{"200"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Individual_Batting_Years_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Batting_Years_Stats() error = %v, wantErr %v", err, tt.wantErr)
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

func TestRead_Individual_Batting_Seasons_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Individual_Batting_Season_Group]
		wantErr bool
	}{
		{
			name: "Individual_Batting_Season_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format": []string{"T20I"},
					"is_male":        []string{"true"},
					"min_start_date": []string{"2019-01-01"},
					"max_start_date": []string{"2023-12-31"},
					"sort_by":        []string{"fifty_plus_scores"},
					"sort_order":     []string{"default"},
					"__page":         []string{"2"},
					"__limit":        []string{"200"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Individual_Batting_Seasons_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Batting_Seasons_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64 = math.MaxInt64
			for _, item := range got.Stats {
				if item.FiftyPlusScores.Int64 > prev {
					t.Error("fifty_plus_scores sorting failed")
					return
				}

				prev = item.FiftyPlusScores.Int64
			}
		})
	}
}
