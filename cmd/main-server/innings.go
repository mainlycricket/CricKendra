package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
	"github.com/mainlycricket/CricKendra/internal/utils"
)

/* Used in matchesRouter */

func inningsRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/{inningsId}/commentary", getMatchInningsDeliveries)

	// auth by controller
	r.Post("/", createInnings)
	r.Patch("/{inningsId}/innings-end", updateInningsEnd)
	r.Patch("/{inningsId}/current-batters", updateInningsCurrentBatters)
	r.Patch("/{inningsId}/current-bowlers", updateInningsCurrentBowlers)

	// auth by own middlewares
	r.Mount("/{inningsId}/batting-scorecards", battingScorecardsRouter())
	r.Mount("/{inningsId}/bowling-scorecards", bowlingScorecardsRouter())
	r.Mount("/{inningsId}/deliveries", deliveriesRouter())

	return r
}

func createInnings(w http.ResponseWriter, r *http.Request) {
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

	var innings models.Innings

	if err := json.NewDecoder(r.Body).Decode(&innings); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	innings.MatchId.Int64, innings.MatchId.Valid = parsedMatchId, true

	inningsId, err := dbutils.InsertInnings(r.Context(), DB_POOL, &innings)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while inserting innings", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "innings created successfully", Data: inningsId}, http.StatusCreated)
}

func updateInningsEnd(w http.ResponseWriter, r *http.Request) {
	_, err := utils.AuthorizeRequest(r, []string{SYSTEM_ADMIN_ROLE})
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "unauthorized request", Data: err}, http.StatusUnauthorized)
		return
	}

	rawInningsId := r.PathValue("inningsId")
	parsedInningsId, err := strconv.ParseInt(rawInningsId, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid innings id", Data: err}, http.StatusBadRequest)
		return
	}

	var input models.InningsEndInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	input.InningsId.Int64, input.InningsId.Valid = parsedInningsId, true

	if err := dbutils.UpdateInningsEnd(r.Context(), DB_POOL, &input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while updating innings end", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "innings end updated successfully", Data: nil}, http.StatusOK)
}

func updateInningsCurrentBatters(w http.ResponseWriter, r *http.Request) {
	_, err := utils.AuthorizeRequest(r, []string{SYSTEM_ADMIN_ROLE})
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "unauthorized request", Data: err}, http.StatusUnauthorized)
		return
	}

	rawInningsId := r.PathValue("inningsId")
	parsedInningsId, err := strconv.ParseInt(rawInningsId, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid innings id", Data: err}, http.StatusBadRequest)
		return
	}

	var input models.InningsCurrentBattersInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	input.InningsId.Int64, input.InningsId.Valid = parsedInningsId, true

	if err := dbutils.UpdateInningsCurrentBatters(r.Context(), DB_POOL, &input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while updating innings current batters", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "innings current batters updated successfully", Data: nil}, http.StatusOK)
}

func updateInningsCurrentBowlers(w http.ResponseWriter, r *http.Request) {
	_, err := utils.AuthorizeRequest(r, []string{SYSTEM_ADMIN_ROLE})
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "unauthorized request", Data: err}, http.StatusUnauthorized)
		return
	}

	rawInningsId := r.PathValue("inningsId")
	parsedInningsId, err := strconv.ParseInt(rawInningsId, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid innings id", Data: err}, http.StatusBadRequest)
		return
	}

	var input models.InningsCurrentBowlersInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	input.InningsId.Int64, input.InningsId.Valid = parsedInningsId, true

	if err := dbutils.UpdateInningsCurrentBowlers(r.Context(), DB_POOL, &input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while updating innings current bowlers", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "innings current bowlers updated successfully", Data: nil}, http.StatusOK)
}

func getMatchInningsDeliveries(w http.ResponseWriter, r *http.Request) {
	matchIdRaw, inningsIdRaw := r.PathValue("matchId"), r.PathValue("inningsId")

	matchId, err := strconv.ParseInt(matchIdRaw, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Message: "invalid match id", Data: nil}, http.StatusBadRequest)
		return
	}

	parsedInningsId, err := strconv.ParseInt(inningsIdRaw, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Message: "invalid innings id", Data: err}, http.StatusBadRequest)
		return
	}

	commentary, err := dbutils.ReadDeliveriesByMatchInnings(r.Context(), DB_POOL, matchId, parsedInningsId)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Message: "error while reading innings commentary", Data: err}, http.StatusInternalServerError)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched innings commentary", Data: commentary}, http.StatusOK)
}
