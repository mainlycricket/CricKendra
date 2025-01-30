package dbutils

import (
	"context"
	"log"
	"math"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/mainlycricket/CricKendra/internal/responses"
)

func TestRead_Overall_Bowling_Bowlers_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Bowling_Bowler_Group]
		wantErr bool
	}{
		{
			name: "Overall_Bowling_Bowler_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":     []string{"ODI"},
					"is_male":            []string{"true"},
					"min_start_date":     []string{"2008-01-01"},
					"max_start_date":     []string{"2023-12-31"},
					"min__wickets_taken": []string{"250"},
					"sort_by":            []string{"wickets_taken"},
					"sort_order":         []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Bowling_Bowlers_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Bowling_Bowlers_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64
			for _, item := range got.Stats {
				if item.WicketsTaken.Int64 < 250 {
					t.Error("wickets_taken qualification failed")
					return
				}
				if item.WicketsTaken.Int64 < prev {
					t.Error("wickets_taken sorting failed")
					return
				}
				prev = item.WicketsTaken.Int64
			}
		})
	}
}

func TestRead_Overall_Bowling_TeamInnings_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Bowling_TeamInnings_Group]
		wantErr bool
	}{
		{
			name: "Overall_Bowling_TeamInnings_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":    []string{"ODI"},
					"is_male":           []string{"true"},
					"season":            []string{"2017/18", "2018", "2018/19"},
					"min__overs_bowled": []string{"50"},
					"sort_by":           []string{"overs_bowled"},
					"sort_order":        []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Bowling_TeamInnings_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Bowling_TeamInnings_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev float64
			for _, item := range got.Stats {
				if item.OversBowled.Float64 < 50 {
					t.Error("overs_bowled qualification failed")
					return
				}
				if item.OversBowled.Float64 < prev {
					t.Errorf(`overs_bowled sorting failed`)
					return
				}
				prev = item.OversBowled.Float64
			}
		})
	}
}

func TestRead_Overall_Bowling_Matches_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Bowling_Match_Group]
		wantErr bool
	}{
		{
			name: "Overall_Bowling_Match_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":      []string{"ODI"},
					"is_male":             []string{"true"},
					"home_or_away":        []string{"home", "away"},
					"max__four_wkt_hauls": []string{"1"},
					"sort_by":             []string{"four_wkt_hauls"},
					"sort_order":          []string{"default"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Bowling_Matches_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Bowling_Matches_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64 = math.MaxInt64
			for _, item := range got.Stats {
				if item.FourWktHauls.Int64 < 1 {
					t.Error("four_wkt_hauls qualifcation failed")
					return
				}
				if item.FourWktHauls.Int64 > prev {
					t.Error(`four_wkt_hauls sorting failed`)
					return
				}
				prev = item.FourWktHauls.Int64
			}
		})
	}
}

