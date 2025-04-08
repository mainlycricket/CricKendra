package dbutils

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/utils"
)

type DB_Exec interface {
	Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, query string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) pgx.Row
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
}

func Connect(ctx context.Context, connectionUrl string) (*pgxpool.Pool, error) {
	if os.Getenv("ENV") == "DOCKER" {
		time.Sleep(5 * time.Second)
	}

	config, err := pgxpool.ParseConfig(connectionUrl)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config: %w", err)
	}

	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		return registerDataTypes(ctx, conn)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	if err := upsertAdmin(pool); err != nil {
		return nil, fmt.Errorf(`failed to upsert admin: %v`, err)
	}

	return pool, nil
}

func registerDataTypes(ctx context.Context, conn *pgx.Conn) error {
	dataTypeNames := []string{
		"career_stats",
		"dismissal_type",
		"innings_end",
	}

	for _, typeName := range dataTypeNames {
		dataType, err := conn.LoadType(ctx, typeName)
		if err != nil {
			return fmt.Errorf("failed to register custom type: %s", typeName)
		}
		conn.TypeMap().RegisterType(dataType)
	}

	return nil
}

func upsertAdmin(db *pgxpool.Pool) error {
	email, password := os.Getenv("ADMIN_EMAIL"), os.Getenv("ADMIN_PASSWORD")
	if email == "" || password == "" {
		return errors.New("email or password not present")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := models.User{
		Name:     pgtype.Text{String: "System Admin", Valid: true},
		Email:    pgtype.Text{String: email, Valid: true},
		Password: pgtype.Text{String: hashedPassword, Valid: true},
		Role:     pgtype.Text{String: "system_admin", Valid: true},
	}

	_, err = UpsertUser(context.Background(), db, &user)
	return err
}
