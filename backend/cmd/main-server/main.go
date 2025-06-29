package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mainlycricket/CricKendra/backend/internal/dbutils"
)

// user roles for auth
const SYSTEM_ADMIN_ROLE string = "system_admin"

var DB_POOL *pgxpool.Pool

func main() {
	if err := initDB(); err != nil {
		log.Fatalf("db init failed: %v", err)
	}

	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   strings.Split(os.Getenv("FRONTEND_URL"), ","),
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// api routes
	r.Mount("/teams", teamsRouter())
	r.Mount("/players", playersRouter())
	r.Mount("/tournaments", tournamentsRouter())
	r.Mount("/series", seriesRouter())
	r.Mount("/seasons", seasonsRouter())
	r.Mount("/continents", continentsRouter())
	r.Mount("/host-nations", hostNationsRouter())
	r.Mount("/cities", citiesRouter())
	r.Mount("/grounds", groundsRouter())
	r.Mount("/matches", matchesRouter())
	r.Mount("/users", usersRouter())
	r.Mount("/options", optionsRouter())

	/* Stats */
	r.Mount("/stats/filter-options", StatFiltersRouter())
	r.Mount("/stats/batting", BattingStatsRouter())
	r.Mount("/stats/bowling", BowlingStatsRouter())
	r.Mount("/stats/team", TeamStatsRouter())

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf(`error while starting server: %v`, err)
	}
}

func initDB() error {
	var err error

	ctx, DB_URL := context.Background(), os.Getenv("DB_URL")
	DB_POOL, err = dbutils.Connect(ctx, DB_URL)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	return nil
}