func TestRead_Overall_Bowling_Teams_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Bowling_Team_Group]
		wantErr bool
	}{
		{
			name: "Overall_Bowling_Team_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format": []string{"ODI"},
					"is_male":        []string{"true"},
					"continent":      []string{"1", "2", "3"},
					"max__average":   []string{"35"},
					"sort_by":        []string{"average"},
					"sort_order":     []string{"default"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Bowling_Teams_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Bowling_Teams_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev float64
			for _, item := range got.Stats {
				if item.Average.Valid && item.Average.Float64 > 35 {
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

func TestRead_Overall_Bowling_Oppositions_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Bowling_Opposition_Group]
		wantErr bool
	}{
		{
			name: "Overall_Bowling_Opposition_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":   []string{"ODI"},
					"is_male":          []string{"true"},
					"host_nation":      []string{"1", "2", "3", "7", "9", "11"},
					"min__strike_rate": []string{"25"},
					"sort_by":          []string{"strike_rate"},
					"sort_order":       []string{"default"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Bowling_Oppositions_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Bowling_Oppositions_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev float64
			for _, item := range got.Stats {
				if item.StrikeRate.Valid && item.StrikeRate.Float64 < 25 {
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

func TestRead_Overall_Bowling_Grounds_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Bowling_Ground_Group]
		wantErr bool
	}{
		{
			name: "Overall_Bowling_Ground_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":      []string{"ODI"},
					"is_male":             []string{"true"},
					"continent":           []string{"1", "2", "3"},
					"min__five_wkt_hauls": []string{"1"},
					"sort_by":             []string{"players_count"},
					"sort_order":          []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Bowling_Grounds_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Bowling_Grounds_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64
			for _, item := range got.Stats {
				if item.FiveWktHauls.Int64 < 1 {
					t.Error("five_wkt_hauls qualification failed")
					return
				}

				if item.PlayersCount.Int64 < prev {
					t.Error("players_count sorting failed")
					return
				}

				prev = item.PlayersCount.Int64
			}
		})
	}
}

func TestRead_Overall_Bowling_HostNations_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Bowling_HostNation_Group]
		wantErr bool
	}{
		{
			name: "Overall_Bowling_HostNation_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":     []string{"ODI"},
					"is_male":            []string{"true"},
					"series":             []string{"1", "2", "3", "7", "9", "11"},
					"max__ten_wkt_hauls": []string{"0"},
					"sort_by":            []string{"best_bowling_innings"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Bowling_HostNations_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Bowling_HostNations_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prevWickets, prevRuns int64 = math.MaxInt64, 0
			for _, item := range got.Stats {
				if item.TenWktHauls.Int64 > 0 {
					t.Error("ten_wkt_hauls qualification failed")
					return
				}

				if item.BestInningsWkts.Int64 > prevWickets {
					t.Error("best_bowling_innings sorting failed")
					return
				} else if item.BestInningsWkts.Int64 == prevWickets && item.BestInningsRuns.Int64 < prevRuns {
					t.Error("best_bowling_innings sorting failed")
					return
				}

				prevWickets = item.BestInningsWkts.Int64
				prevRuns = item.BestInningsRuns.Int64
			}
		})
	}
}

func TestRead_Overall_Bowling_Continents_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Bowling_Continent_Group]
		wantErr bool
	}{
		{
			name: "Overall_Bowling_Continent_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":      []string{"ODI"},
					"is_male":             []string{"true"},
					"tournament":          []string{"1", "2"},
					"min__fours_conceded": []string{"50"},
					"sort_by":             []string{"fours_conceded"},
					"sort_order":          []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Bowling_Continents_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Bowling_Continents_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64
			for _, item := range got.Stats {
				if item.FoursConceded.Int64 < 50 {
					t.Error("fours_conceded qualification failed")
					return
				}

				if item.FoursConceded.Int64 < prev {
					t.Error("fours_conceded sorting failed")
					return
				}

				prev = item.FoursConceded.Int64
			}
		})
	}
}

func TestRead_Overall_Bowling_Series_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Bowling_Series_Group]
		wantErr bool
	}{
		{
			name: "Overall_Bowling_Series_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":      []string{"ODI"},
					"is_male":             []string{"true"},
					"match_result":        []string{"won", "tied", "drawn"},
					"min__sixes_conceded": []string{"25"},
					"sort_by":             []string{"sixes_conceded"},
					"sort_order":          []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Bowling_Series_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Bowling_Series_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64
			for _, item := range got.Stats {
				if item.SixesConceded.Int64 < 25 {
					t.Error("sixes_conceded qualification failed")
					return
				}

				if item.SixesConceded.Int64 < prev {
					t.Error("sixes_conceded sorting failed")
					return
				}

				prev = item.SixesConceded.Int64
			}
		})
	}
}

func TestRead_Overall_Bowling_Tournaments_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Bowling_Tournament_Group]
		wantErr bool
	}{
		{
			name: "Overall_Bowling_Tournament_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":      []string{"ODI"},
					"is_male":             []string{"true"},
					"toss_result":         []string{"won"},
					"min__innings_bowled": []string{"50"},
					"sort_by":             []string{"innings_bowled"},
					"sort_order":          []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Bowling_Tournaments_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Bowling_Tournaments_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64
			for _, item := range got.Stats {
				if item.InningsBowled.Int64 < 50 {
					t.Error("innings_bowled qualification failed")
					return
				}

				if item.InningsBowled.Int64 < prev {
					t.Error("innings_bowled sorting failed")
					return
				}

				prev = item.InningsBowled.Int64
			}
		})
	}
}

func TestRead_Overall_Bowling_Years_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Bowling_Year_Group]
		wantErr bool
	}{
		{
			name: "Overall_Bowling_Year_Group",
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
			got, err := Read_Overall_Bowling_Years_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Bowling_Years_Stats() error = %v, wantErr %v", err, tt.wantErr)
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

func TestRead_Overall_Bowling_Seasons_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Bowling_Season_Group]
		wantErr bool
	}{
		{
			name: "Overall_Bowling_Season_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":    []string{"ODI"},
					"is_male":           []string{"true"},
					"innings_number":    []string{"2"},
					"min__maiden_overs": []string{"2"},
					"sort_by":           []string{"maiden_overs"},
					"sort_order":        []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Bowling_Seasons_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Bowling_Seasons_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64
			for _, item := range got.Stats {
				if item.MaidenOvers.Int64 < 2 {
					t.Error("maiden_overs qualification failed")
					return
				}

				if item.MaidenOvers.Int64 < prev {
					t.Error("maiden_overs sorting failed")
					return
				}

				prev = item.MaidenOvers.Int64
			}
		})
	}
}

