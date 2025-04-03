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

func bowlingScorecardsRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(utils.AuthorizationMiddleware([]string{SYSTEM_ADMIN_ROLE}))

	r.Post("/", createBowlingScorecardEntry)

	return r
}

func createBowlingScorecardEntry(w http.ResponseWriter, r *http.Request) {
	rawInningsId := r.PathValue("inningsId")
	parsedInningsId, err := strconv.ParseInt(rawInningsId, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid innings id", Data: err}, http.StatusBadRequest)
		return
	}

	var entry models.BowlingScorecard

	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	entry.InningsId.Int64, entry.InningsId.Valid = parsedInningsId, true

	if err := dbutils.InsertBowlingScorecardEntry(r.Context(), DB_POOL, &entry); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while inserting bowling scorecard entry", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "bowling scorecard entry created successfully", Data: nil}, http.StatusCreated)
}
