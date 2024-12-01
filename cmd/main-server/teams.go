package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
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
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	teamId, err := dbutils.InsertTeam(r.Context(), DB_POOL, &team)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while inserting team", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "team created successfully", Data: teamId}, http.StatusCreated)
}

func getTeams(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.ReadTeams(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading teams", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "teams read successfully", Data: response}, http.StatusOK)
}
