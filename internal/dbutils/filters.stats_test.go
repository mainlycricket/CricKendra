package dbutils

import (
	"context"
	"log"
	"net/url"
	"os"
	"testing"

	"github.com/mainlycricket/CricKendra/internal/responses"
)

func TestRead_Stat_Filter_Options(t *testing.T) {
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
		want    responses.StatsFilters
		wantErr bool
	}{
		{
			name: "Read_Stat_Filter_Options",
			args: args{
				ctx: context.Background(),
				db:  DB_POOL,
				queryMap: url.Values{
					"playing_format": []string{"ODI"},
					"is_male":        []string{"true"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read_Stat_Filter_Options(tt.args.ctx, tt.args.db, tt.args.queryMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read_Stat_Filter_Options() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !got.MinDate.Valid {
				t.Error("invalid min_date")
				return
			}

			if !got.MaxDate.Valid {
				t.Error("invalid max_date")
				return
			}

			if len(got.Seasons) == 0 {
				t.Error("seasons less than 1")
				return
			}

			if len(got.Teams) == 0 {
				t.Error("teams less than 1")
				return
			}

			if len(got.Grounds) == 0 {
				t.Error("grounds less than 1")
				return
			}

			if len(got.HostNations) == 0 {
				t.Error("hostNatins less than 1")
				return
			}

			if len(got.Continents) == 0 {
				t.Error("continents less than 1")
				return
			}

			if len(got.Series) == 0 {
				t.Error("series less than 1")
				return
			}

			if len(got.Tournaments) == 0 {
				t.Error("tournaments less than 1")
				return
			}
		})
	}
}
