package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/pkg/dotenv"
)

var DB_POOL *pgxpool.Pool

func main() {
	err := initDB()
	if err != nil {
		log.Fatalf("db init failed: %v", err)
	}

	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{os.Getenv("FRONTEND_URL")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

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
	r.Mount("/innings", inningsRouter())
	r.Mount("/users", usersRouter())
	r.Mount("/blog-articles", blogArticlesRouter())

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf(`error while starting server: %v`, err)
	}
}

func initDB() error {
	basePath, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to read base path: %v", err)
	}

	dotEnvPath := filepath.Join(basePath, ".env")
	err = dotenv.ReadDotEnv(dotEnvPath)
	if err != nil {
		return fmt.Errorf("error while reading .env file: %v", err)
	}

	ctx, DB_URL := context.Background(), os.Getenv("DB_URL")
	DB_POOL, err = dbutils.Connect(ctx, DB_URL)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	return nil
}
