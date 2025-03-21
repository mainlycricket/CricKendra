package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
)

func squadsRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/match-entries", createMatch)

	return r
}

func upsertMatchSquadEntries(w http.ResponseWriter, r *http.Request) {
	var entries []models.MatchSquad

	err := json.NewDecoder(r.Body).Decode(&entries)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	if err := dbutils.UpsertMatchSquadEntries(r.Context(), DB_POOL, entries); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while upserting match squad entries", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "match squad entries upserted successfully", Data: nil}, http.StatusCreated)
}
