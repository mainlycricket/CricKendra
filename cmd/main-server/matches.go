package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
)

func matchesRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/match-result-options", getMatchResultOptions)

	r.Get("/match-type-options", getMatchTypeOptions)

	r.Get("/match-format-options", getMatchFormats)

	r.Get("/match-level-options", getMatchLevels)

	r.Post("/", createMatch)
	r.Get("/", getMatches)
	r.Get("/{matchId}", getMatchById)

	return r
}

func createMatch(w http.ResponseWriter, r *http.Request) {
	var match models.Match

	err := json.NewDecoder(r.Body).Decode(&match)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	matchId, err := dbutils.InsertMatch(r.Context(), DB_POOL, &match)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while inserting match", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "match created successfully", Data: matchId}, http.StatusCreated)
}

func getMatches(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.ReadMatches(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading matches", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "matches read successfully", Data: response}, http.StatusOK)
}

func getMatchById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("matchId")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid match id", Data: err}, http.StatusBadRequest)
		return
	}

	match, err := dbutils.ReadMatchById(r.Context(), DB_POOL, int_id)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading match", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched match successfully", Data: match}, http.StatusOK)
}

func getMatchResultOptions(w http.ResponseWriter, r *http.Request) {
	teams, err := dbutils.ReadMatchResultOptions(r.Context(), DB_POOL)

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading match result options", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "match result options read successfully", Data: teams}, http.StatusOK)
}

func getMatchTypeOptions(w http.ResponseWriter, r *http.Request) {
	teams, err := dbutils.ReadMatchTypeOptions(r.Context(), DB_POOL)

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading match type options", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "match type options read successfully", Data: teams}, http.StatusOK)
}

func getMatchFormats(w http.ResponseWriter, r *http.Request) {
	teams, err := dbutils.ReadMatchFormats(r.Context(), DB_POOL)

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading match formats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "match formats read successfully", Data: teams}, http.StatusOK)
}

func getMatchLevels(w http.ResponseWriter, r *http.Request) {
	teams, err := dbutils.ReadMatchLevels(r.Context(), DB_POOL)

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading match levels", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "match levels read successfully", Data: teams}, http.StatusOK)
}
