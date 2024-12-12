package statqueries

import (
	"fmt"
	"net/url"
	"strings"
)

func HandlePlayingFormat(params *url.Values, filters *[]string, args *[]any) {
	playing_format := params.Get("playing_format")
	if playing_format != "" {
		*args = append(*args, playing_format)
		*filters = append(*filters, fmt.Sprintf(`matches.playing_format = $%d`, len(*args)))
	}
}

func HandleIsMale(params *url.Values, filters *[]string, args *[]any) {
	is_male := params.Get("is_male")
	if is_male != "" {
		*args = append(*args, is_male)
		*filters = append(*filters, fmt.Sprintf(`matches.is_male::bool = $%d`, len(*args)))
	}
}

func HandleMinStartDate(params *url.Values, filters *[]string, args *[]any) {
	min_start_date := params.Get("min_start_date")
	if min_start_date != "" {
		*args = append(*args, min_start_date)
		*filters = append(*filters, fmt.Sprintf(`matches.start_date::date >= $%d`, len(*args)))
	}
}

func HandleMaxStartDate(params *url.Values, filters *[]string, args *[]any) {
	max_start_date := params.Get("max_start_date")
	if max_start_date != "" {
		*args = append(*args, max_start_date)
		*filters = append(*filters, fmt.Sprintf(`matches.start_date::date <= $%d`, len(*args)))
	}
}

func HandleSeasons(params *url.Values, filters *[]string, args *[]any) {
	var placeholders []string

	seasons := (*params)["season"]
	for _, season := range seasons {
		*args = append(*args, season)
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, len(*args)))
	}

	if len(placeholders) > 0 {
		*filters = append(*filters, fmt.Sprintf(`matches.season IN (%s)`, strings.Join(placeholders, ", ")))
	}
}

func HandleGround(params *url.Values, filters *[]string, args *[]any) {
	var placeholders []string

	grounds_id := (*params)["ground"]
	for _, ground_id := range grounds_id {
		*args = append(*args, ground_id)
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, len(*args)))
	}

	if len(placeholders) > 0 {
		*filters = append(*filters, fmt.Sprintf(`matches.ground_id IN (%s)`, strings.Join(placeholders, ", ")))
	}
}

func HandleBattingTeam(params *url.Values, filters *[]string, args *[]any) {
	var placeholders []string

	teams_id := (*params)["batting_team"]
	for _, team_id := range teams_id {
		*args = append(*args, team_id)
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, len(*args)))
	}

	if len(placeholders) > 0 {
		*filters = append(*filters, fmt.Sprintf(`innings.batting_team_id IN (%s)`, strings.Join(placeholders, ", ")))
	}
}

func HandleBowlingTeam(params *url.Values, filters *[]string, args *[]any) {
	var placeholders []string

	teams_id := (*params)["bowling_team"]
	for _, team_id := range teams_id {
		*args = append(*args, team_id)
		placeholders = append(placeholders, fmt.Sprintf(`$%d`, len(*args)))
	}

	if len(placeholders) > 0 {
		*filters = append(*filters, fmt.Sprintf(`innings.bowling_team_id IN (%s)`, strings.Join(placeholders, ", ")))
	}
}
