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

	r.Post("/", createMatch)
	r.Get("/", getMatches)

	r.Get("/{matchId}/summary", getMatchSummary)
	r.Get("/{matchId}/full-scorecard", getMatchFullScorecard)
	r.Get("/{matchId}/squads", getMatchSquad)

	r.Mount("/{matchId}/innings", inningsRouter())

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

func getMatchSummary(w http.ResponseWriter, r *http.Request) {
	matchIdRaw := r.PathValue("matchId")

	matchId, err := strconv.ParseInt(matchIdRaw, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Message: "invalid match id", Data: nil}, http.StatusBadRequest)
		return
	}

	scorecard, err := dbutils.ReadMatchSummary(r.Context(), DB_POOL, matchId)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Message: "error while reading match summary", Data: err}, http.StatusInternalServerError)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched match summary", Data: scorecard}, http.StatusOK)
}

func getMatchFullScorecard(w http.ResponseWriter, r *http.Request) {
	matchIdRaw := r.PathValue("matchId")

	matchId, err := strconv.ParseInt(matchIdRaw, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Message: "invalid match id", Data: nil}, http.StatusBadRequest)
		return
	}

	scorecard, err := dbutils.ReadMatchFullScorecard(r.Context(), DB_POOL, matchId)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Message: "error while reading match scorecard", Data: err}, http.StatusInternalServerError)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched match scorecard", Data: scorecard}, http.StatusOK)
}

func getMatchSquad(w http.ResponseWriter, r *http.Request) {
	matchIdRaw := r.PathValue("matchId")

	matchId, err := strconv.ParseInt(matchIdRaw, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Message: "invalid match id", Data: nil}, http.StatusBadRequest)
		return
	}

	scorecard, err := dbutils.ReadSquadByMatchId(r.Context(), DB_POOL, matchId)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Message: "error while reading match squad", Data: err}, http.StatusInternalServerError)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched match squad", Data: scorecard}, http.StatusOK)
}
