package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/backend/internal/dbutils"
	"github.com/mainlycricket/CricKendra/backend/internal/models"
	"github.com/mainlycricket/CricKendra/backend/internal/responses"
	"github.com/mainlycricket/CricKendra/backend/internal/utils"
)

func playersRouter() *chi.Mux {
	r := chi.NewRouter()

	// auth by controller
	r.Post("/", createPlayer)

	r.Get("/", getPlayers)
	r.Get("/{playerId}", getPlayerById)

	return r
}

func createPlayer(w http.ResponseWriter, r *http.Request) {
	_, err := utils.AuthorizeRequest(r, []string{SYSTEM_ADMIN_ROLE})
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "unauthorized request", Data: err}, http.StatusUnauthorized)
		return
	}

	var player models.Player

	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	playerId, err := dbutils.InsertPlayer(r.Context(), DB_POOL, &player)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while inserting player", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "player created successfully", Data: playerId}, http.StatusCreated)
}

func getPlayers(w http.ResponseWriter, r *http.Request) {
	players, err := dbutils.ReadPlayers(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading players", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched players successfully", Data: players}, http.StatusOK)
}

func getPlayerById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("playerId")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid player id", Data: err}, http.StatusBadRequest)
		return
	}

	players, err := dbutils.ReadPlayerById(r.Context(), DB_POOL, int_id)

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading player", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched player successfully", Data: players}, http.StatusOK)
}