func TestRead_Overall_Bowling_Decades_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Bowling_Decade_Group]
		wantErr bool
	}{
		{
			name: "Overall_Bowling_Decade_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":             []string{"ODI"},
					"is_male":                    []string{"true"},
					"max__innings_runs_conceded": []string{"50"},
					"sort_by":                    []string{"economy"},
					"sort_order":                 []string{"default"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Bowling_Decades_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Bowling_Decades_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev float64
			for _, item := range got.Stats {
				if item.Economy.Valid && item.Economy.Float64 < prev {
					t.Error("economy sorting failed")
					return
				}

				if item.Economy.Valid {
					prev = item.Economy.Float64
				}
			}
		})
	}
}

func TestRead_Overall_Bowling_Aggregate_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Overall_Bowling_Aggregate_Group]
		wantErr bool
	}{
		{
			name: "Overall_Bowling_Aggregate_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":                []string{"ODI"},
					"min_start_date":                []string{"2018-01-01"},
					"is_male":                       []string{"true"},
					"min__innings_bowling_position": []string{"1"},
					"max__innings_bowling_position": []string{"3"},
					"toss_result":                   []string{"lost"},
					"bat_field_first":               []string{"field"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Overall_Bowling_Aggregate_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Overall_Bowling_Aggregate_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got.Stats) > 1 {
				t.Error("got more than more 1 item")
				return
			}
		})
	}
}

func TestRead_Individual_Bowling_Innings_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Individual_Bowling_Innings_Group]
		wantErr bool
	}{
		{
			name: "Individual_Bowling_Innings_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":                []string{"ODI"},
					"is_male":                       []string{"true"},
					"max__innings_runs_conceded":    []string{"50"},
					"min__innings_balls_bowled":     []string{"30"},
					"min__innings_wickets_taken":    []string{"4"},
					"min__innings_bowling_position": []string{"1"},
					"max__innings_bowling_position": []string{"4"},
					"sort_by":                       []string{"economy"},
					"sort_order":                    []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Individual_Bowling_Innings_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Bowling_Innings_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev float64 = math.MaxFloat64
			for _, item := range got.Stats {
				if item.RunsConceded.Int64 > 50 {
					t.Error("max__innings_runs_conceded qualification failed")
					return
				}

				if item.OversBowled.Float64 < 5 {
					t.Error("min__innings_balls_bowled qualification failed")
					return
				}

				if item.WicketsTaken.Int64 < 4 {
					t.Error("min__innings_wickets_taken qualification failed")
					return
				}

				if item.Economy.Valid && item.Economy.Float64 > prev {
					t.Error("economy sorting failed")
					return
				}

				if item.Economy.Valid {
					prev = item.Economy.Float64
				}
			}
		})
	}
}

