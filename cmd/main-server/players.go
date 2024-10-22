package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/models"
)

func playersRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/{playerId}", getPlayerById)
	r.Get("/", getPlayers)
	r.Post("/", createPlayer)
	return r
}

func createPlayer(w http.ResponseWriter, r *http.Request) {
	var player models.Player

	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		writeJsonResponse(w, r, ApiResponse{false, "error while decoding json", err}, http.StatusBadRequest)
		return
	}

	err = dbutils.InsertPlayer(r.Context(), DB_POOL, &player)
	if err != nil {
		writeJsonResponse(w, r, ApiResponse{false, "error while inserting player", err}, http.StatusBadRequest)
		return
	}

	writeJsonResponse(w, r, ApiResponse{true, "player created successfully", nil}, http.StatusCreated)
}

func getPlayers(w http.ResponseWriter, r *http.Request) {
	players, err := dbutils.ReadPlayers(r.Context(), DB_POOL)

	if err != nil {
		writeJsonResponse(w, r, ApiResponse{false, "error while reading players", err}, http.StatusBadRequest)
		return
	}

	writeJsonResponse(w, r, ApiResponse{true, "fetched players successfully", players}, http.StatusOK)
}

func getPlayerById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("playerId")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		writeJsonResponse(w, r, ApiResponse{false, "invalid player id", err}, http.StatusBadRequest)
		return
	}

	players, err := dbutils.ReadPlayerById(r.Context(), DB_POOL, int_id)

	if err != nil {
		writeJsonResponse(w, r, ApiResponse{false, "error while reading player", err}, http.StatusBadRequest)
		return
	}

	writeJsonResponse(w, r, ApiResponse{true, "fetched player successfully", players}, http.StatusOK)
}
