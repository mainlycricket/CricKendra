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

func matchesRouter() *chi.Mux {
	r := chi.NewRouter()

	// auth by controller
	r.Post("/", createMatch)
	r.Post("/{matchId}/upsert-squad", upsertMatchSquadEntries)
	r.Patch("/{matchId}/toss-decision", updateMatchTossDecision)
	r.Patch("/{matchId}/match-result", updateMatchResult)
	r.Patch("/{matchId}/match-state", updateMatchState)
	r.Post("/{matchId}/player-awards", upsertMatchPlayerAwards)

	r.Get("/", getMatches)
	r.Get("/{matchId}/summary", getMatchSummary)
	r.Get("/{matchId}/full-scorecard", getMatchFullScorecard)
	r.Get("/{matchId}/stats", getMatchStats)
	r.Get("/{matchId}/squads", getMatchSquad)

	// auth mixed - see inningsRouter
	r.Mount("/{matchId}/innings", inningsRouter())

	return r
}

func createMatch(w http.ResponseWriter, r *http.Request) {
	_, err := utils.AuthorizeRequest(r, []string{SYSTEM_ADMIN_ROLE})
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "unauthorized request", Data: err}, http.StatusUnauthorized)
		return
	}

	var match models.Match

	if err := json.NewDecoder(r.Body).Decode(&match); err != nil {
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

func upsertMatchSquadEntries(w http.ResponseWriter, r *http.Request) {
	_, err := utils.AuthorizeRequest(r, []string{SYSTEM_ADMIN_ROLE})
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "unauthorized request", Data: err}, http.StatusUnauthorized)
		return
	}

	matchIdRaw := r.PathValue("matchId")
	parsedMatchId, err := strconv.ParseInt(matchIdRaw, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Message: "invalid match id", Data: nil}, http.StatusBadRequest)
		return
	}

	var entries []models.MatchSquad
	if err := json.NewDecoder(r.Body).Decode(&entries); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	for idx, entry := range entries {
		entry.MatchId.Int64, entry.MatchId.Valid = parsedMatchId, true
		entries[idx] = entry
	}

	if err := dbutils.UpsertMatchSquadEntries(r.Context(), DB_POOL, entries); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while upserting match squad entries", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "match squad entries upserted successfully", Data: nil}, http.StatusCreated)
}

func updateMatchTossDecision(w http.ResponseWriter, r *http.Request) {
	_, err := utils.AuthorizeRequest(r, []string{SYSTEM_ADMIN_ROLE})
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "unauthorized request", Data: err}, http.StatusUnauthorized)
		return
	}

	matchIdRaw := r.PathValue("matchId")

	parsedMatchId, err := strconv.ParseInt(matchIdRaw, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Message: "invalid match id", Data: nil}, http.StatusBadRequest)
		return
	}

	var input models.TossDecisionInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	input.MatchId.Int64, input.MatchId.Valid = parsedMatchId, true

	if err := dbutils.UpdateMatchTossDecisionById(r.Context(), DB_POOL, &input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while updating match toss decision", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "match toss decision updated successfully", Data: nil}, http.StatusOK)
}

func updateMatchResult(w http.ResponseWriter, r *http.Request) {
	_, err := utils.AuthorizeRequest(r, []string{SYSTEM_ADMIN_ROLE})
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "unauthorized request", Data: err}, http.StatusUnauthorized)
		return
	}

	matchIdRaw := r.PathValue("matchId")

	parsedMatchId, err := strconv.ParseInt(matchIdRaw, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Message: "invalid match id", Data: nil}, http.StatusBadRequest)
		return
	}

	var input models.MatchResultInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	input.MatchId.Int64, input.MatchId.Valid = parsedMatchId, true

	if err := dbutils.UpdateMatchResultById(r.Context(), DB_POOL, &input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while updating match result", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "match result updated successfully", Data: nil}, http.StatusOK)
}

func updateMatchState(w http.ResponseWriter, r *http.Request) {
	_, err := utils.AuthorizeRequest(r, []string{SYSTEM_ADMIN_ROLE})
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "unauthorized request", Data: err}, http.StatusUnauthorized)
		return
	}

	matchIdRaw := r.PathValue("matchId")
	parsedMatchId, err := strconv.ParseInt(matchIdRaw, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Message: "invalid match id", Data: nil}, http.StatusBadRequest)
		return
	}

	var input models.MatchStateInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	input.MatchId.Int64, input.MatchId.Valid = parsedMatchId, true

	if err := input.Validate(r.Context(), DB_POOL); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid operation", Data: err}, http.StatusBadRequest)
		return
	}

	if err := dbutils.UpdateMatchStateById(r.Context(), DB_POOL, &input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while updating match state", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "match state updated successfully", Data: nil}, http.StatusOK)
}

func upsertMatchPlayerAwards(w http.ResponseWriter, r *http.Request) {
	_, err := utils.AuthorizeRequest(r, []string{SYSTEM_ADMIN_ROLE})
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "unauthorized request", Data: err}, http.StatusUnauthorized)
		return
	}

	matchIdRaw := r.PathValue("matchId")
	parsedMatchId, err := strconv.ParseInt(matchIdRaw, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Message: "invalid match id", Data: nil}, http.StatusBadRequest)
		return
	}

	var entries []models.PlayerAwardEntry

	if err := json.NewDecoder(r.Body).Decode(&entries); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	for idx, entry := range entries {
		entry.MatchId.Int64, entry.MatchId.Valid = parsedMatchId, true
		entries[idx] = entry
	}

	if err := dbutils.UpsertMatchAwardEntries(r.Context(), DB_POOL, entries); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while upserting match award entries", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "match award entries upserted successfully", Data: nil}, http.StatusCreated)
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

func getMatchStats(w http.ResponseWriter, r *http.Request) {
	matchIdRaw := r.PathValue("matchId")

	matchId, err := strconv.ParseInt(matchIdRaw, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Message: "invalid match id", Data: nil}, http.StatusBadRequest)
		return
	}

	scorecard, err := dbutils.ReadMatchStats(r.Context(), DB_POOL, matchId)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Message: "error while reading match stats", Data: err}, http.StatusInternalServerError)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched match stats", Data: scorecard}, http.StatusOK)
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
