package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/models"
)

func teamsRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", getTeams)
	r.Post("/", createTeam)
	return r
}

func createTeam(w http.ResponseWriter, r *http.Request) {
	var team models.Team

	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		writeJsonResponse(w, r, ApiResponse{false, "error while decoding json", err}, http.StatusBadRequest)
		return
	}

	err = dbutils.InsertTeam(r.Context(), DB_POOL, &team)
	if err != nil {
		writeJsonResponse(w, r, ApiResponse{false, "error while inserting team", err}, http.StatusBadRequest)
		return
	}

	writeJsonResponse(w, r, ApiResponse{true, "team created successfully", nil}, http.StatusCreated)
}

func getTeams(w http.ResponseWriter, r *http.Request) {
	teams, err := dbutils.ReadTeams(r.Context(), DB_POOL)

	if err != nil {
		writeJsonResponse(w, r, ApiResponse{false, "error while reading teams", err}, http.StatusBadRequest)
		return
	}

	writeJsonResponse(w, r, ApiResponse{true, "teams read successfully", teams}, http.StatusOK)
}