func TestRead_Individual_Bowling_MatchTotals_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Individual_Bowling_MatchTotals_Group]
		wantErr bool
	}{
		{
			name: "Individual_Bowling_MatchTotals_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":                []string{"ODI"},
					"is_male":                       []string{"true"},
					"min__innings_runs_conceded":    []string{"30"},
					"max__innings_runs_conceded":    []string{"50"},
					"min__innings_bowling_position": []string{"3"},
					"max__innings_bowling_position": []string{"5"},
					"sort_by":                       []string{"start_date"},
					"sort_order":                    []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Individual_Bowling_MatchTotals_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Bowling_MatchTotals_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev time.Time
			for _, item := range got.Stats {
				if item.RunsConceded.Int64 < 30 {
					t.Error("min__innings_runs_conceded qualification failed")
					return
				}

				if item.RunsConceded.Int64 > 50 {
					t.Error("max__innings_runs_conceded qualification failed")
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

func TestRead_Individual_Bowling_Series_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Individual_Bowling_Series_Group]
		wantErr bool
	}{
		{
			name: "Individual_Bowling_Series_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":                []string{"ODI"},
					"is_male":                       []string{"true"},
					"tournament":                    []string{"1"},
					"min__innings_bowling_position": []string{"1"},
					"max__innings_bowling_position": []string{"7"},
					"sort_by":                       []string{"player_name"},
					"sort_order":                    []string{"default"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Read_Individual_Bowling_Series_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Bowling_Series_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRead_Individual_Bowling_Tournaments_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Individual_Bowling_Tournament_Group]
		wantErr bool
	}{
		{
			name: "Individual_Bowling_Tournament_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":   []string{"ODI"},
					"is_male":          []string{"true"},
					"home":             []string{"home"},
					"max__strike_rate": []string{"35"},
					"sort_by":          []string{"average"},
					"sort_order":       []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Individual_Bowling_Tournaments_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Bowling_Tournaments_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev float64 = math.MaxFloat64
			for _, item := range got.Stats {
				if item.StrikeRate.Valid && item.StrikeRate.Float64 > 35 {
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

func TestRead_Individual_Bowling_Grounds_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Individual_Bowling_Ground_Group]
		wantErr bool
	}{
		{
			name: "Individual_Bowling_Ground_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":   []string{"ODI"},
					"is_male":          []string{"true"},
					"host_nation":      []string{"1", "2"},
					"min__average":     []string{"25"},
					"max__strike_rate": []string{"35"},
					"min__economy":     []string{"5"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Individual_Bowling_Grounds_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Bowling_Grounds_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64 = math.MaxInt64
			for _, item := range got.Stats {
				if item.Average.Valid && item.Average.Float64 < 25 {
					t.Error("average qualification failed")
					return
				}

				if item.StrikeRate.Valid && item.StrikeRate.Float64 > 35 {
					t.Error("strike_rate qualification failed")
					return
				}

				if item.Economy.Valid && item.Economy.Float64 < 5 {
					t.Error("economy qualification failed")
					return
				}

				if item.WicketsTaken.Int64 > prev {
					t.Error("wickets_taken sorting failed")
				}

				prev = item.WicketsTaken.Int64
			}
		})
	}
}

func TestRead_Individual_Bowling_HostNations_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Individual_Bowling_HostNation_Group]
		wantErr bool
	}{
		{
			name: "Individual_Bowling_HostNation_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":      []string{"ODI"},
					"is_male":             []string{"true"},
					"continent":           []string{"3"},
					"min__matches_played": []string{"5"},
					"sort_by":             []string{"best_bowling_match"},
					"sort_order":          []string{"reverse"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Individual_Bowling_HostNations_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Bowling_HostNations_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prevWickets, prevRuns int64 = 0, math.MaxInt64
			for _, item := range got.Stats {
				if item.MatchesPlayed.Int64 < 5 {
					t.Error("matches_played qualification failed")
					return
				}

				if item.BestMatchWkts.Int64 < prevWickets {
					t.Error("best_bowling_match sorting failed")
					return
				} else if item.BestMatchWkts.Int64 == prevWickets && item.BestMatchRuns.Int64 > prevRuns {
					t.Error("best_bowling_match sorting failed")
					return
				}

				prevWickets = item.BestMatchWkts.Int64
				prevRuns = item.BestMatchRuns.Int64
			}
		})
	}
}

func TestRead_Individual_Bowling_Oppositions_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Individual_Bowling_Opposition_Group]
		wantErr bool
	}{
		{
			name: "Individual_Bowling_Opposition_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format":      []string{"ODI"},
					"is_male":             []string{"true"},
					"primary_team":        []string{"1", "2", "3"},
					"opposition_team":     []string{"1", "2", "3", "4", "5"},
					"max__innings_bowled": []string{"15"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Individual_Bowling_Oppositions_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Bowling_Oppositions_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64 = math.MaxInt64
			for _, item := range got.Stats {
				if item.InningsBowled.Int64 > 15 {
					t.Error("innings_bowled qualification failed")
					return
				}

				if item.WicketsTaken.Int64 > prev {
					t.Error("runs_scored sorting failed")
				}

				prev = item.WicketsTaken.Int64
			}
		})
	}
}

func TestRead_Individual_Bowling_Years_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Individual_Bowling_Year_Group]
		wantErr bool
	}{
		{
			name: "Individual_Bowling_Year_Group",
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
					"__page":         []string{"1"},
					"__limit":        []string{"200"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Individual_Bowling_Years_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Bowling_Years_Stats() error = %v, wantErr %v", err, tt.wantErr)
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

func TestRead_Individual_Bowling_Seasons_Stats(t *testing.T) {
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
		want    responses.StatsResponse[responses.Individual_Bowling_Season_Group]
		wantErr bool
	}{
		{
			name: "Individual_Bowling_Season_Group",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format": []string{"T20I"},
					"is_male":        []string{"true"},
					"min_start_date": []string{"2019-01-01"},
					"max_start_date": []string{"2023-12-31"},
					"sort_by":        []string{"five_wkt_hauls"},
					"sort_order":     []string{"reverse"},
					"__page":         []string{"1"},
					"__limit":        []string{"200"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Individual_Bowling_Seasons_Stats(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Individual_Bowling_Seasons_Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var prev int64 = math.MaxInt64
			for _, item := range got.Stats {
				if item.FiveWktHauls.Int64 > prev {
					t.Error("five_wkt_hauls sorting failed")
					return
				}

				prev = item.FiveWktHauls.Int64
			}
		})
	}
}
