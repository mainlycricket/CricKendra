package dbutils

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

/* Blog Articles */

func ReadArticleCategories(ctx context.Context, db *pgxpool.Pool) ([]string, error) {
	query := `SELECT unnest(enum_range(null::article_category))`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	categories, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (string, error) {
		var category string
		err := row.Scan(&category)
		return category, err
	})

	return categories, err
}

func ReadArticleStatusOptions(ctx context.Context, db *pgxpool.Pool) ([]string, error) {
	query := `SELECT unnest(enum_range(null::article_status))`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	statusOptions, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (string, error) {
		var status string
		err := row.Scan(&status)
		return status, err
	})

	return statusOptions, err
}

/* Innings */

func ReadInningsEndOptions(ctx context.Context, db *pgxpool.Pool) ([]string, error) {
	query := `SELECT unnest(enum_range(null::innings_end))`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	inningsEndOptions, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (string, error) {
		var inningsEnd string
		err := row.Scan(&inningsEnd)
		return inningsEnd, err
	})

	return inningsEndOptions, err
}

/* Players */

func ReadBowlingStyleOptions(ctx context.Context, db *pgxpool.Pool) ([]string, error) {
	query := `SELECT unnest(enum_range(null::bowling_style))`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	bowlingStyles, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (string, error) {
		var bowlingStyle string
		err := row.Scan(&bowlingStyle)
		return bowlingStyle, err
	})

	return bowlingStyles, err
}

func ReadDismissalTypeOptions(ctx context.Context, db *pgxpool.Pool) ([]string, error) {
	query := `SELECT unnest(enum_range(null::dismissal_type))`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	dismissalTypes, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (string, error) {
		var dismissalType string
		err := row.Scan(&dismissalType)
		return dismissalType, err
	})

	return dismissalTypes, err
}

func ReadPlayingStatusOptions(ctx context.Context, db *pgxpool.Pool) ([]string, error) {
	query := `SELECT unnest(enum_range(null::playing_status))`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	playingStatusOptions, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (string, error) {
		var playingStatus string
		err := row.Scan(&playingStatus)
		return playingStatus, err
	})

	return playingStatusOptions, err
}

/* Matches */

func ReadMatchResultOptions(ctx context.Context, db *pgxpool.Pool) ([]string, error) {
	query := `SELECT unnest(enum_range(null::match_final_result))`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	matchResultOptions, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (string, error) {
		var matchResult string
		err := row.Scan(&matchResult)
		return matchResult, err
	})

	return matchResultOptions, err
}

func ReadMatchTypeOptions(ctx context.Context, db *pgxpool.Pool) ([]string, error) {
	query := `SELECT unnest(enum_range(null::match_type))`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	matchTypeOptions, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (string, error) {
		var matchType string
		err := row.Scan(&matchType)
		return matchType, err
	})

	return matchTypeOptions, err
}

func ReadMatchFormats(ctx context.Context, db *pgxpool.Pool) ([]string, error) {
	query := `SELECT unnest(enum_range(null::playing_format))`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	matchFormats, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (string, error) {
		var matchFormat string
		err := row.Scan(&matchFormat)
		return matchFormat, err
	})

	return matchFormats, err
}

func ReadMatchLevels(ctx context.Context, db *pgxpool.Pool) ([]string, error) {
	query := `SELECT unnest(enum_range(null::playing_level))`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	playingLevels, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (string, error) {
		var playingLevel string
		err := row.Scan(&playingLevel)
		return playingLevel, err
	})

	return playingLevels, err
}

/* Users */

func ReadUserRoleOptions(ctx context.Context, db *pgxpool.Pool) ([]string, error) {
	query := `SELECT unnest(enum_range(null::user_role))`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	userRoleOptions, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (string, error) {
		var role string
		err := row.Scan(&role)
		return role, err
	})

	return userRoleOptions, err
}
