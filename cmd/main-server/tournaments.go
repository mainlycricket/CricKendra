package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
	"github.com/mainlycricket/CricKendra/internal/utils"
)

func tournamentsRouter() *chi.Mux {
	r := chi.NewRouter()

	// auth by controller
	r.Post("/", createTournament)

	r.Get("/", getTournaments)

	return r
}

func createTournament(w http.ResponseWriter, r *http.Request) {
	_, err := utils.AuthorizeRequest(r, []string{SYSTEM_ADMIN_ROLE})
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "unauthorized request", Data: err}, http.StatusUnauthorized)
		return
	}

	var tournament models.Tournament

	if err := json.NewDecoder(r.Body).Decode(&tournament); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	tournamentId, err := dbutils.InsertTournament(r.Context(), DB_POOL, &tournament)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while inserting tournament", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "tournament created successfully", Data: tournamentId}, http.StatusCreated)
}

func getTournaments(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.ReadTournaments(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading tournaments", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "tournaments read successfully", Data: response}, http.StatusOK)
}
