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

/* Used in inningsRouter */

func battingScorecardsRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(utils.AuthorizationMiddleware([]string{SYSTEM_ADMIN_ROLE}))

	r.Post("/entries", createBattingScorecardEntries)
	r.Patch("/{batterId}/batting-position", updateBatterPositionByInningsId)

	return r
}

func createBattingScorecardEntries(w http.ResponseWriter, r *http.Request) {
	rawInningsId := r.PathValue("inningsId")
	parsedInningsId, err := strconv.ParseInt(rawInningsId, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid innings id", Data: err}, http.StatusBadRequest)
		return
	}

	var entries []models.BattingScorecard

	if err := json.NewDecoder(r.Body).Decode(&entries); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	for idx, entry := range entries {
		entry.InningsId.Int64, entry.InningsId.Valid = parsedInningsId, true
		entries[idx] = entry
	}

	if err := dbutils.InsertBattingScorecardEntries(r.Context(), DB_POOL, entries); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while inserting batting scorecard entries", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "batting scorecard entries created successfully", Data: nil}, http.StatusCreated)
}

func updateBatterPositionByInningsId(w http.ResponseWriter, r *http.Request) {
	rawInningsId, rawBatterId := r.PathValue("inningsId"), r.PathValue("batterId")

	parsedInningsId, err := strconv.ParseInt(rawInningsId, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid innings id", Data: err}, http.StatusBadRequest)
		return
	}

	parsedBatterId, err := strconv.ParseInt(rawBatterId, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid batter id", Data: err}, http.StatusBadRequest)
		return
	}

	var input models.BatterPositionInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	input.InningsId.Int64, input.InningsId.Valid = parsedInningsId, true
	input.BatterId.Int64, input.BatterId.Valid = parsedBatterId, true

	if err := dbutils.UpdateBatterPositionByInningsId(r.Context(), DB_POOL, &input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while updating batter batting position", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "batter batting position updated successfully", Data: nil}, http.StatusOK)
}
